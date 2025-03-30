package core

import (
	"encoding/gob"
	"os"
)

type PageData struct {
	IDToTitle map[int32]string
	TitleToID map[string]int32
}

type PageLinksData struct {
	PageLinksMap        map[int32][]int32
	PageLinksMapReverse map[int32][]int32
}

func ReadPagesData(path string) (*PageData, error) {
	if path == "" {
		path = "/data/pages.dat"
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := &PageData{}
	dec := gob.NewDecoder(file)
	if err := dec.Decode(data); err != nil {
		panic(err)
	}
	return data, nil
}

func ReadLinkTargetsData(path string) (map[int32]string, error) {
	if path == "" {
		path = "/data/linkTargets.dat"
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data map[int32]string
	dec := gob.NewDecoder(file)
	if err := dec.Decode(&data); err != nil {
		panic(err)
	}
	return data, nil
}

func ReadPageLinksData(path string) (*PageLinksData, error) {
	if path == "" {
		path = "/data/pagelinks.dat"
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := &PageLinksData{}
	dec := gob.NewDecoder(file)
	if err := dec.Decode(data); err != nil {
		return nil, err
	}
	return data, nil
}
