package main

import (
	"fmt"
	"strconv"
)

// Generate linkTargetsMap from jawiki-20250301-linktarget.sql
// Columns: lt_id, lt_namespace, lt_title
// output: linkTargetsMap := make(map[int32]string)
// data type: lt_id: int32, lt_title: string
func GenerateLinkTargetsData() error {
	valsChan, errChan := SQLValueRowIterator("./sql/jawiki-20250301-linktarget.sql")

	count := 0

	linkTargetsMap := make(map[int32]string)

	for vals := range valsChan {
		lt_id, err := strconv.ParseInt(vals[0], 10, 32)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		lt_namespace, err := strconv.ParseInt(vals[1], 10, 32)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		lt_title := vals[2]

		if lt_namespace != 0 {
			continue
		}

		linkTargetsMap[int32(lt_id)] = lt_title

		count++
		if count%100000 == 0 {
			fmt.Printf("Generated %d link targets data\n", count)
		}
	}

	fmt.Println("Finished reading valsChan")

	if err := <-errChan; err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	Output(linkTargetsMap, "./data/linkTargets.dat")

	fmt.Printf("Generated %d link targets data\n", count)
	return nil
}
