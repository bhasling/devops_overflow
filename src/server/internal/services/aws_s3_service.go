/*
	This file implements AwsS3Service that is an implementation of the PersistedFileInterface
	supported by the service provider.
*/
package services

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"bytes"
	"strings"
)

type AwsS3Service struct {
	config		*Config
	awsS3		*s3.S3
}

func NewAwsS3Service(config *Config) * AwsS3Service {
	return &AwsS3Service{config: config}
}

func (s *AwsS3Service ) GetFolders(key string) ([] string, error) {
	log.Printf("AwsS3Service.GetFolders '%s'", key)
	awsS3, err := s.getAwsS3()
	if (err != nil) {
		return nil, err
	}
	input := s3.ListObjectsV2Input{Bucket : &s.config.S3BucketName, Prefix : &key}
	resp, err := awsS3.ListObjectsV2(&input)
	if err != nil {
		log.Println(err)
		return make([]string, 0), err
	}
	result := make([]string, len(resp.Contents))
	for i, item := range resp.Contents {
		result[i] = *item.Key
	}
	log.Printf("AwsS3Service.GetFolders found %d rows.", len(result))
	return result, nil
}

func (s *AwsS3Service ) GetFile(key string) (string, error) {
	log.Printf("AwsS3Service.GetFile '%s'", key)
	awsS3, err := s.getAwsS3()
	if (err != nil) {
		return "", err
	}
	input := s3.GetObjectInput{Bucket : &s.config.S3BucketName, Key : &key}
	rawObject, err := awsS3.GetObject(&input)
	if err != nil {
		log.Println(err)
		return "", err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(rawObject.Body)
	result := buf.String()
	return result, nil
}

func (s *AwsS3Service ) GetBinaryFile(key string) ([]byte, error) {
	log.Printf("AwsS3Service.GetBinaryFile '%s'", key)
	awsS3, err := s.getAwsS3()
	if (err != nil) {
		return make([]byte, 0), err
	}
	input := s3.GetObjectInput{Bucket : &s.config.S3BucketName, Key : &key}
	rawObject, err := awsS3.GetObject(&input)
	if err != nil {
		log.Println(err)
		return make([]byte, 0), err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(rawObject.Body)
	result := buf.Bytes()
	return result, nil
}

func (s *AwsS3Service ) WriteFile(key string, value string) error {
	log.Printf("AwsS3Service.WriteFile '%s'", key)
	awsS3, err := s.getAwsS3()
	if (err != nil) {
		log.Println(err)
		return err
	}
	reader := strings.NewReader(value)
	input := s3.PutObjectInput{Bucket : &s.config.S3BucketName, Key : &key, Body:reader}
	_, err = awsS3.PutObject(&input)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
func (s *AwsS3Service ) DeleteFile(key string) error {
	log.Printf("AwsS3Service.DeleteFile '%s'", key)
	awsS3, err := s.getAwsS3()
	if (err != nil) {
		log.Println(err)
		return err
	}
	input := s3.DeleteObjectInput{Bucket : &s.config.S3BucketName, Key : &key}
	_, err = awsS3.DeleteObject(&input)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *AwsS3Service )  getAwsS3() (*s3.S3, error) {
	if s.awsS3 == nil {
		sess, err := session.NewSession(&aws.Config{Region: aws.String(s.config.Region)})
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		s.awsS3 = s3.New(sess)
	}
	return s.awsS3, nil

}
