package model

import (
	"time"
)

// User is the structure which holds one user from the database.
type TestUser struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Password  string `json:"-"`
	Active    bool
	IsAdmin   bool
	CreatedAt time.Time
	UpdatedAt time.Time
	Plan      *Plan
}

// GetAll returns a slice of all users, sorted by last name
func (u *TestUser) GetAll() ([]*User, error) {
	var users []*User

	user := User{
		ID:        1,
		Email:     "admin@mail.com",
		FirstName: "Admin",
		LastName:  "Mk11",
		Password:  "123",
		Active:    true,
		IsAdmin:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users = append(users, &user)

	return users, nil
}

// GetByEmail returns one user by email
func (u *TestUser) GetByEmail(email string) (*User, error) {
	user := User{
		ID:        1,
		Email:     "admin@mail.com",
		FirstName: "Admin",
		LastName:  "Mk11",
		Password:  "123",
		Active:    true,
		IsAdmin:   true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return &user, nil
}

// GetOne returns one user by id
func (u *TestUser) GetOne(id int) (*User, error) {
	return u.GetByEmail("")
}

// Update updates one user in the database, using the information
// stored in the receiver u
func (u *TestUser) Update() error {
	return nil
}

// Delete deletes one user from the database, by User.ID
func (u *TestUser) Delete() error {
	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *TestUser) DeleteByID(id int) error {
	return nil
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (u *TestUser) Insert(user User) (int, error) {
	return 2, nil
}

// ResetPassword is the method we will use to change a user's password.
func (u *TestUser) ResetPassword(password string) error {
	return nil
}

// PasswordMatches uses Go's bcrypt package to compare a user supplied password
// with the hash we have stored for a given user in the database. If the password
// and hash match, we return true; otherwise, we return false.
func (u *TestUser) PasswordMatches(plainText string) (bool, error) {
	return true, nil
}
