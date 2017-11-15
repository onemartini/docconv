// +build ocr

package main

import (
	"fmt"
	"io"
)

func ConvertPDF(r io.Reader) (string, map[string]string, error) {

	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return bodyResult, metaResult, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	// Verify if pdf has images or is pdf only-text
	if PDFHasImage(f.Name()) {
		bodyResult, err := ConvertImagePDF(f.Name())
		return bodyResult.body, nil, nil
	}
	bodyResult, metaResult, err := ConvertTextPDF(f.Name())
	return bodyResult.body, metaResult.meta, nil

}
