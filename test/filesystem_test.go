package test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	file := "sample-file.tar.gz"
	addr := "22aabb811efca4e6f4748bd18a46b502fa85549df9fa07da649c0a148d7d5530"
	b, w := createMultipartFormData(
		t,
		MultipartField{Key: "file", Type: File, FilePath: file},
		MultipartField{Key: "address", Type: String, Str: addr},
	)

	req, err := http.NewRequest("POST", "/filesystem", &b)
	if err != nil {
		t.Errorf("Error while uploading file: %v", err)
	}

	req.Header.Set("Content-Type", w.FormDataContentType())
}

func TestDownload(t *testing.T) {
	addr := "22aabb811efca4e6f4748bd18a46b502fa85549df9fa07da649c0a148d7d5530"
	res, err := http.NewRequest("GET", fmt.Sprintf("/filesystem?address=%s", addr), nil)
	if err != nil {
		t.Errorf("Error while downloading file: %v", err)
	}

	out, err := os.Create("sample-file2.tar.gz")
	if err != nil {
		t.Errorf("Error while creating a file: %v", err)
	}

	defer out.Close()
	_, err = io.Copy(out, res.Body)
	if err != nil {
		t.Errorf("Error while copying file: %v", err)
	}
}
