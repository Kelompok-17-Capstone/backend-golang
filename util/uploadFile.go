package util

import (
	"context"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func UploadFile(file *multipart.FileHeader) (*manager.UploadOutput, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Error loading .env file")
	}
	f, err := file.Open()
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to upload photo")
	}
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}

	client := s3.NewFromConfig(cfg)

	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("capstone-kelompok17"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Failed to upload photo")
	}
	return result, nil
}
