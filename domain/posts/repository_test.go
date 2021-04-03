package posts

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPostRepository(t *testing.T) {
	repo, err := NewPostsRepository()
	require.NoError(t, err)
	tester := Tester{repo: repo}

	t.Run("GetById Test", tester.GetByIdTest)
	t.Run("Create Test", tester.CreateTest)
	t.Run("Update Test", tester.UpdateTest)
	t.Run("Delete Test", tester.DeleteTest)
}

type Tester struct {
	repo *Repository
}

func (t2 *Tester) GetByIdTest(t *testing.T) {
	panic("implement me")
}

func (t2 *Tester) CreateTest(t *testing.T) {
	panic("implement me")
}

func (t2 *Tester) UpdateTest(t *testing.T) {
	panic("implement me")
}

func (t2 *Tester) DeleteTest(t *testing.T) {
	panic("implement me")
}
