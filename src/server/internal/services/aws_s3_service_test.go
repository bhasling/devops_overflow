/*
	Unit test orfthe AWSS3Service.
	This uses test helper methods found in services_test_helper_test.go
*/
package services

import (
	"testing"
)
func TestShouldReadFolders(t *testing.T) {
	// There is no mock for this test. It runs the test on the AWSS3Service with no existing S3 bucket
	// So the test expects an error and zero keys returned
	// If we mocked out the AWSService the test would just test the mock.
	var config = NewConfig()
	var providerService = NewServiceProvider(config)
	var fileService = providerService.GetPersistedFileService()

	var key string = "tasks/"
	keys, err := fileService.GetFolders(key)
	ExpectError(t, err)
	ExpectEquals(t, len(keys), 0)
}
