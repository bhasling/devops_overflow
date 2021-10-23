/*
	This file implements Issue service. This provides business logic operations to support
	the Issues that are submitted by a user.
*/
package services

import (
	"fmt"
	"errors"
	"encoding/json"
	"time"
	"github.com/google/uuid"
	"log"
)

type IssueServiceImpl struct {
	persistedFileService		PersistedFileService
}

type Answer struct {
    AnswerId        string          `json:answer_id"`
    User            string          `json:"user_id"`
    AnswerTime	    time.Time       `json:answer_time"`
    Description     string          `json:"desciption"`
}
type Issue struct {
    IssueId         string          `json:"issue_id"`
    User            string          `json:"user_id"`
    Title           string          `json:"title"`
    Product         string          `json:"product"`
    SubmitTime      time.Time       `json:response_time"`
    Description     string          `json:"description"`
    Answers	       []Answer     	`json:"answers"`
}

func ParseDate(value string) time.Time {
    var result time.Time
    result, _ = time.Parse("2006-01-02", value)
    return result
}

func NewIssueService(persistedFileService PersistedFileService) * IssueServiceImpl {
	return &IssueServiceImpl{
		persistedFileService: persistedFileService,
	}
}

func (s *IssueServiceImpl ) GetAllIssues() ([]Issue, error) {
	result := make([]Issue, 0)
	keys, err := s.persistedFileService.GetFolders("issues/")
	if err != nil {
		log.Println(err)
		return result, err
	}
	for _, key := range keys {
		content, _ := s.persistedFileService.GetFile(key)
		issue := &Issue{}
		json.Unmarshal([]byte(content), &issue)
		result = append(result, *issue)
	}
	return result, nil
}

func (s *IssueServiceImpl ) CreateIssue() (*Issue, error) {
	id := uuid.New().String()
	result := Issue {IssueId: id}
	return &result, nil
}

func (s *IssueServiceImpl ) CreateAnswer(issue *Issue) (*Answer, error) {
	id := uuid.New().String()
	answer := Answer {AnswerId: id}
	issue.Answers = append(issue.Answers, answer)
	return &answer, nil
}

func (s *IssueServiceImpl ) UpdateAnswer(issue *Issue, answer *Answer) {
	for i, a := range issue.Answers {
		if (a.AnswerId == answer.AnswerId) {
			issue.Answers[i] = *answer
			break
		}
	}
}

func (s *IssueServiceImpl ) FindAnswerById(issue *Issue, answerId string) *Answer {
	for _, answer := range issue.Answers {
		if (answer.AnswerId == answerId) {
			return &answer
		}
	}
	return nil
}

func (s *IssueServiceImpl ) DeleteAnswerById(issue *Issue, answerId string) error {
	for i, answer := range issue.Answers {
		if (answer.AnswerId == answerId) {
			issue.Answers = append(issue.Answers[:i], issue.Answers[i+1:]...)
			return nil
		}
	}
	return errors.New(fmt.Sprintf("No such answer id %s", answerId))
}

func (s *IssueServiceImpl ) SaveIssue(issue *Issue) error {
	content, _ := json.Marshal(issue)
	var key = fmt.Sprintf("issues/%s", issue.IssueId)
	err := s.persistedFileService.WriteFile(key, string(content))
	if err != nil {
		log.Printf("Unable to save issue '%s': %v", key, err)
		return err
	}
	return nil
}

func (s *IssueServiceImpl ) GetIssueById(issueId string) (*Issue, error) {
	var key = fmt.Sprintf("issues/%s", issueId)
	content, err := s.persistedFileService.GetFile(key)
	if (err != nil) {
		log.Println(err)
		return nil, err
	}
	issue := &Issue{}
	json.Unmarshal([]byte(content), &issue)
	return issue, nil
}

func (s *IssueServiceImpl ) DeleteIssueById(issueId string) error {
	var key = fmt.Sprintf("issues/%s", issueId)
	err := s.persistedFileService.DeleteFile(key)
	if (err != nil) {
		log.Println(err)
		return err
	}
	return nil
}

