package upload

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"de.marcules.httptesting/util"
)

const tmpDir = "tmp"

type FileMeta struct {
	Filename string `json:"filename"`
	Hash     string `json:"hash"`
}

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {

	fileMeta, statusCode, err := createFileFromMultipart(r)
	if err != nil {
		slog.Error("could not create file from multipart form", "error", err.Error())
		w.WriteHeader(statusCode)
		return
	}

	statusCode, err = checkFileHash(fileMeta)
	if err != nil {
		slog.Error(err.Error())
		w.WriteHeader(statusCode)
		return
	}

	w.WriteHeader(statusCode)
}

func checkFileHash(fileMeta *FileMeta) (int, error) {
	fileHash, err := util.HashFile(tmpDir + "/" + fileMeta.Filename)
	if err != nil {
		err = errors.New("could not hash file. error:" + err.Error())
		return http.StatusInternalServerError, err

	}

	if fileMeta.Hash != fileHash {
		_ = os.Remove(tmpDir + "/" + fileMeta.Filename)
		err = errors.New("file hashes do not match")
		return http.StatusBadRequest, err

	}

	return http.StatusOK, nil
}

func createFileFromMultipart(r *http.Request) (*FileMeta, int, error) {
	err := r.ParseMultipartForm(1000 * 1000 * 400)
	if err != nil {
		err = errors.New(fmt.Sprintf("could not parse multipart form. %s: %s", "error", err.Error()))
		return nil, http.StatusBadRequest, err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		err = errors.New(fmt.Sprintf("could not get file from form. %s: %s", "error", err.Error()))
		return nil, http.StatusBadRequest, err
	}
	defer file.Close()

	fileMetaStr := r.FormValue("fileMeta")

	var fileMeta FileMeta
	err = json.Unmarshal([]byte(fileMetaStr), &fileMeta)
	if err != nil {
		err = errors.New(fmt.Sprintf("could not parse file meta data. %s: %s", "error", err.Error()))
		return nil, http.StatusInternalServerError, err
	}

	err = os.MkdirAll(tmpDir, 0o755)
	if err != nil {
		err = errors.New(fmt.Sprintf("could not create temp dir. %s: %s", "error", err.Error()))
		return nil, http.StatusInternalServerError, err
	}

	tmpFile, err := os.CreateTemp(tmpDir, "uploaded-*")
	if err != nil {
		err = errors.New(fmt.Sprintf("could not create temp file. %s: %s", "error", err.Error()))
		return nil, http.StatusInternalServerError, err
	}
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, file)
	defer r.Body.Close()
	if err != nil {
		err = errors.New(fmt.Sprintf("could not write body to tmp file. %s: %s", "error", err.Error()))
		return nil, http.StatusInternalServerError, err
	}

	err = os.Rename(tmpFile.Name(), tmpDir+"/"+fileMeta.Filename)
	if err != nil {
		err = errors.New(fmt.Sprintf("could not rename file. %s: %s", "error", err.Error()))
		return nil, http.StatusInternalServerError, err
	}

	return &fileMeta, http.StatusOK, nil
}
