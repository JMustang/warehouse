package user

import (
	"testing"

	"github.com/JMustang/warehouse/types"
)

func TestUserServiceHandles(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
}

type mockUserStore struct{}

// GetUserByID implements types.UserStore.
func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserById(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(*types.User) error {
	return nil
}
