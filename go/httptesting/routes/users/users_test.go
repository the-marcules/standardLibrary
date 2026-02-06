package users

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

var PORT = "8080"

var expectedBodyJson = Users{
	Users: []User{
		{
			ID:    1,
			Name:  "Alice",
			Email: "alice@wonderland.wo",
		},
		{
			ID:    2,
			Name:  "Bob",
			Email: "bob@bobs-burgers.com",
		},
	}}

func Test_rootEndPoint(t *testing.T) {
	t.Run("should retrun status of 200 and body of json form user list", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:"+PORT+"/users", nil)
		recorder := httptest.NewRecorder()
		t.Setenv("USER_FILE", "../../testdata/users.json")
		GetUsers(recorder, req)

		expectedStatusCode := 200
		var gotBodyJson Users
		err := json.NewDecoder(recorder.Body).Decode(&gotBodyJson)
		require.NoError(t, err)
		require.Equal(t, expectedStatusCode, recorder.Code)
		require.Equal(t, gotBodyJson, expectedBodyJson)

	})
	t.Run("should return 500 and empty body if no user file is set", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:"+PORT+"/users", nil)
		recorder := httptest.NewRecorder()
		t.Setenv("USER_FILE", "../../testdata/non-existing-file.json")
		GetUsers(recorder, req)

		expectedStatusCode := 500
		require.Equal(t, expectedStatusCode, recorder.Code)
		require.Equal(t, recorder.Body.String(), "")
	})

	t.Run("should return 500 and empty body if user does not contain parseable json", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://localhost:"+PORT+"/users", nil)
		recorder := httptest.NewRecorder()
		t.Setenv("USER_FILE", "../../testdata/corrupt-users.json")
		GetUsers(recorder, req)

		expectedStatusCode := 500
		require.Equal(t, expectedStatusCode, recorder.Code)
		require.Equal(t, recorder.Body.String(), "")
	})
}
