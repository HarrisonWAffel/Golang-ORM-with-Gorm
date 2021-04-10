package user

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

func TestUserService(t *testing.T) {
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
	require.NoError(t, tstr.service.SaveUser(testUser))
}

func testGetPostsById(t *testing.T) {
	x, err := tstr.service.GetUserById(testUser.ID)
	require.NoError(t, err)
	require.Equal(t, x.ID, testUser.ID)
	t.Log(x)
}

func testUpdatePost(t *testing.T) {
	testUser.Email = "NEW"
	require.NoError(t, tstr.service.UpdateUser(testUser))
	x, err := tstr.service.GetUserById(testUser.ID)
	require.NoError(t, err)
	t.Log(x)
}

func testDeletePost(t *testing.T) {
	require.NoError(t, tstr.service.DeleteUser(testUser))
	_, err := tstr.service.GetUserById(testUser.ID)
	require.Error(t, err)
}
