package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
	"os"
)

type S3Config struct {
	AccessKeyId     string
	SecretAccessKey string
	Region          string
	BucketName      string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	sc := S3Config{
		AccessKeyId:     os.Getenv("ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("SECRET_ACCESS_KEY"),
		Region:          os.Getenv("REGION"),
		BucketName:      os.Getenv("BUCKET_NAME"),
	}

	sc.UploadFile("./simple.svg")
}

func (sc *S3Config) UploadFile(filepath string) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(sc.Region),
		Credentials: credentials.NewStaticCredentials(sc.AccessKeyId, sc.SecretAccessKey, ""),
	}))
	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(sc.BucketName),
		Key:    aws.String("simple/simple.svg"),
		Body:   f,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("file uploaded to, %s\n", aws.String(result.Location))
}
