package usecase

import (
	"backend-golang/repository/database"
	"context"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

func RegisterAsMember(id uint) error {
	user, err := database.ViewMemberInformation(id)
	if err != nil {
		return err
	}

	user.Role = "member"
	err = database.UpdateUserRole(user.ID, user.Role)
	if err != nil {
		return err
	}

	memberCode := database.GenerateMemberCode(user.ID)
	user.MemberCode = memberCode

	err = database.UpdateUser(&user)
	if err != nil {
		return err
	}

	// Generate barcode png
	err = godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	writer := oned.NewCode128Writer()
	img, err := writer.Encode(memberCode, gozxing.BarcodeFormat_CODE_128, 250, 50, nil)
	if err != nil {
		log.Fatalf("impossible to encode barcode: %s", err)
		return err
	}
	file, err := os.Create(memberCode + ".png")
	if err != nil {
		log.Fatalf("impossible to create file: %s", err)
		return err
	}

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("impossible to encode barcode in PNG: %s", err)
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	// Save the barcode to S3
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}
	f, err := os.Open(memberCode + ".png")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("capstone-kelompok17"),
		Key:    aws.String(f.Name()),
		Body:   f,
		ACL:    "public-read",
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to upload photo")
	}
	user.Barcode = result.Location
	if err := database.UpdateUser(&user); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Remove(memberCode + ".png"); err != nil {
		return err
	}

	return nil
}
