package storage

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const fileName = "testFile_tests.txt"
const corruptFileName = "notExistantFolder/testFile_tests.txt"

func TestStorageWrite(t *testing.T) {
	t.Run("write should write expected text into the given file and response with successfull msg", func(t *testing.T) {
		store := helperCreateValidStorage()
		text := "hallo welt!"
		expected := fmt.Sprintf("%s\n",text)
		result := store.writeToFile(text)

		filecontent, err := os.ReadFile(fileName)
		require.NoError(t, err)

		require.Equal(t, FileWrittenResponse.string(), result)
		require.Equal(t, expected, string(filecontent))
	})

	t.Run("return err response on fail writing to file", func(t *testing.T) {
		
		store := helperCreateInvalidStorage()
		text := "hallo welt!"
		result := store.writeToFile(text)

		require.Contains(t, result, FileOperationFailed.string())
	})
}

func TestStorageRead(t *testing.T) {
	t.Run("Reads content w/o err", func(t *testing.T) {
		store := helperCreateValidStorage()
		text := "This is an example text."
		expectedText := fmt.Sprintf("%s\n", text)
		writeResponse := store.writeToFile(text)

		require.Equal(t, FileWrittenResponse.string(), writeResponse)

		readResponse := store.Read() 
		require.Equal(t, expectedText, readResponse)

	})

	t.Run("Returns '-1' if error occured", func(t *testing.T) {
		store := helperCreateInvalidStorage()
		
		readResponse := store.Read() 
		require.Equal(t, "-1", readResponse)

	})
}

func TestWrite(t *testing.T) {
	t.Run("should append text to file", func(t *testing.T) {
		store := helperCreateValidStorage()
		response := store.Write("hallo")
		require.Equal(t, FileWrittenResponse.string(), response)

		response = store.Write("welt")
		require.Equal(t, FileWrittenResponse.string(), response)
		require.Contains(t, store.Read(), "hallo\nwelt\n")
	})
}



func TestTruncate(t *testing.T) {
	t.Run("Empties the storage.", func(t *testing.T) {
		store := helperCreateValidStorage()
		store.Write("this should vanish.")
		store.Truncate()

		require.Empty(t, store.Read())
	})
}


func helperCreateValidStorage() *Storage {
	return NewStorage(fileName)
}

func helperCreateInvalidStorage() *Storage {
	return NewStorage(corruptFileName)
}