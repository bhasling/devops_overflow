/*
	Unit test of the AWSS3Service.
	This uses test helper methods found in services_test_helper_test.go
*/
package services

import (
	"testing"
	"encoding/json"
	"strings"
)
func TestIssueHappyPath(t *testing.T) {
	var config = NewConfig()
	serviceProvider := NewServiceProvider(config)
	mockFileService := CreateMockfileService(serviceProvider)
	mockFileService.AddGetFolderResult([]string {}, nil)
	mockFileService.AddGetFolderResult([]string {"1"}, nil)
	mockFileService.AddGetFileResult(`{"issue_id":"1", "title":"my title"}`, nil)
	mockFileService.AddGetFileResult(`{"issue_id":"1", "title":"my title"}`, nil)

	// Get initial issues
	issueService := serviceProvider.GetIssueService()
	initialIssues, _ := issueService.GetAllIssues()

	// Create a new issue
	issue,_ := issueService.CreateIssue()
	issue.Title = "a title"
	err := issueService.SaveIssue(issue)
	ExpectNoError(t, err)

	// Read back the issue by the id
	readBackIssue, _ := issueService.GetIssueById(issue.IssueId)
	ExpectNotEquals(t, readBackIssue, nil)
	ExpectEquals(t, readBackIssue.Title, "my title")

	// Read back new list of issues
	newListOfIssues, _ := issueService.GetAllIssues()
	ExpectEquals(t, len(newListOfIssues), len(initialIssues) + 1)

	// Create a new answer in the issue
	answer, _ := issueService.CreateAnswer(readBackIssue)
	answer.Description = "my answer"
	issueService.UpdateAnswer(readBackIssue, answer)
	content, _ := json.Marshal(readBackIssue)
	ExpectTrue(t, strings.Contains(string(content), "my answer"), "Expected my answer in issue")

	// Delete the new answer
	issueService.DeleteAnswerById(issue, answer.AnswerId)
	content, _ = json.Marshal(issue)
	ExpectEquals(t, len(issue.Answers), 0)

	// Delete the new issue
	err = issueService.DeleteIssueById(issue.IssueId)
	ExpectNoError(t, err)
}
