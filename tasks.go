package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var client *s3.Client

// Create a new S3 Client
func createClient() *s3.Client {
	cfg, err := config.LoadDefaultConfig(context.Background())
	inspect(err)

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
}

// List all buckets in the Object Store
func listBuckets(client *s3.Client) {
	result, err := client.ListBuckets(context.Background(), &s3.ListBucketsInput{})
	inspect(err)

	fmt.Println("List of all buckets:")
	for _, bucket := range result.Buckets {
		fmt.Println(*bucket.Name)
	}
}

// Upload an object (file) to the object store
func uploadObject(client *s3.Client, bucket, source, key string) {
	file, err := os.Open(source)
	inspect(err)
	defer file.Close()

	_, err = client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})
	inspect(err)

	fmt.Println("File uploaded successfully")
}

// Upload a large object (file) to the object store
func uploadLargeObject(client *s3.Client, bucket, source, key string) {
	file, err := os.Open(source)
	inspect(err)
	defer file.Close()

	uploader := transfermanager.New(client)

	_, err = uploader.UploadObject(context.Background(), &transfermanager.UploadObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	})
	inspect(err)

	fmt.Println("Large file uploaded successfully")
}

// Download an object (file) from the object store
func downloadObject(client *s3.Client, bucket, key, dest string) {
	result, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	inspect(err)
	defer result.Body.Close()

	file, err := os.Create(dest)
	inspect(err)
	defer file.Close()

	_, err = io.Copy(file, result.Body)
	inspect(err)

	fmt.Println("File downloaded successfully")
}

// Delet an object (file) from the object store
func deleteObject(client *s3.Client, bucket, key string) {
	_, err := client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	inspect(err)

	fmt.Println("File deleted successfully")
}

// Create a new bucket in the object store
func createBucket(client *s3.Client, bucket string) {
	_, err := client.CreateBucket(context.Background(), &s3.CreateBucketInput{
		Bucket: &bucket,
	})
	inspect(err)

	fmt.Println("Bucket created successfully")
}

// Delete an existing bucket from the object store
func deleteBucket(client *s3.Client, bucket string) {
	_, err := client.DeleteBucket(context.Background(), &s3.DeleteBucketInput{
		Bucket: &bucket,
	})
	inspect(err)

	fmt.Println("Bucket deleted successfully")
}
