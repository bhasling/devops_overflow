/*
	Unit test of the service_config.
	This uses test helper methods found in services_test_helper_test.go
*/
package services

import (
	"testing"
)
func TestShouldReadConfig(t *testing.T) {
	// For this test the code reads the default configuration in the top level application folder
	config := NewConfig()
	err := config.LoadConfig("../../config.yaml")
	ExpectNoError(t, err)
	ExpectEquals(t, config.Region, "us-east-1")
}
