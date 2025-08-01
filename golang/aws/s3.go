package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const MSG_ERROR_CREATE_FILE = "Unable to open file\n %q, %v"
const MSG_ERROR_UPLOAD_FILE = "Unable to upload %q to %q, %v\n"

func Upload(filename string) error {
	cfg, err := LoadCustomAWSConfig()

	if err == nil {
		file, err := os.Open("tmp/" + filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, MSG_ERROR_CREATE_FILE, filename, err)
			return err
		}

		defer file.Close()

		bucket := os.Getenv("BUCKET_S3")
		client := s3.NewFromConfig(cfg)
		uploader := manager.NewUploader(client)

		_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
			Body:   file,
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, MSG_ERROR_UPLOAD_FILE, filename, bucket, err)
			return err
		}

		return nil
	}

	return err
}

func LoadCustomAWSConfig() (aws.Config, error) {
	profile := os.Getenv("KEY_PROFILE")
	region := os.Getenv("S3_REGION")

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region),
	)

	return cfg, err
}
