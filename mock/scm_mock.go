// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/code-devel-cover/CodeCover/core (interfaces: SCMService,Client,RepoService,UserService,ContentService,GitService)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	core "github.com/code-devel-cover/CodeCover/core"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSCMService is a mock of SCMService interface
type MockSCMService struct {
	ctrl     *gomock.Controller
	recorder *MockSCMServiceMockRecorder
}

// MockSCMServiceMockRecorder is the mock recorder for MockSCMService
type MockSCMServiceMockRecorder struct {
	mock *MockSCMService
}

// NewMockSCMService creates a new mock instance
func NewMockSCMService(ctrl *gomock.Controller) *MockSCMService {
	mock := &MockSCMService{ctrl: ctrl}
	mock.recorder = &MockSCMServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSCMService) EXPECT() *MockSCMServiceMockRecorder {
	return m.recorder
}

// Client mocks base method
func (m *MockSCMService) Client(arg0 core.SCMProvider) (core.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client", arg0)
	ret0, _ := ret[0].(core.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Client indicates an expected call of Client
func (mr *MockSCMServiceMockRecorder) Client(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockSCMService)(nil).Client), arg0)
}

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Contents mocks base method
func (m *MockClient) Contents() core.ContentService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contents")
	ret0, _ := ret[0].(core.ContentService)
	return ret0
}

// Contents indicates an expected call of Contents
func (mr *MockClientMockRecorder) Contents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contents", reflect.TypeOf((*MockClient)(nil).Contents))
}

// Git mocks base method
func (m *MockClient) Git() core.GitService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Git")
	ret0, _ := ret[0].(core.GitService)
	return ret0
}

// Git indicates an expected call of Git
func (mr *MockClientMockRecorder) Git() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Git", reflect.TypeOf((*MockClient)(nil).Git))
}

// Repositories mocks base method
func (m *MockClient) Repositories() core.RepoService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Repositories")
	ret0, _ := ret[0].(core.RepoService)
	return ret0
}

// Repositories indicates an expected call of Repositories
func (mr *MockClientMockRecorder) Repositories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Repositories", reflect.TypeOf((*MockClient)(nil).Repositories))
}

// Token mocks base method
func (m *MockClient) Token(arg0 *core.User) core.Token {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token", arg0)
	ret0, _ := ret[0].(core.Token)
	return ret0
}

// Token indicates an expected call of Token
func (mr *MockClientMockRecorder) Token(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockClient)(nil).Token), arg0)
}

// Users mocks base method
func (m *MockClient) Users() core.UserService {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Users")
	ret0, _ := ret[0].(core.UserService)
	return ret0
}

// Users indicates an expected call of Users
func (mr *MockClientMockRecorder) Users() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Users", reflect.TypeOf((*MockClient)(nil).Users))
}

// MockRepoService is a mock of RepoService interface
type MockRepoService struct {
	ctrl     *gomock.Controller
	recorder *MockRepoServiceMockRecorder
}

// MockRepoServiceMockRecorder is the mock recorder for MockRepoService
type MockRepoServiceMockRecorder struct {
	mock *MockRepoService
}

// NewMockRepoService creates a new mock instance
func NewMockRepoService(ctrl *gomock.Controller) *MockRepoService {
	mock := &MockRepoService{ctrl: ctrl}
	mock.recorder = &MockRepoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepoService) EXPECT() *MockRepoServiceMockRecorder {
	return m.recorder
}

// CloneURL mocks base method
func (m *MockRepoService) CloneURL(arg0 context.Context, arg1 *core.User, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloneURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloneURL indicates an expected call of CloneURL
func (mr *MockRepoServiceMockRecorder) CloneURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloneURL", reflect.TypeOf((*MockRepoService)(nil).CloneURL), arg0, arg1, arg2)
}

// Find mocks base method
func (m *MockRepoService) Find(arg0 context.Context, arg1 *core.User, arg2 string) (*core.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2)
	ret0, _ := ret[0].(*core.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockRepoServiceMockRecorder) Find(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockRepoService)(nil).Find), arg0, arg1, arg2)
}

// List mocks base method
func (m *MockRepoService) List(arg0 context.Context, arg1 *core.User) ([]*core.Repo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*core.Repo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockRepoServiceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockRepoService)(nil).List), arg0, arg1)
}

// NewReportID mocks base method
func (m *MockRepoService) NewReportID(arg0 *core.Repo) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewReportID", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// NewReportID indicates an expected call of NewReportID
func (mr *MockRepoServiceMockRecorder) NewReportID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewReportID", reflect.TypeOf((*MockRepoService)(nil).NewReportID), arg0)
}

// MockUserService is a mock of UserService interface
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUserService) Create(arg0 context.Context, arg1 *core.Token) (*core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(*core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserService)(nil).Create), arg0, arg1)
}

// Find mocks base method
func (m *MockUserService) Find(arg0 context.Context, arg1 *core.Token) (*core.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(*core.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockUserServiceMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockUserService)(nil).Find), arg0, arg1)
}

// MockContentService is a mock of ContentService interface
type MockContentService struct {
	ctrl     *gomock.Controller
	recorder *MockContentServiceMockRecorder
}

// MockContentServiceMockRecorder is the mock recorder for MockContentService
type MockContentServiceMockRecorder struct {
	mock *MockContentService
}

// NewMockContentService creates a new mock instance
func NewMockContentService(ctrl *gomock.Controller) *MockContentService {
	mock := &MockContentService{ctrl: ctrl}
	mock.recorder = &MockContentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContentService) EXPECT() *MockContentServiceMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockContentService) Find(arg0 context.Context, arg1 *core.User, arg2, arg3, arg4 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockContentServiceMockRecorder) Find(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockContentService)(nil).Find), arg0, arg1, arg2, arg3, arg4)
}

// ListAllFiles mocks base method
func (m *MockContentService) ListAllFiles(arg0 context.Context, arg1 *core.User, arg2, arg3 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllFiles", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllFiles indicates an expected call of ListAllFiles
func (mr *MockContentServiceMockRecorder) ListAllFiles(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllFiles", reflect.TypeOf((*MockContentService)(nil).ListAllFiles), arg0, arg1, arg2, arg3)
}

// MockGitService is a mock of GitService interface
type MockGitService struct {
	ctrl     *gomock.Controller
	recorder *MockGitServiceMockRecorder
}

// MockGitServiceMockRecorder is the mock recorder for MockGitService
type MockGitServiceMockRecorder struct {
	mock *MockGitService
}

// NewMockGitService creates a new mock instance
func NewMockGitService(ctrl *gomock.Controller) *MockGitService {
	mock := &MockGitService{ctrl: ctrl}
	mock.recorder = &MockGitServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGitService) EXPECT() *MockGitServiceMockRecorder {
	return m.recorder
}

// FindCommit mocks base method
func (m *MockGitService) FindCommit(arg0 context.Context, arg1 *core.User, arg2 *core.Repo) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindCommit", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	return ret0
}

// FindCommit indicates an expected call of FindCommit
func (mr *MockGitServiceMockRecorder) FindCommit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindCommit", reflect.TypeOf((*MockGitService)(nil).FindCommit), arg0, arg1, arg2)
}

// GitRepository mocks base method
func (m *MockGitService) GitRepository(arg0 context.Context, arg1 *core.User, arg2 string) (core.GitRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GitRepository", arg0, arg1, arg2)
	ret0, _ := ret[0].(core.GitRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GitRepository indicates an expected call of GitRepository
func (mr *MockGitServiceMockRecorder) GitRepository(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GitRepository", reflect.TypeOf((*MockGitService)(nil).GitRepository), arg0, arg1, arg2)
}
