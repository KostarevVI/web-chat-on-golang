// Package model as Database tables
package model

// UserData table
type UserData struct {
	ID              int
	PasswordHash    string
	FirstName       string
	LastName        string
	Email           string
	CreateDate      string
	UpdateDate      string
	FriendsChatOnly bool
	IsDeleted       bool
}
