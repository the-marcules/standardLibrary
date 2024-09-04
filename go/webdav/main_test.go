package main

import (
	"context"
	"encoding/json"
	"fmt"
	webdav "github.com/emersion/go-webdav"
	"github.com/stretchr/testify/require"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func TestWebDav(t *testing.T) {

	t.Run("basic getstat should not return an error", func(t *testing.T) {

		server := SetupFakeDtServer(t)
		webDavClient := NewWebDavCommunicator(server.URL)
		defer server.Close()
		response := webDavClient.GetStat("adam")
		require.NoError(t, response.Err)
	})

	t.Run("massive count of parallel requests should not fail", func(t *testing.T) {

		server := SetupFakeDtServer(t)
		defer server.Close()
		webDavClient := NewWebDavCommunicator(server.URL)
		sampleSize := 10000

		testResult := make(chan *WebDavResponse, sampleSize)

		fmt.Printf("initiate %d requests\n", sampleSize)
		for i := 0; i < sampleSize; i++ {
			go func(run int) {
				testResult <- webDavClient.GetStat("adam")
			}(i)
		}

		println("evaluating results")
		for i := 0; i < sampleSize; i++ {
			response := <-testResult
			msg := fmt.Sprintf("failed in run %d", i)
			require.NoError(t, response.Err, msg)
			println(response.Msg)
		}

		println("finished evaluating results")

	})

}

func TestContext(t *testing.T) {
	server := SetupFakeDtServer(t)
	defer server.Close()
	webDavClient := NewWebDavCommunicator(server.URL)
	t.Run("cancel request", func(t *testing.T) {

		response := make(chan *WebDavResponse, 1)
		go func() {
			response <- webDavClient.GetStat("adam")

		}()
		webDavClient.cancel()
		evaluate := <-response
		require.Error(t, evaluate.Err)
		require.Contains(t, evaluate.Err.Error(), "context canceled")
	})
}

func TestNewWebDavCommunicator(t *testing.T) {
	t.Run("should not panic with transport", func(t *testing.T) {
		server := SetupFakeDtServer(t)
		defer server.Close()

		require.NotPanics(t, func() {
			_ = NewWebDavCommunicator(server.URL)
		})
	})

	t.Run("cancel manually", func(t *testing.T) {
		server := SetupFakeDtServer(t)
		defer server.Close()

		webdavClient := NewWebDavCommunicator(server.URL)
		response := make(chan *WebDavResponse, 1)
		go func() {
			response <- webdavClient.GetStat("adam")

		}()
		require.NotPanics(t, func() {
			webdavClient.cancel()
		})
		evaluate := <-response
		require.Error(t, evaluate.Err)
		require.Contains(t, evaluate.Err.Error(), "context canceled")
	})

	t.Run("should timeout", func(t *testing.T) {
		server := SetupFakeDtServer(t)
		defer server.Close()

		webdavClient := NewWebDavCommunicator(server.URL, 1*time.Millisecond)
		response := make(chan *WebDavResponse, 1)
		go func() {
			response <- webdavClient.GetStat("adam")

		}()

		evaluate := <-response
		require.Error(t, evaluate.Err)
		require.Contains(t, evaluate.Err.Error(), "deadline exceeded")
	})
}

func TestWebDavResponse_PutFile(t *testing.T) {
	defer cleanUp(t)
	t.Run("put file should not return an error", func(t *testing.T) {
		server := SetupFakeDtServer(t)
		webDavClient := NewWebDavCommunicator(server.URL)

		startTime := time.Now().Unix()
		fmt.Printf("starttime %d", startTime)
		wg := sync.WaitGroup{}
		batch := 5000

		var i int
		for i = 0; i < batch; i++ {
			wg.Add(1)
			go func(count int) {
				defer wg.Done()
				remoteFileName := fmt.Sprintf("testdata/testfile_%d.txt", count)

				response := webDavClient.PutFile("testdata/testfile.txt", remoteFileName)
				require.NoError(t, response.Err, "got error on putting but should not")

				response = webDavClient.GetStat(remoteFileName)
				require.NoError(t, response.Err, "should not return an error, on finding file")

				var stat webdav.FileInfo
				err := json.Unmarshal([]byte(response.Msg), &stat)
				require.NoError(t, err)

				require.GreaterOrEqual(t, stat.ModTime.Unix(), startTime)

			}(i)

		}
		wg.Wait()
		t.Logf("finished %d runs of %d", i, batch)
	})

	t.Run("put file should not return an error", func(t *testing.T) {
		server := SetupFakeDtServer(t)
		webDavClient := NewWebDavCommunicator(server.URL)

		startTime := time.Now().Unix()
		fmt.Printf("starttime %d", startTime)

		remoteFileName := "testdata/C8521DB.zip")

		response := webDavClient.PutFile("testdata/C8521DB.zip", remoteFileName)
		require.NoError(t, response.Err, "got error on putting but should not")

		response = webDavClient.GetStat(remoteFileName)
		require.NoError(t, response.Err, "should not return an error, on finding file")

		var stat webdav.FileInfo
		err := json.Unmarshal([]byte(response.Msg), &stat)
		require.NoError(t, err)

		require.GreaterOrEqual(t, stat.ModTime.Unix(), startTime)

	})
}

func SetupFakeDtServer(t *testing.T) *httptest.Server {

	t.Helper()

	fs := webdav.LocalFileSystem("testdata/webdavdir")

	err := fs.Mkdir(context.TODO(), "/testdata")
	if err != nil {
		t.Logf("got error on creating webdavdir: %v", err)
	}
	err = fs.Mkdir(context.TODO(), "/adam")
	if err != nil {
		t.Logf("got error on creating webdavdir: %v", err)
	}
	handler := &webdav.Handler{
		FileSystem: fs,
	}

	server := httptest.NewServer(handler)
	return server

}

func cleanUp(t *testing.T) {
	fs := webdav.LocalFileSystem("testdata/webdavdir")
	_ = fs.RemoveAll(context.TODO(), "/testdata")
	_ = fs.RemoveAll(context.TODO(), "/adam")
}
