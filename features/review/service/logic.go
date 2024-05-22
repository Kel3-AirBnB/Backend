package service

import (
	"airbnb/features/review"
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type reviewService struct {
	reviewData   review.DataInterface
	s3           *s3.S3
	s3BucketName string
}

// Delete implements review.ServiceInterface.
func (r *reviewService) Delete(id uint) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	return r.reviewData.Delete(id)
}

// GetReviews implements review.ServiceInterface.
func (r *reviewService) GetReviews(id uint) (data *review.Core, err error) {
	if id <= 0 {
		return nil, errors.New("[validation] id not valid")
	}
	return r.reviewData.SelectById(id)
}

// Create implements review.ServiceInterface.
func (r *reviewService) Create(input review.Core, file io.Reader, handlerFilename string) (string, error) {

	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
	photoFileName, errPhoto := r.UploadFileToS3(file, fileName)
	if errPhoto != nil {
		return "", errPhoto
	}
	input.Foto = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", r.s3BucketName, photoFileName)
	err := r.reviewData.Insert(input)
	if err != nil {
		return "", err
	}
	return input.Foto, nil
}

// GetAll implements review.ServiceInterface.
func (r *reviewService) GetAll() ([]review.Core, error) {
	return r.reviewData.SelectAll()
}

func New(rd review.DataInterface, s3 *s3.S3, bucketName string) review.ServiceInterface {
	return &reviewService{
		reviewData:   rd,
		s3:           s3,
		s3BucketName: bucketName,
	}
}

func (r *reviewService) UploadFileToS3(file io.Reader, fileName string) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	_, err := r.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(r.s3BucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", r.s3BucketName, aws.StringValue(r.s3.Config.Region), fileName), err
}
