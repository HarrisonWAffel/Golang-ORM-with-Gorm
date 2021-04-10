package user

import (
	"github.com/HarrisonWAffel/dbTrain/config"
	"github.com/HarrisonWAffel/dbTrain/domain"
	"github.com/HarrisonWAffel/dbTrain/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestUserRepository(t *testing.T) {
	tester = RepoTester{}
	tester.StartMockDatabase()
	err := config.Read()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(tester.DSN), nil)
	require.NoError(t, err)
	repo, err := NewUserRepository(db)
	require.NoError(t, err)
	tester.repo = repo

	t.Run("Create Test", tester.CreateTest)
	t.Run("GetById Test", tester.GetByIdTest)
	t.Run("Update Test", tester.UpdateTest)
	t.Run("Delete Test", tester.DeleteTest)

	tester.StopMockDatabase()
}

type RepoTester struct {
	repo *Repository
	test.BaseRepoTester
}

var (
	testUser = User{
		BaseEntity: domain.BaseEntity{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		UserName:  "username",
		Password:  "password",
		Email:     "email",
		LastLogin: time.Now(),
	}
	tester RepoTester
)

func (t2 *RepoTester) CreateTest(t *testing.T) {
	require.NoError(t, t2.repo.Create(testUser))
}

func (t2 *RepoTester) GetByIdTest(t *testing.T) {
	rt, err := t2.repo.GetById(testUser.ID)
	require.NoError(t, err)
	t.Log(rt)
}

func (t2 *RepoTester) UpdateTest(t *testing.T) {
	testUser.Email = "UPDATED"
	require.NoError(t, t2.repo.Update(testUser))
}

func (t2 *RepoTester) DeleteTest(t *testing.T) {
	require.NoError(t, t2.repo.Delete(testUser))
	_, err := t2.repo.GetById(testUser.ID)
	require.Error(t, err)
}
