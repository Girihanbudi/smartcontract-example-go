package io

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func ExtractTarGz(gzipStream io.Reader, parentPath string) error {
	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		return err
	}

	tarReader := tar.NewReader(uncompressedStream)
	for header, err := tarReader.Next(); err == nil; header, err = tarReader.Next() {
		path := fmt.Sprintf("%s/%s", parentPath, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(path, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
		case tar.TypeReg:
			outFile, err := os.Create(path)
			if err != nil {
				return fmt.Errorf("failed create file: %w", err)
			}

			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				return fmt.Errorf("failed to copy file: %w", err)
			}
			if err := outFile.Close(); err != nil {
				return fmt.Errorf("failed to close file: %w", err)
			}
		default:
			return fmt.Errorf("unknown type: %b in %s", header.Typeflag, header.Name)
		}
	}
	if err != io.EOF {
		return fmt.Errorf("header not found: %w", err)
	}
	return nil
}

func ExtensionExist(filename string, extensions ...string) bool {
	fileLength := len(filename)
	for _, ext := range extensions {
		extLength := len(ext)
		if fileLength > extLength {
			fileExt := filename[fileLength-extLength:]
			if strings.ToLower(fileExt) == strings.ToLower(ext) {
				return true
			}
		}
	}

	return false
}
