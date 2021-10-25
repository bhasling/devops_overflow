/*
	Unit test of the service_provider
	This uses test helper methods found in services_test_helper_test.go
*/
package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
func TestServiceProviderHappyPath(t *testing.T) {
	var config = NewConfig()
	serviceProvider := NewServiceProvider(config)
	issueService := serviceProvider.GetIssueService()
	assert.NotEqual(t, nil, issueService)
}
