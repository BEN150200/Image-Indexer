package metadata

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"path/filepath"

	"os"

	"github.com/rwcarlsen/goexif/exif"
)

type ImageMetadata struct {
	Format       string
	FileName     string
	AbsolutePath string
	ExifRaw      []byte `json:"exif_raw,omitempty"`
}

type Reader interface {
	Read(filePath string) (*ImageMetadata, error)
}

type DefaultReader struct{}

func (r DefaultReader) Read(filePath string) (*ImageMetadata, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, format, err := image.DecodeConfig(file)
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	absolutePath, _ := filepath.Abs(filePath)
	metadata := ImageMetadata{
		Format:       format,
		FileName:     file.Name(),
		AbsolutePath: absolutePath,
	}

	// Read EXIF data
	x, err := exif.Decode(file)
	if err != nil && err != io.EOF {
		// Some images might not have EXIF, so ignore EOF for non-EXIF files
		fmt.Println("Warning: no EXIF data:", err)
	}

	if x != nil {
		metadata.ExifRaw = x.Raw
	}

	return &metadata, nil
}

func (m ImageMetadata) Debug() {
	fmt.Println("File: ", m.FileName)
	fmt.Println("Format:", m.Format)

	if m.AbsolutePath != "" {
		fmt.Println("Path: ", m.AbsolutePath)
	}
	if m.ExifRaw != nil {
		x, _ := exif.Decode(bytes.NewReader(m.ExifRaw))
		fmt.Println("Metadata: \n", x.String())
	}
}
