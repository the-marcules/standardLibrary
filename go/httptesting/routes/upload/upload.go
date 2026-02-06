package upload

import (
	"encoding/json"
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

	err := r.ParseMultipartForm(1000 * 1000 * 400)
	if err != nil {
		slog.Error("could not parse multipart form", "error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		slog.Error("could not get file from form", "error", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	wert := r.FormValue("fileMeta")

	var fileMeta FileMeta
	err = json.Unmarshal([]byte(wert), &fileMeta)
	if err != nil {
		slog.Error("could not parse file meta data", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = os.MkdirAll(tmpDir, 0o755)
	if err != nil {
		slog.Error("could not create temp dir", "error", err.Error(), "tmpDir", tmpDir)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tmpFile, err := os.CreateTemp(tmpDir, "uploaded-*")
	defer tmpFile.Close()
	if err != nil {
		slog.Error("could not create temp file", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(tmpFile, file)
	defer r.Body.Close()
	if err != nil {
		slog.Error("could not write body to tmp file", "error", err.Error(), "tmpFile", tmpDir+"/"+tmpFile.Name())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = os.Rename(tmpFile.Name(), tmpDir+"/"+fileMeta.Filename)
	if err != nil {
		slog.Error("could not rename file", "error", err.Error(), "tmpFile", tmpDir+"/"+tmpFile.Name(), "targetFile", tmpDir+"/"+fileMeta.Filename)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fileHash, err := util.HashFile(tmpDir + "/" + fileMeta.Filename)
	if err != nil {
		slog.Error("could not hash file", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if fileMeta.Hash != fileHash {
		_ = os.Remove(tmpDir + "/" + fileMeta.Filename)

		slog.Error("file hashes do not match")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func createFileFromMultipart(r *http.Request) (FileMeta, error) {

}
