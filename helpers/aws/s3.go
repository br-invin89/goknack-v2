package aws

import (
	"bytes"
	"fmt"
	//"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	//"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	maxPartSize        = int64(512 * 1000)
	maxRetries         = 3
	awsAccessKeyID     = "AKIAJX7U6FORX4IB5KXQ"
	awsSecretAccessKey = "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"
	awsBucketRegion    = "us-east-1"
	awsBucketName      = "goknack2"
)

// SVC var
var SVC *s3.S3

var Path string

func Upload(file *os.File, fileType string) (string, error){
	creds := credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, "")
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
	}
	cfg := aws.NewConfig().WithRegion(awsBucketRegion).WithCredentials(creds)
	SVC = s3.New(session.New(), cfg)

	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)

	if fileType == "audio/mpeg3" || fileType == "audio/mp3" || fileType == "audio/mpeg" {
		Path = "/content/audio/" + file.Name()
	} else {
		Path = "/content/image/" + file.Name()
	}
	
	params := &s3.PutObjectInput{
		Bucket:      aws.String(awsBucketName),
		Key:         aws.String(Path),
		ContentType: aws.String(fileType),
		Body: fileBytes,
		ContentLength: aws.Int64(size),
		ACL: aws.String("public-read"),
	}

	_, err = SVC.PutObject(params) 
	if err != nil { 
	  return "", err
	} 

	return Path, nil
}
