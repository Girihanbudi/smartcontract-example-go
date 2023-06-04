package test

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"testing"
)

type MultipartType int

const (
	File MultipartType = iota
	String
)

type MultipartField struct {
	Key      string
	Type     MultipartType
	Str      string
	FilePath string
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Println("PWD: ", pwd)
		panic(err)
	}
	return r
}

func createMultipartFormData(t *testing.T, fields ...MultipartField) (bytes.Buffer, *multipart.Writer) {
	var b bytes.Buffer
	var err error
	w := multipart.NewWriter(&b)
	var fw io.Writer

	for _, field := range fields {
		if field.Type == File {
			file := mustOpen(field.FilePath)
			if fw, err = w.CreateFormFile(field.Key, file.Name()); err != nil {
				t.Errorf("Error creating writer: %v", err)
			}
			if _, err = io.Copy(fw, file); err != nil {
				t.Errorf("Error with io.Copy: %v", err)
			}
		} else {
			if err = w.WriteField(field.Key, field.Str); err != nil {
				t.Errorf("Error write field: %v", err)
			}
		}
	}

	w.Close()
	return b, w
}
