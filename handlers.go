package main

import (
	"fmt"
	"indexer/metadata"
)

func AddImage(filename string) error {
	reader := metadata.DefaultReader{}

	meta, err := reader.Read(filename)

	if err != nil {
		fmt.Printf("%s Error: %s", filename, err.Error())
	}

	return Database.SaveData("images", meta.AbsolutePath, meta)
}

func ReadImage(filename string) error {
	var img metadata.ImageMetadata
	err := Database.GetData("images", filename, &img)
	if err != nil {
		return err
	}

	img.Debug()
	return nil
}
