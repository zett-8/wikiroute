package main

import (
	"fmt"
	"strconv"
	"wikiroute/core"
)

// Generate pagesMap from jawiki-20250301-page.sql
// Columns: page_id, page_namespace, page_title, page_is_redirect, page_is_new, page_random, page_touched, page_links_updated, page_latest, page_len, page_content_model, page_lang
// output: PageData
// data type: page_id: int64, page_title: string
func GeneratePagesData() error {
	valsChan, errChan := SQLValueRowIterator("./sql/jawiki-20250301-page.sql")

	pageData := &core.PageData{
		IDToTitle: make(map[int32]string),
		TitleToID: make(map[string]int32),
	}

	count := 0

	for vals := range valsChan {
		page_namespace, err := strconv.ParseInt(vals[1], 10, 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		if page_namespace != 0 {
			continue
		}

		page_id, err := strconv.ParseInt(vals[0], 10, 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		page_title := vals[2]

		pageData.IDToTitle[int32(page_id)] = page_title
		pageData.TitleToID[page_title] = int32(page_id)

		count++
		if count%100000 == 0 {
			fmt.Printf("Generated %d pages data\n", count)
		}
	}

	if err := <-errChan; err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	Output(pageData, "./data/pages.dat")

	fmt.Printf("Generated %d pages data\n", count)
	return nil
}
