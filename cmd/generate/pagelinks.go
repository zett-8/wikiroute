package main

import (
	"fmt"
	"strconv"
	"wikiroute/core"
)

// Generate pageLinksMap 
// requires: jawiki-20250301-pagelinks.sql, linkTargets.dat, pages.dat
// Columns: pl_from, pl_from_namespace, pl_target_id
// output: PageLinksData
// data type: pl_from: int64, pl_target_id: int64
func GeneratePageLinksData() error {
	linkTargets, err := core.ReadLinkTargetsData("./data/linkTargets.dat")
	if err != nil {
		return err
	}
	pages, err := core.ReadPagesData("./data/pages.dat")
	if err != nil {
		return err
	}

	valsChan, errChan := SQLValueRowIterator("./sql/jawiki-20250301-pagelinks.sql")

	data := &core.PageLinksData{
		PageLinksMap:        make(map[int32][]int32),
		PageLinksMapReverse: make(map[int32][]int32),
	}

	count := 0
	for vals := range valsChan {
		_pl_namespace, err := strconv.ParseInt(vals[1], 10, 64)
		if err != nil {
			return err
		}
		pl_namespace := int32(_pl_namespace)

		if pl_namespace != 0 {
			continue
		}

		_pl_from, err := strconv.ParseInt(vals[0], 10, 64)
		if err != nil {
			return err
		}
		pl_from := int32(_pl_from)

		_pl_target_id, err := strconv.ParseInt(vals[2], 10, 64)
		if err != nil {
			return err
		}
		pl_target_id := int32(_pl_target_id)

		targetPageId := pages.TitleToID[linkTargets[pl_target_id]]

		data.PageLinksMap[pl_from] = append(data.PageLinksMap[pl_from], targetPageId)
		data.PageLinksMapReverse[targetPageId] = append(data.PageLinksMapReverse[targetPageId], pl_from)

		count++
		if count%1000000 == 0 {
			fmt.Printf("Generated %d page link data\n", count)
		}
	}

	if err := <-errChan; err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	if err := Output(data, "./data/pagelinks.dat"); err != nil {
		return err
	}

	fmt.Printf("Generated %d page links data\n", count)
	return nil
}
