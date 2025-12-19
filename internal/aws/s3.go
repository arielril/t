package aws

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadFiles(cfg aws.Config, bucketName string, fileNames []string) {
	for _, file := range fileNames {

		fileContent, err := os.Open(file)
		if err != nil {
			fmt.Printf("failed to open file %s: %s\n", file, err)
			continue
		}

		client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.EndpointOptions.DisableHTTPS = true
		})
		putOutput, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
			Body:   fileContent,
			Bucket: aws.String(bucketName),
			Key:    aws.String(file),
		})
		if err != nil {
			fmt.Printf("failed to upload file %s: %s\n", file, err)
			continue
		}

		fmt.Printf("successfully uploaded file %s (upload sha256 %v)\n", file,
			putOutput.ChecksumSHA256)

		_ = fileContent.Close()
	}

	fmt.Printf("finished uploading %d files: %v\n", len(fileNames), fileNames)
}
