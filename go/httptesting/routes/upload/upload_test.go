package upload

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"de.marcules.httptesting/util"
	"github.com/stretchr/testify/require"
)

var PORT = "8080"

const (
	TestFilesPath = "../../testdata"
	TmpDir        = "./tmp/"
)

func Test_upload(t *testing.T) {
	t.Run("should return status of 200 when file is uploaded", func(t *testing.T) {
		fileName := "file-to-upload.txt"
		body, contentType := getMultipartRequestBody(t, fileName, "")

		req := httptest.NewRequest("PUT", "http://localhost:"+PORT+"/upload", body)
		req.Header.Set("Content-Type", contentType)
		recorder := httptest.NewRecorder()
		FileUploadHandler(recorder, req)
		require.Equal(t, http.StatusOK, recorder.Code)
		require.FileExists(t, "./tmp/file-to-upload.txt")
	})

	t.Run("should return status of 400 if file hashes do not match and file is deleted", func(t *testing.T) {
		fileName := "file-to-upload.txt"
		body, contentType := getMultipartRequestBody(t, fileName, "invalidHash")

		req := httptest.NewRequest("PUT", "http://localhost:"+PORT+"/upload", body)
		req.Header.Set("Content-Type", contentType)
		recorder := httptest.NewRecorder()
		FileUploadHandler(recorder, req)

		require.Equal(t, http.StatusBadRequest, recorder.Code)
		require.NoFileExists(t, TmpDir+fileName)
	})

	t.Run("should return status of 400 if multipart form is invalid", func(t *testing.T) {
		fileName := "file-to-upload.txt"
		body := bytes.NewBufferString("invalid multipart form data")
		req := httptest.NewRequest("PUT", "http://localhost:"+PORT+"/upload", body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary=invalidBoundary")
		recorder := httptest.NewRecorder()
		FileUploadHandler(recorder, req)

		require.Equal(t, http.StatusBadRequest, recorder.Code)
		require.NoFileExists(t, TmpDir+fileName)
	})

	t.Run("should return status of 400 if file is invalid in multipart form", func(t *testing.T) {
		fileName := "file-to-upload.txt"
		body := bytes.NewBufferString(`--585019cbf5086faf39ac381dc42beb36ced90dccddc71c06fa21f1d09ca0
Content-Disposition: form-data; name="fileMeta"

{"filename":"file-to-upload.txt","hash":"ead823998779483cf26cfae2d4ab62e1f5954d07f36e99196dbb3c106e9ae76a"}
--585019cbf5086faf39ac381dc42beb36ced90dccddc71c06fa21f1d09ca0--
`)
		req := httptest.NewRequest("PUT", "http://localhost:"+PORT+"/upload", body)
		req.Header.Set("Content-Type", "multipart/form-data; boundary=585019cbf5086faf39ac381dc42beb36ced90dccddc71c06fa21f1d09ca0")
		recorder := httptest.NewRecorder()
		FileUploadHandler(recorder, req)

		require.Equal(t, http.StatusBadRequest, recorder.Code)
		require.NoFileExists(t, TmpDir+fileName)
	})

	t.Cleanup(func() {
		err := os.RemoveAll(TmpDir)
		if err != nil {
			t.Error("could not remove uploaded file", "error", err.Error())
		}
	})

}

func getMultipartRequestBody(t *testing.T, fileName string, hashOverWrite string) (*bytes.Buffer, string) {

	pathAndFilename := TestFilesPath + "/" + fileName
	fileToUpload, err := os.Open(pathAndFilename)
	require.NoError(t, err)

	fileHash := hashOverWrite

	if hashOverWrite == "" {
		fileHash, err = util.HashFile(pathAndFilename)
		require.NoError(t, err)
	}

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	defer writer.Close()

	part, err := writer.CreateFormFile("file", fileName)
	require.NoError(t, err)
	_, err = io.Copy(part, fileToUpload)
	require.NoError(t, err)

	fileMeta := FileMeta{
		Filename: fileName,
		Hash:     fileHash,
	}

	metaString, err := json.Marshal(fileMeta)
	require.NoError(t, err)

	err = writer.WriteField("fileMeta", string(metaString))
	require.NoError(t, err)

	return &body, writer.FormDataContentType()
}
