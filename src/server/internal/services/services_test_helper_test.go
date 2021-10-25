/*
	The file contains functions to support unit tests in the services package.
*/
package services

import (
	"testing"
	"path/filepath"
	"fmt"
	"runtime"
	"reflect"
	"errors"
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

// Class to represent mocked call input/outputs to GetFile,WriteFile,or DeleteFile methods
type MockedPersistedFileResults struct {
	position					int					// Nth call to mocked method
	mockResults					[]string			// simulated content result to return from call
	mockErrors					[]error				// simulated error result to return from call
	callKeys					[]string			// key passed in call
	callContents				[]string			// contents passed in call
}

// Class to represent mocked call input/outputs to GetFolder method
type MockedPersistedFolderResults struct {
	position					int					// Nth call to mocked method
	mockResults					[][]string			// simulated list of keys to return from call
	errorResults				[]error				// simulated error result to return from call
}

// Create a Mocked version of the PersistantFileService
type MockedPersistedFileService struct {
	getFoldersMock			MockedPersistedFolderResults
	getFileMock				MockedPersistedFileResults
	writeFileMock			MockedPersistedFileResults
	deleteFileMock			MockedPersistedFileResults
}

// Methods in the mocked class to add mock return results
func (s *MockedPersistedFileService ) AddGetFolderResult(result []string, err error) {
	s.getFoldersMock.mockResults = append(s.getFoldersMock.mockResults, result)
	s.getFoldersMock.errorResults = append(s.getFoldersMock.errorResults, err)
}

func (s *MockedPersistedFileService ) AddGetFileResult(result string, err error) {
	s.getFileMock.mockResults = append(s.getFileMock.mockResults, result)
	s.getFileMock.mockErrors = append(s.getFileMock.mockErrors, err)
}

func (s *MockedPersistedFileService ) AddWriteFileResult(err error) {
	s.writeFileMock.mockErrors = append(s.writeFileMock.mockErrors, err)
}

func (s *MockedPersistedFileService ) AddDeleteFileResult(err error) {
	s.deleteFileMock.mockErrors = append(s.deleteFileMock.mockErrors, err)
}

// Method to create a mocked PersistantFileService and inject it into the serviceProvider for testing
func CreateMockfileService(serviceProvider *ServiceProvider) *MockedPersistedFileService {
	mockedPersistedFileService := &MockedPersistedFileService{
	}	
	serviceProvider.SetPersistedFileService(mockedPersistedFileService)
	return mockedPersistedFileService
}

// Methods in the mocked class to satisfy the interface and return the mocked return values
func (s *MockedPersistedFileService ) GetFolders(key string) ([] string, error) {
	var result [] string = nil
	var err error = nil
	if s.getFoldersMock.position < len(s.getFoldersMock.mockResults) {
		result = s.getFoldersMock.mockResults[s.getFoldersMock.position]
		s.getFoldersMock.position = s.getFoldersMock.position + 1
	}
	if result == nil {
		err = errors.New("Error getting folders")
	}
	return result, err
}
func (s *MockedPersistedFileService ) GetFile(key string) (string, error) {
	var result string = ""
	var err error = nil

	if s.getFileMock.position < len(s.getFileMock.mockResults) {
		result = s.getFileMock.mockResults[s.getFileMock.position]
		err = s.getFileMock.mockErrors[s.getFileMock.position]
		s.getFileMock.position = s.getFileMock.position + 1
	}

	return result, err
}
func (s *MockedPersistedFileService ) GetBinaryFile(key string) ([] byte, error) {
	return make([] byte, 0), nil
}
func (s *MockedPersistedFileService ) WriteFile(key string, value string) error {
	var err error = nil

	if s.writeFileMock.position < len(s.writeFileMock.mockResults) {
		err = s.writeFileMock.mockErrors[s.writeFileMock.position]
		s.writeFileMock.position = s.writeFileMock.position + 1
	}
	return err
}
func (s *MockedPersistedFileService ) DeleteFile(key string) error {
	var err error = nil
	if s.deleteFileMock.position < len(s.deleteFileMock.mockResults) {
		err = s.deleteFileMock.mockErrors[s.deleteFileMock.position]
		s.deleteFileMock.position = s.deleteFileMock.position + 1
	}

	return err
}
