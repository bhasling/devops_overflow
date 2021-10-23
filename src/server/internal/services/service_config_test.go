/*
	Unit test of the service_config.
*/
package services

import (
	"testing"
)
func TestConfig(t *testing.T) {
	config := NewConfig()
	err := config.LoadConfig("../../config.yaml")
	if (err != nil) {
		t.Errorf("Unable to read configuration file.")
		return
	}
	if (config.Region != "us-east-1") {
		t.Errorf("Expected region read from config file")
		return
	}
}
