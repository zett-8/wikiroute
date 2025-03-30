package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/xwb1989/sqlparser"
)

// Make a channel that yields rows from a SQL file raw by raw.
func SQLValueRowIterator(filePath string) (<-chan []string, <-chan error) {
	out := make(chan []string)
	errCh := make(chan error, 1)

	go func() {
		defer close(out)
		defer close(errCh)

		file, err := os.Open(filePath)
		if err != nil {
			errCh <- fmt.Errorf("failed to open file: %w", err)
			return
		}
		defer file.Close()

		var reader io.Reader = file

		// If the file is gzipped, create a gzip reader
		if filepath.Ext(filePath) == ".gz" {
			gzipReader, err := gzip.NewReader(file)
			if err != nil {
				errCh <- fmt.Errorf("failed to create gzip reader: %w", err)
				return
			}
			defer gzipReader.Close()
			reader = gzipReader
		}

		tokenizer := sqlparser.NewTokenizer(bufio.NewReader(reader))

		for {
			stmt, err := sqlparser.ParseNext(tokenizer)
			if err != nil {
				if err == io.EOF {
					break
				}
				continue
			}

			insertStmt, ok := stmt.(*sqlparser.Insert)
			if !ok {
				continue
			}

			switch rows := insertStmt.Rows.(type) {
			case sqlparser.Values:
				for _, row := range rows {
					var strRow []string
					for _, val := range row {
						switch v := val.(type) {
						case *sqlparser.SQLVal:
							if v.Type == sqlparser.StrVal {
								strRow = append(strRow, string(v.Val))
							} else {
								strRow = append(strRow, sqlparser.String(val))
							}
						default:
							strRow = append(strRow, sqlparser.String(val))
						}
					}
					out <- strRow
				}
			}
		}

		errCh <- nil
	}()

	return out, errCh
}
