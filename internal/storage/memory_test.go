package storage

import "testing"

func TestAddUser(t *testing.T) {
	repo := New()
	err := repo.AddUser("Alice")
	if err != nil {
		t.Fatalf("could not add user, error: %v", err)
	}

	user, err := repo.GetUser("Alice")
	if err != nil {
		t.Fatalf("could not get user, error: %v", err)
	}

	if user.Name != "Alice" {
		t.Fatalf("expected user name to be Alice, got %s", user.Name)
	}
}
