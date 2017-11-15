// +build ocr

package docconv

import (
	"fmt"
	"io"
)

func ConvertPDF(r io.Reader) (string, map[string]string, error) {

	f, err := NewLocalFile(r, "/tmp", "sajari-convert-")
	if err != nil {
		return "", nil, fmt.Errorf("error creating local file: %v", err)
	}
	defer f.Done()

	// Verify if pdf has images or is pdf only-text
	if PDFHasImage(f.Name()) {
		bodyResult, imageConvertErr := ConvertImagePDF(f.Name())
		if imageConvertErr != nil {
			return "", nil, imageConvertErr
		}
		return bodyResult.body, nil, nil
	}
	bodyResult, metaResult, textConvertErr := ConvertTextPDF(f.Name())
	if textConvertErr != nil {
		return "", nil, imageConvertErr
	}
	return bodyResult.body, metaResult.meta, nil

}
