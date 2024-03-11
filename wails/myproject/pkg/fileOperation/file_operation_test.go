package fileOperation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_getFileInfo(t *testing.T) {
	fO := FileOperation{}
	t.Run("retrieves file information of a file", func(t *testing.T) {
		filePath := "./testData/testFile.txt"

		info := fO.GetFileInfo(filePath)

		require.NotEmpty(t, info)
	})
	t.Run("err if file does not exist", func(t *testing.T) {
		filePath := "./testData/gibts_nicht.txt"

		info := fO.GetFileInfo(filePath)

		require.Empty(t, info)
	})
}

func TestExifResponse_ParseExifString(t *testing.T) {
	meta := ExifResponse{}
	testString := `PixelYDimension: 903 Orientation: 1 Software: "ILCE-7M4 v1.11" DateTime: "2023:03:29 11:49:43" DateTimeDigitized: "2023:03:28 10:28:11" ColorSpace: 1 PixelXDimension: 772 XResolution: "350/1" YResolution: "350/1" ExifIFDPointer: 138`
	meta.ParseExifString(testString)
	require.Equal(t, "772", meta.ExifInfo.ExifMeta.PixelXDimension)
	require.Equal(t, "903", meta.ExifInfo.ExifMeta.PixelYDimension)
	require.Equal(t, "1", meta.ExifInfo.ExifMeta.Orientation)
	require.Equal(t, "ILCE-7M4 v1.11", meta.ExifInfo.ExifMeta.Software)
	require.Equal(t, "2023:03:29 11:49:43", meta.ExifInfo.ExifMeta.DateTime)
	require.Equal(t, "2023:03:28 10:28:11", meta.ExifInfo.ExifMeta.DateTimeDigitized)
	require.Equal(t, "1", meta.ExifInfo.ExifMeta.ColorSpace)
	require.Equal(t, "350/1", meta.ExifInfo.ExifMeta.XResolution)
	require.Equal(t, "350/1", meta.ExifInfo.ExifMeta.YResolution)
	require.Equal(t, "138", meta.ExifInfo.ExifMeta.ExifIFDPointer)
}
