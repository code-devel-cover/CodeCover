package core

import (
	"context"
	"time"
)

//go:generate mockgen -package mock -destination ../mock/scm_mock.go . SCMService,Client,RepoService,UserService,ContentService,GitService

type SCMService interface {
	Client(scm SCMProvider) (Client, error)
}

type Token struct {
	Token   string
	Refresh string
	Expires time.Time
}

type Client interface {
	Repositories() RepoService
	Users() UserService
	Git() GitService
	Contents() ContentService
	Token(user *User) Token
}

// RepoService provides operations with SCM
type RepoService interface {
	NewReportID(repo *Repo) string
	// List repositories from SCM context
	List(ctx context.Context, user *User) ([]*Repo, error)
	Find(ctx context.Context, user *User, name string) (*Repo, error)
	CloneURL(ctx context.Context, user *User, name string) (string, error)
}

// UserService defines operations with SCM
type UserService interface {
	Find(ctx context.Context, token *Token) (*User, error)
	Create(ctx context.Context, token *Token) (*User, error)
}

type GitService interface {
	GitRepository(ctx context.Context, user *User, repo string) (GitRepository, error)
	FindCommit(ctx context.Context, user *User, repo *Repo) string
}

type ContentService interface {
	ListAllFiles(ctx context.Context, user *User, repo, ref string) ([]string, error)
	Find(ctx context.Context, user *User, repo, path, ref string) ([]byte, error)
}
