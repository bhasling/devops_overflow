/*
	The file contains functions to support unit tests in the resources package.
*/
package resources

import (
	"testing"
	"path/filepath"
	"fmt"
	"runtime"
	"reflect"
	"errors"
	"main/internal/services"
)

func ExpectTrue(t *testing.T, condition bool, message string) {
	if ! condition {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("[%s:%d] Expected true. %s\n", filepath.Base(file), line, message)
		t.FailNow()
	}
}

func ExpectNoError(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("[%s:%d] Expected no error. Got '%s'.\n", filepath.Base(file), line, err)
		t.FailNow()
	}
}

func ExpectError(t *testing.T, err error) {
	if err == nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("[%s:%d] Expected an error.\n", filepath.Base(file), line)
		t.FailNow()
	}
}

func ExpectEquals(t *testing.T, act, exp interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("[%s:%d] Expected '%v' Got '%v'.\n", filepath.Base(file), line, exp, act)
		t.FailNow()
	}
}

func ExpectNotEquals(t *testing.T, act, exp interface{}) {
	if reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("[%s:%d] Expected '%v' to be not equal to '%v'.\n", filepath.Base(file), line, exp, act)
		t.FailNow()
	}
}

// Create a Mocked version of the PersistantFileService
type MockedPersistedFileService struct {
	getFolderResultsIndex int
	getFolderResults [][]string
	getFolderErrors []error
	getFileResultsIndex int
	getFileResults []string
	getFileErrors [] error
	writeFileErrorsIndex int
	writeFileErrors [] error
	deleteFileErrorsIndex int
	deleteFileErrors [] error
}

// Methods in the mocked class to add mock return results
func (s *MockedPersistedFileService ) AddGetFolderResult(result []string, err error) {
	s.getFolderResults = append(s.getFolderResults, result)
	s.getFolderErrors = append(s.getFolderErrors, err)
}

func (s *MockedPersistedFileService ) AddGetFileResult(result string, err error) {
	s.getFileResults = append(s.getFileResults, result)
	s.getFileErrors = append(s.getFileErrors, err)
}

func (s *MockedPersistedFileService ) AddWriteFileResult(err error) {
	s.writeFileErrors = append(s.writeFileErrors, err)
}

func (s *MockedPersistedFileService ) AddDeleteFileResult(err error) {
	s.deleteFileErrors = append(s.deleteFileErrors, err)
}

// Method to create a mocked PersistantFileService and inject it into the serviceProvider for testing
func CreateMockfileService(serviceProvider *services.ServiceProvider) *MockedPersistedFileService {
	mockedPersistedFileService := &MockedPersistedFileService{
		getFolderResultsIndex : 0,
		getFolderResults : make([][] string, 0),
	}
	serviceProvider.SetPersistedFileService(mockedPersistedFileService)
	return mockedPersistedFileService
}

// Methods in the mocked class to satisfy the interface and return the mocked return values
func (s *MockedPersistedFileService ) GetFolders(key string) ([] string, error) {
	var result [] string = nil
	var err error = nil
	if s.getFolderResultsIndex < len(s.getFolderResults) {
		result = s.getFolderResults[s.getFolderResultsIndex]
		s.getFolderResultsIndex = s.getFolderResultsIndex + 1
	}
	if result == nil {
		err = errors.New("Error getting folders")
	}
	return result, err
}
func (s *MockedPersistedFileService ) GetFile(key string) (string, error) {
	var result string = ""
	var err error = nil
	if s.getFileResultsIndex < len(s.getFileResults) {
		result = s.getFileResults[s.getFileResultsIndex]
		s.getFileResultsIndex = s.getFileResultsIndex + 1
	}
	return result, err
}
func (s *MockedPersistedFileService ) GetBinaryFile(key string) ([] byte, error) {
	return make([] byte, 0), nil
}
func (s *MockedPersistedFileService ) WriteFile(key string, value string) error {
	var err error = nil
	if s.writeFileErrorsIndex < len(s.writeFileErrors) {
		err = s.writeFileErrors[s.writeFileErrorsIndex]
		s.writeFileErrorsIndex = s.writeFileErrorsIndex + 1
	}
	return err
}
func (s *MockedPersistedFileService ) DeleteFile(key string) error {
	var err error = nil
	if s.deleteFileErrorsIndex < len(s.deleteFileErrors) {
		err = s.deleteFileErrors[s.deleteFileErrorsIndex]
		s.deleteFileErrorsIndex = s.deleteFileErrorsIndex + 1
	}
	return err
}
