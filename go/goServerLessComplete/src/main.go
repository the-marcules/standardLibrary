package main

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type App struct {
	ctx        context.Context
	s3Client   PutterApi
	bucketName string
}

type PutterApi interface {
	PutObject(context.Context, *s3.PutObjectInput, ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

func (a App) HandleRequest() error {

	fileContent := "Hallo Welt!"
	reader := strings.NewReader(fileContent)

	params := &s3.PutObjectInput{
		Bucket: aws.String(a.bucketName),
		Key:    aws.String("myFile.txt"),
		Body:   reader,
	}

	_, err := a.s3Client.PutObject(a.ctx, params)
	if err != nil {
		slog.Error("could not put object %s", "Error", err.Error())
	}
	return err
}

func main() {
	bucketName := os.Getenv("BUCKET_NAME")

	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		slog.Error("got error initiating default config %s", "Error", err.Error())
	}

	s3Client := s3.NewFromConfig(cfg)

	app := App{
		ctx:        ctx,
		s3Client:   s3Client,
		bucketName: bucketName,
	}

	app.HandleRequest()

}
