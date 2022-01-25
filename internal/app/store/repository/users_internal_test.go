package sqlstore_test

import (
	"testing"

	sqlstore "github.com/anayks/golang-avito-trainee-tech-task/internal/app/store/repository"
	"github.com/stretchr/testify/assert"
)

func TestUser_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, "../../../../.env")
	defer db.Close()
	s := sqlstore.New(db)
	_, user, _, teardownFiller := sqlstore.DBFiller(t, db, teardown)
	defer teardownFiller()

	t.Run("TestUser_Create", func(t *testing.T) {
		newUser, err := s.RepositoryUsers.Create("heyboy")
		assert.NoError(t, err)
		assert.Equal(t, user.ID+1, newUser.ID)
	})
}
