package posts

import (
	"github.com/HarrisonWAffel/dbTrain/domain"
	"github.com/HarrisonWAffel/dbTrain/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestPostRepository(t *testing.T) {
	tester = RepoTester{}
	tester.StartMockDatabase()
	db, err := gorm.Open(postgres.Open(tester.DSN), nil)
	require.NoError(t, err)
	repo, err := NewPostsRepository(db)
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
	testPost = Post{
		BaseEntity: domain.BaseEntity{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		Content: "CONTENT",
		Private: false,
	}
	tester RepoTester
)

func (t2 *RepoTester) CreateTest(t *testing.T) {
	require.NoError(t, t2.repo.Create(testPost))
}

func (t2 *RepoTester) GetByIdTest(t *testing.T) {
	rt, err := t2.repo.GetById(testPost.ID)
	require.NoError(t, err)
	t.Log(rt)
}

func (t2 *RepoTester) UpdateTest(t *testing.T) {
	testPost.Content = "UPDATED"
	require.NoError(t, t2.repo.Update(testPost))
}

func (t2 *RepoTester) DeleteTest(t *testing.T) {
	require.NoError(t, t2.repo.Delete(testPost))
	_, err := t2.repo.GetById(testPost.ID)
	require.Error(t, err)
}
