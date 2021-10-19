/*
	The services package contains services to support the resources.
	Some service are DAO service that read and read to persistent storage.
	Some services are business object services that may use DAO services or other services.
	Services are stateless except for caches.

	Each service has a service interface to allow dependency injection for unit testing.
	This file contains a ServiceProvider that provides access to the implementation
	of every service. A configuration is provided to the NewServiceProvider that creates
	the service provider instance. This configuration is provided to every service
	created by the service provider. This means services can be aquired by resources
	using GetXXX() methods with no arguments, but the services are configured for the
	environment as required.

	The main function of the application is responsible to create the configuration
	and create the ServiceProvider and put the service provider in the the middleware
	of the HTTP server to be accessible by resources that handle HTTP events.

	The following services are avaiable:
		persistedFileService		- Provide operations on persisted files.
		userService					- Provide operations in user objects.
		issueService				- Provide operations on DEVOP issue objects.
*/

package services

type ServiceProvider struct {
	config	 					*Config
	persistedFileService		PersistedFileService
	userService					UserService
	issueService				IssueService
}

type PersistedFileService interface {
	GetFolders(key string) ([]string, error)
	GetFile(key string) (string, error)
	WriteFile(key string, value string) error
	DeleteFile(key string) error
}

type UserService interface {
	GetAllUsers() ([]User, error)
	CreateUser(userId string) (*User, error)
	SaveUser(user *User) error
	GetUserById(userId string) (*User, error)
	DeleteUserById(userId string) error
}

type IssueService interface {
	GetAllIssues() ([]Issue, error)
	CreateIssue() (*Issue, error)
	CreateAnswer(issue *Issue) (*Answer, error)
	UpdateAnswer(issue *Issue, answer *Answer)
	FindAnswerById(issue *Issue, answerId string) *Answer
	DeleteAnswerById(issue *Issue, answerId string) error
	SaveIssue(issue *Issue) error
	GetIssueById(issueId string) (*Issue, error)
	DeleteIssueById(issueId string) error
}

func (serviceProvider *ServiceProvider) GetConfig() *Config {
	return serviceProvider.config
}

func (serviceProvider *ServiceProvider) getPersistedFileService() PersistedFileService {
	if (serviceProvider.persistedFileService == nil) {
		serviceProvider.persistedFileService = NewAwsS3Service(serviceProvider.config)
	}
	return serviceProvider.persistedFileService
}

func (serviceProvider *ServiceProvider) GetUserService() UserService {
	if (serviceProvider.userService == nil) {
		serviceProvider.userService = NewUserService(serviceProvider.getPersistedFileService())
	}
	return serviceProvider.userService
}

func (serviceProvider *ServiceProvider) GetIssueService() IssueService {
	if (serviceProvider.issueService == nil) {
		serviceProvider.issueService = NewIssueService(serviceProvider.getPersistedFileService())
	}
	return serviceProvider.issueService
}

func NewServiceProvider(config *Config) *ServiceProvider {
	return &ServiceProvider {
		config:config,
	}
}