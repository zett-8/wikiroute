package main

import (
	"encoding/gob"
	"os"
)

func Output(data interface{}, outPath string) error {
	outFile, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	enc := gob.NewEncoder(outFile)
	if err := enc.Encode(data); err != nil {
		return err
	}

	return nil
}
