package io

import (
	"testing"
)

func TestExtensionExist(t *testing.T) {
	fileName := "my-file.tar.gz"
	if exist := ExtensionExist(fileName, ".tar.gz"); !exist {
		t.Error("Validation failed on checking same extension")
	}

	fileName = "my-file.pdf"
	if exist := ExtensionExist(fileName, ".tar.gz"); exist {
		t.Error("Validation failed on checking diferent extension")
	}

	fileName = ".tar.gz"
	if exist := ExtensionExist(fileName, ".tar.gz"); exist {
		t.Error("Validation failed on checking no file name")
	}
}
