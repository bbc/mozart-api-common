package storage

import (
	"bytes"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Storage interface {
	Get(key string) (*s3.GetObjectOutput, *Error)
	Set(key string, data []byte) (*s3.PutObjectOutput, *Error)
}

type S3Storage struct {
	service *s3.S3
}

func (s *S3Storage) getService() *s3.S3 {
	if s.service == nil {
		config := &aws.Config{
			Region:   aws.String(os.Getenv("AWS_REGION")),
			Endpoint: aws.String(os.Getenv("S3_ENDPOINT")),
		}

		if os.Getenv("APP_ENV") != "production" {
			config.DisableSSL = aws.Bool(true)
			config.S3ForcePathStyle = aws.Bool(true)
		}

		s.service = s3.New(config)
	}

	return s.service
}

func (s *S3Storage) Get(key string) (*s3.GetObjectOutput, *Error) {
	response, err := s.getService().GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(key),
	})
	return response, handleError(err)
}

func (s *S3Storage) Set(key string, data []byte) (*s3.PutObjectOutput, *Error) {
	response, err := s.getService().PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET")),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/json"),
	})
	return response, handleError(err)
}

func handleError(err error) *Error {
	var storageError *Error

	if awsErr, ok := err.(awserr.Error); ok {
		storageError = &Error{}
		storageError.Message = awsErr.Code() + ": " + awsErr.Message()
		if reqErr, ok := err.(awserr.RequestFailure); ok {
			storageError.Status = reqErr.StatusCode()
		}
	}

	return storageError
}

type Error struct {
	Message string
	Status  int
}

func (e *Error) Error() string {
	return e.Message
}
