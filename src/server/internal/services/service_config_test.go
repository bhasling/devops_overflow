/*
	Unit test of the service_config.
	This uses test helper methods found in services_test_helper_test.go
*/
package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestShouldReadConfig(t *testing.T) {
	// For this test the code reads the default configuration in the top level application folder
	config := NewConfig()
	err := config.LoadConfig("../../config.yaml")
	assert.Equal(t, nil, err)
	assert.Equal(t, "us-east-1", config.Region)
}
