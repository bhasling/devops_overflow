/*
	Unit test or the AWSS3Service.
*/
package services

import (
	"testing"
)
func TestS3ReadFolder(t *testing.T) {
	var config = NewConfig()
	var providerService = NewServiceProvider(config)
	var fileService = providerService.getPersistedFileService()

	var key string
	key = "tasks/"
	keys, _ := fileService.GetFolders(key)
	if (len(keys) == 0) {
		t.Errorf("Expected number of keys to > zero.")
	}
}
