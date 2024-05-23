package service

import (
	"airbnb/features/review"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type reviewService struct {
	reviewData   review.DataInterface
	s3           *s3.S3
	s3BucketName string
	env          string
}

// UpdateById implements review.ServiceInterface.
func (r *reviewService) UpdateById(id uint, input review.Core, file io.Reader, handlerFilename string) (string, error) {
	// Memeriksa apakah ID ulasan valid
	if id <= 0 {
		return "", errors.New("id not valid")
	}

	// Mengupdate data ulasan berdasarkan ID
	err := r.reviewData.EditById(id, input)
	if err != nil {
		return "", err
	}

	// Jika ada file yang diunggah, lakukan pembaruan pada foto ulasan
	if file != nil {
		// Upload file baru ke penyimpanan (S3 atau lokal)
		var photoURL string
		var errUpload error
		if r.env == "production" {
			photoURL, errUpload = r.UploadFileToS3(file, handlerFilename)
		} else {
			photoURL, errUpload = r.SaveFileLocally(file, handlerFilename)
		}
		if errUpload != nil {
			return "", errUpload
		}

		// Update URL foto pada data ulasan
		input.Foto = photoURL

		// Simpan perubahan data ulasan dengan URL foto yang baru
		err := r.reviewData.EditById(id, input)
		if err != nil {
			return "", err
		}
	}

	// Jika tidak ada file yang diunggah, kembalikan URL foto yang sudah ada
	return input.Foto, nil
}

func New(rd review.DataInterface, s3 *s3.S3, bucketName, env string) review.ServiceInterface {
	return &reviewService{
		reviewData:   rd,
		s3:           s3,
		s3BucketName: bucketName,
		env:          env,
	}
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

	var photoURL string
	var err error

	if r.env == "production" {
		photoURL, err = r.UploadFileToS3(file, fileName)
		if err != nil {
			return "", err
		}
		input.Foto = photoURL
	} else {
		photoURL, err = r.SaveFileLocally(file, fileName)
		if err != nil {
			return "", err
		}
		input.Foto = photoURL
	}

	err = r.reviewData.Insert(input)
	if err != nil {
		return "", err
	}
	return input.Foto, nil
}

// func (r *reviewService) Create(input review.Core, file io.Reader, handlerFilename string) (string, error) {

// 	timestamp := time.Now().Unix()
// 	fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
// 	// photoFileName, errPhoto := r.UploadFileToS3(file, fileName)
// 	localFilePath, errPhoto := r.SaveFileLocally(file, fileName)
// 	if errPhoto != nil {
// 		return "", errPhoto
// 	}
// 	input.Foto = localFilePath
// 	// input.Foto = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", r.s3BucketName, photoFileName)
// 	err := r.reviewData.Insert(input)
// 	if err != nil {
// 		return "", err
// 	}
// 	return input.Foto, nil
// }

// GetAll implements review.ServiceInterface.
func (r *reviewService) GetAll() ([]review.Core, error) {
	return r.reviewData.SelectAll()
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
func (r *reviewService) SaveFileLocally(file io.Reader, fileName string) (string, error) {
	localDir := "./utils/uploads"

	// Pastikan direktori ada
	if _, err := os.Stat(localDir); os.IsNotExist(err) {
		errDir := os.MkdirAll(localDir, 0755)
		if errDir != nil {
			return "", errDir
		}
	}

	localFilePath := fmt.Sprintf("%s/%s", localDir, fileName)

	localFile, err := os.Create(localFilePath)
	if err != nil {
		return "", err
	}
	defer localFile.Close()

	if _, err := io.Copy(localFile, file); err != nil {
		return "", err
	}

	return localFilePath, nil
}
