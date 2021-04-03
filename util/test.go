package util

import (
	"testing"
)

type BaseRepositoryTest interface {
	GetByIdTest(t *testing.T)
	CreateTest(t *testing.T)
	UpdateTest(t *testing.T)
	DeleteTest(t *testing.T)
}

//BaseRepoTester provides docker test suites - needs db migration!
type BaseRepoTester struct {
}

type BaseHandlerTest interface {
	GETTest(t *testing.T)
	POSTTest(t *testing.T)
	PUTTest(t *testing.T)
	DELETETest(t *testing.T)
}
