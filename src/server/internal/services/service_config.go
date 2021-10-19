/*
	This file contains the configuration object used by the service provider.
	This allows configuration of the AWS S3 bucket used for file persistance
	and the AWS region. Other configuration can be added as required for other
	services.
*/
package services

type Config struct {
	S3BucketName	string
	Region			string
	StaticFolder	string
}

func NewConfig() * Config {
	return &Config {
		S3BucketName:	"stem-practice-demo",
		Region:			"us-east-1",
		StaticFolder:	"../ui/out",
	}
}
