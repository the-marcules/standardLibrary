package main

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_HandleRequest(t *testing.T) {
	t.Run("should put file with a name to expected bucket", func(t *testing.T) {
		s3Mock := &S3MOCK{}
		expectedBucketName := "schorsch"
		ctx := context.TODO()
		app := App{
			s3Client:   s3Mock,
			ctx:        ctx,
			bucketName: expectedBucketName,
		}

		app.HandleRequest()

		assert.Equal(t, expectedBucketName, s3Mock.GivenBucket)
		assert.NotEmpty(t, s3Mock.GivenFile)
	})

	t.Run("should write content to file in bucket", func(t *testing.T) {
		s3Mock := &S3MOCK{}
		expectedBucketName := "schorsch"
		expectedFileContent := "Hallo Welt!"
		ctx := context.TODO()
		app := App{
			s3Client:   s3Mock,
			ctx:        ctx,
			bucketName: expectedBucketName,
		}

		err := app.HandleRequest()
		require.NoError(t, err)
		require.Equal(t, expectedFileContent, s3Mock.GivenContent)
	})

	t.Run("should return error", func(t *testing.T) {
		s3Mock := &S3MOCK{
			ForceError: true,
		}
		expectedBucketName := "schorsch"
		ctx := context.TODO()
		app := App{
			s3Client:   s3Mock,
			ctx:        ctx,
			bucketName: expectedBucketName,
		}

		err := app.HandleRequest()
		require.Error(t, err)
		require.EqualErrorf(t, err, "some err", "%s", "formatted")
	})
}

type S3MOCK struct {
	GivenBucket  string
	GivenFile    string
	GivenContent string
	ForceError   bool
}

func (s3 *S3MOCK) PutObject(ctx context.Context, inputParams *s3.PutObjectInput, options ...func(*s3.Options)) (out *s3.PutObjectOutput, err error) {

	if s3.ForceError {
		err = errors.New("some err")
		return
	}

	buffer, err := io.ReadAll(inputParams.Body)
	if err != nil {
		return
	}

	bodyString := string(buffer)

	s3.GivenBucket = *inputParams.Bucket
	s3.GivenFile = *inputParams.Key
	s3.GivenContent = string(bodyString)

	return
}
