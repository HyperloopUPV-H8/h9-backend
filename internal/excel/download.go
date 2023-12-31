package excel

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	trace "github.com/rs/zerolog/log"
	"github.com/xuri/excelize/v2"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"

	_ "embed"
)

const SHEETS_MIME_TYPE = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

var (
	ErrClient   = "getting client"
	ErrDownload = "getting file"
	ErrSaving   = "saving file"
	ErrOpening  = "opening file"
)

type DownloadConfig struct {
	Id   string
	Path string
	Name string
}

//go:embed secret.json
var apiKey []byte

func Download(config DownloadConfig) (*excelize.File, error) {
	trace.Trace().Str("id", config.Id).Str("path", config.Path).Str("name", config.Name).Msg("download file")
	client, errClient := getClient(apiKey)
	if errClient != nil {
		trace.Error().Str("id", config.Id).Str("path", config.Path).Str("name", config.Name).Stack().Err(errClient).Msg("")
		return nil, errors.Wrap(errClient, ErrClient)
	}

	fileBuf, errFile := getFile(client, config.Id, SHEETS_MIME_TYPE)
	if errFile != nil {
		trace.Warn().Str("id", config.Id).Str("path", config.Path).Str("name", config.Name).Stack().Err(errFile).Msg("")
	} else {
		errSaving := saveFile(fileBuf, config.Path, config.Name)
		if errSaving != nil {
			trace.Error().Str("id", config.Id).Str("path", config.Path).Str("name", config.Name).Stack().Err(errSaving).Msg("")
			return nil, errSaving
		}
	}

	file, errOpening := excelize.OpenFile(filepath.Join(config.Path, config.Name))
	if errOpening != nil {
		return nil, errors.Wrap(errOpening, ErrOpening)
	}

	return file, nil
}

func getClient(apiKey []byte) (*drive.Service, error) {
	trace.Trace().Msg("get client")
	ctx := context.Background()

	client, err := drive.NewService(ctx, option.WithCredentialsJSON(apiKey))

	if err != nil {
		return nil, err
	}

	return client, nil
}

func getFile(client *drive.Service, id string, mimeType string) ([]byte, error) {
	trace.Trace().Str("id", id).Msg("get file")
	resp, err := client.Files.Export(id, mimeType).Download()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func saveFile(content []byte, path string, name string) error {
	trace.Trace().Str("path", path).Str("name", name).Msg("save file")

	err := os.Mkdir(path, 0777)

	if !os.IsExist(err) {
		return err
	}

	err = os.WriteFile(filepath.Join(path, name), content, 0644) // rw-r--r--
	if err != nil {
		return err
	}
	return nil
}
