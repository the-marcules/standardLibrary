package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const bucketName = "mrcls-litter-bucket"

func downloadFile(dlUrl string) string {
	u, err := url.Parse(dlUrl)
	if err != nil {
		log.Fatalf("error on parisng url: %s", err.Error())
	}
	filename := u.Path
	downloadPath := "./download"
	out, err := os.Create(downloadPath + filename)
	if err != nil {
		log.Fatalf("error on creating local file output: %s", err.Error())
	}
	defer out.Close()

	resp, err := http.Get(dlUrl)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("error on downloading file: %s", err.Error())
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalf("error on writing dl to output file: %s", err.Error())
	}
	return downloadPath + filename
}

func getHash(filename string) string {
	//hashing
	/*	file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("error on opening file: %s", err.Error())
		}*/

	bytes2, _ := os.ReadFile(filename)

	fileHash := sha256.Sum256(bytes2)
	//if _, err = io.Copy(fileHash, file); err != nil {
	//	log.Fatal("error on hashing " + err.Error())
	//}
	buf := new(bytes.Buffer)
	enc := base64.NewEncoder(base64.StdEncoding, buf)
	enc.Write(fileHash[:])
	return fmt.Sprintf("%v", buf.String())
}

func initAWS() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"),
	})

	if err != nil {
		return nil, errors.New("error on initiating AWS Session")
	}
	return sess, nil
}

func GetPresignedUrl(sess *session.Session, filename string) string {
	svc := s3.New(sess)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	return urlStr
}

func GetS3Checksum(sess *session.Session, filename string) (checksum string) {
	svc := s3.New(sess)
	req, err := svc.GetObjectAttributes(&s3.GetObjectAttributesInput{
		Bucket:           aws.String(bucketName),
		Key:              aws.String(filename),
		ObjectAttributes: []*string{aws.String("x-amz-checksum-sha256")},
	})

	if err != nil {
		log.Fatalf("failed to get attribute request:%s", err)
	}

	log.Print(req.GoString())
	return req.Checksum.GoString()
}

func putFileToS3(sess *session.Session, filename string) {
	svc := s3.New(sess)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		//Body:   strings.NewReader("EXPECTED CONTENTS"),
	})

	str, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Fatalf("error getting presigned put url: %s", err.Error())
	} else {
		log.Printf("url for uploading %s:\n%s", filename, str)
	}

}

func main() {
	sess, err := initAWS()
	if err != nil {
		log.Fatal(err.Error())
	}

	//putFileToS3(sess, "testfile.txt")

	//checksum := GetS3Checksum(sess, "testfile.txt")
	//print(checksum)

	urlStr := GetPresignedUrl(sess, "testfile.txt")

	log.Println("The URL is", urlStr)

	downloadedFile := downloadFile(urlStr)

	fileHash := getHash(downloadedFile)
	fmt.Printf("File %s has Hash (SHA256) of '%s'\n", downloadedFile, fileHash)

}
