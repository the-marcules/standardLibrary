package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

var PORT = "8080"

func Test_rootEndPoint(t *testing.T) {
	t.Run("should retrun status of 200 and body of 'hello, world'", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:"+PORT, nil)
		recorder := httptest.NewRecorder()
		RootHandler(recorder, req)

		expectedStatusCode := 200

		require.Equal(t, expectedStatusCode, recorder.Code)
		require.Equal(t, "Hello, World!", recorder.Body.String())

	})
}
