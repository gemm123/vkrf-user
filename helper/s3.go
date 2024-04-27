package helper

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	configaws "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gemm123/vkrf-user/config"
	"hash/fnv"
	"log"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"
)

func UploadToS3(image *multipart.FileHeader) (string, error) {
	file, err := image.Open()
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
		return "", err
	}
	defer file.Close()

	cfg, err := configaws.LoadDefaultConfig(context.Background(),
		configaws.WithRegion(config.AWSRegion()),
		configaws.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     config.AWSAccessKey(),
				SecretAccessKey: config.AWSSecretKey(),
			}, nil
		})),
	)
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
		return "", err
	}

	hashNameFile := hashFileName(image.Filename)
	s3Client := s3.NewFromConfig(cfg)
	_, err = s3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(config.AWSBucket()),
		Key:    aws.String(hashNameFile),
		Body:   file,
		ACL:    "public-read",
	})
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
		return "", err
	}

	imageUrl := GetS3Link(config.AWSBucket(), hashNameFile)

	return imageUrl, nil
}

func GetS3Link(bucket, key string) string {
	region := config.AWSRegion()
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, key)
}

func hashFileName(fileName string) string {
	currentTime := time.Now().UnixNano()
	timestampString := strconv.FormatInt(currentTime, 10)

	baseName := filepath.Base(fileName)
	extension := filepath.Ext(baseName)
	baseName = baseName[:len(baseName)-len(extension)]

	combinedString := baseName + timestampString

	hasher := fnv.New32a()
	hasher.Write([]byte(combinedString))

	hashValue := fmt.Sprintf("%x", hasher.Sum32())
	uniqueFileName := hashValue + extension

	return uniqueFileName
}
