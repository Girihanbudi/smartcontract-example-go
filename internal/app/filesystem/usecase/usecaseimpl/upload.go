package usecaseimpl

import (
	"context"
	"encoding/json"
	"fmt"
	goio "io"
	"os"
	"regexp"
	module "smartm2m-technical-test/internal/app/filesystem"
	"smartm2m-technical-test/internal/app/filesystem/preset/request"
	"smartm2m-technical-test/internal/app/filesystem/preset/response"
	"smartm2m-technical-test/internal/pkg/io"
	"smartm2m-technical-test/internal/pkg/stderror"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (u Usecase) Upload(ctx context.Context, cmd request.Upload) (res response.Upload, err *stderror.StdError) {
	// Validate input command
	if _, err = cmd.Validate(); err != nil {
		return
	}

	// Generate random nano id
	tempPath, generateIdErr := gonanoid.New()
	if generateIdErr != nil {
		err = stderror.ErrInternalServerError().SetError(generateIdErr)
		return
	}
	tempPath = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(tempPath, "")

	// Create a temporary directory for calculation
	path := fmt.Sprintf("tmp/%s", tempPath)
	if createDirErr := os.MkdirAll(path, os.ModePerm); createDirErr != nil {
		err = stderror.ErrInternalServerError().SetError(createDirErr)
		return
	}

	// Delete temporary folder to the closest return
	defer os.RemoveAll(path)

	// Extract zip to temporary folder
	openedFile, openFileErr := cmd.File.Open()
	if openFileErr != nil {
		err = stderror.ErrInternalServerError().SetError(openFileErr)
		return
	}

	if extractErr := io.ExtractTarGz(openedFile, path); err != nil {
		err = stderror.ErrBadRequest().SetError(extractErr)
		return
	}

	// Read all content from extracted zip
	var files []module.File
	fInfos, _ := os.ReadDir(path)
	for _, f := range fInfos {
		b, readFileErr := os.ReadFile(fmt.Sprintf("%s/%s", path, f.Name()))
		if readFileErr != nil {
			err = stderror.ErrInternalServerError().SetError(readFileErr)
			return
		}
		var content map[string]interface{}
		if convertErr := json.Unmarshal(b, &content); convertErr != nil {
			err = stderror.ErrInternalServerError().SetError(convertErr)
			return
		}
		files = append(files, module.File{
			Filename: f.Name(),
			Contents: content,
		})
	}

	// Store file to smart contract
	trxOpts, conn, err := u.GetTrxOptsAndConn(ctx, cmd.Address)
	if err != nil {
		return
	}

	openedFile, openFileErr = cmd.File.Open()
	if openFileErr != nil {
		err = stderror.ErrInternalServerError().SetError(openFileErr)
		return
	}

	content, readFileErr := goio.ReadAll(openedFile)
	if readFileErr != nil {
		err = stderror.ErrInternalServerError().SetError(readFileErr)
		return
	}

	_, uploadErr := conn.AddFile(trxOpts, cmd.File.Filename, string(content))
	if uploadErr != nil {
		err = stderror.ErrInternalServerError().SetError(uploadErr)
		return
	}

	// Assign value to response
	res.Files = files

	return
}
