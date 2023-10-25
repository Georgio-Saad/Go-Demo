package helpers

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func ConnectToBucket() *session.Session {
	bucketKeyId := os.Getenv("AWS_S3_ID")
	bucketSecret := os.Getenv("AWS_S3_SECRET")
	bucketRegion := os.Getenv("AWS_S3_REGION")

	sess, sessErr := session.NewSession(&aws.Config{
		Region:      aws.String(bucketRegion),
		Credentials: credentials.NewStaticCredentials(bucketKeyId, bucketSecret, ""),
	})

	if sessErr != nil {
		panic(sessErr)
	}

	return sess
}
