package storage

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Storage interface {
	Get(key string) (string, *Error)
	Set(key string, data string) *Error
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

func (s *S3Storage) Get(key string) (string, *Error) {
	response, err := s.getService().GetObject(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    aws.String(key),
	})

	var object []byte

	if err != nil {
		return string(object), handleAWSError(err)
	}

	object, readErr := ioutil.ReadAll(response.Body)

	if readErr != nil {
		return string(object), &Error{Message: readErr.Error()}
	}

	return string(object), nil
}

func (s *S3Storage) Set(key string, data string) *Error {
	_, err := s.getService().PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET")),
		Key:         aws.String(key),
		Body:        bytes.NewReader([]byte(data)),
		ContentType: aws.String("application/json"),
	})

	return handleAWSError(err)
}

func handleAWSError(err error) *Error {
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
