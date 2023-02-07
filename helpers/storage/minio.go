package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage struct {
	minioClient *minio.Client
}

var h = Storage{}

func MinioConnection() (*minio.Client, error) {
	ctx := context.Background()
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESSKEY")
	secretAccessKey := os.Getenv("MINIO_SECRETKEY")
	useSSL := true
	minioClient, errInit := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if errInit != nil {
		log.Fatalln(errInit)
	}
	bucketName := os.Getenv("MINIO_BUCKET")

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
	h.minioClient = minioClient
	return minioClient, errInit
}

func GetFile(objectName string) []byte {
	object, err := h.minioClient.GetObject(context.Background(), os.Getenv("MINIO_BUCKET"), objectName, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer object.Close()
	if err != nil {
		log.Println("", err)
		return nil
	}
	defer object.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, object); err != nil {
		log.Println("", err)
		return nil
	}
	return buf.Bytes()
}
