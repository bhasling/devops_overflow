/*
	This file contains the configuration object used by the service provider.
	This allows configuration of the AWS S3 bucket used for file persistance
	and the AWS region. Other configuration can be added as required for other
	services.
*/
package services

import (
	"log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	S3BucketName	string 	`yaml:"S3BucketName"`
	Region			string 	`yaml:"region"`
	StaticFolder	string 	`yaml:"StaticFolder"`
}

func NewConfig() * Config {
	return &Config {
		S3BucketName:	"configure_this_bucstemket_name",
		Region:			"us-east-1",
		StaticFolder:	"../ui/out",
	}
}

func (config *Config) LoadConfig(fileName string) error {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("Error loading config file: %v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Printf("Error parsing config file: %v", err)
		return err
	}
	return nil
}
