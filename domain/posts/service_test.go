package posts

import (
	"github.com/HarrisonWAffel/dbTrain/test"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

type ServiceTester struct {
	service Service
	test.BaseRepoTester
}

var tstr ServiceTester

func TestPostService(t *testing.T) {
	tstr = ServiceTester{}
	tstr.StartMockDatabase()
	db, err := gorm.Open(postgres.Open(tstr.DSN), nil)
	require.NoError(t, err)
	serv, err := NewService(db)
	require.NoError(t, err)
	tstr.service = serv

	t.Run("", testCreateNewPost)
	t.Run("", testGetPostsById)
	t.Run("", testUpdatePost)
	t.Run("", testDeletePost)

	tstr.StopMockDatabase()
}

func testCreateNewPost(t *testing.T) {
	require.NoError(t, tstr.service.CreateNewPost(testPost))
}

func testGetPostsById(t *testing.T) {
	x, err := tstr.service.GetPostById(testPost.ID)
	require.NoError(t, err)
	require.Equal(t, x.ID, testPost.ID)
	t.Log(x)
}

func testUpdatePost(t *testing.T) {
	testPost.Content = "NEW"
	require.NoError(t, tstr.service.UpdatePost(testPost))
	x, err := tstr.service.GetPostById(testPost.ID)
	require.NoError(t, err)
	t.Log(x)
}

func testDeletePost(t *testing.T) {
	require.NoError(t, tstr.service.DeletePost(testPost))
	_, err := tstr.service.GetPostById(testPost.ID)
	require.Error(t, err)
}
