/*
	Unit test or the AWSS3Service.
*/
package services

import (
	"testing"
	"encoding/json"
	"strings"
)
func TestServiceHappyPath(t *testing.T) {
	var config = NewConfig()
	serviceProvider := NewServiceProvider(config)
	issueService := serviceProvider.GetIssueService()
	issues, _ := issueService.GetAllIssues()
	issueCount := len(issues)
	issue,_ := issueService.CreateIssue()
	issue.Title = "a title"
	err := issueService.SaveIssue(issue)
	if (err != nil) {
		t.Errorf("Save issue failed %s", err.Error())
	}
	issueId := issue.IssueId
	issue, _ = issueService.GetIssueById(issueId)
	if (issue == nil) {
		t.Errorf("Expected to get issue by Id")
	}
	issues, _ = issueService.GetAllIssues()
	if (len(issues) != issueCount + 1) {
		t.Errorf("Expected %d users got %d.", issueCount + 1, len(issues))
	}
	answer, _ := issueService.CreateAnswer(issue)
	answer.Description = "my answer"
	issueService.UpdateAnswer(issue, answer)
	content, _ := json.Marshal(issue)
	if (! strings.Contains(string(content), "my answer")) {
		t.Errorf("Expected my answer in issue")
	}
	issueService.DeleteAnswerById(issue, answer.AnswerId)
	content, _ = json.Marshal(issue)
	if (len(issue.Answers) != 0) {
		t.Errorf("Did not delete answer")
	}
	err = issueService.DeleteIssueById(issueId)
	if (err != nil) {
		t.Errorf("Unable to delete issue")
	}
	issue, _ = issueService.CreateIssue()
	issue.Title = "My first issue"
	answer, _ = issueService.CreateAnswer(issue)
	answer.Description = "That is just the way it is."
	issueService.SaveIssue(issue)

}
