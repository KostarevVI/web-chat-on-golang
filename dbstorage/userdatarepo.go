// Package dbstorage
package dbstorage

import "web-chat-on-golang/internal/app/model"

// UserDataRepo repository structure for UserData table
type UserDataRepo struct {
	dbstorage *DbStorage
}

// Create entry in UserData table
func (r *UserDataRepo) Create(u *model.UserData) (*model.UserData, error) {
	if err := r.dbstorage.db.QueryRow(
		"INSERT INTO user_data (email, password_hash, first_name, last_name, "+
			"friends_chat_only, is_deleted) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		u.Email,
		u.PasswordHash,
		u.FirstName,
		u.LastName,
		u.FriendsChatOnly,
		u.IsDeleted,
	).Scan(&u.ID); err != nil {
		return nil, err
	}
	return u, nil
}

// FindUser with email in UserData table
func (r *UserDataRepo) FindUser(email string) (*model.UserData, error) {
	u := &model.UserData{}
	if err := r.dbstorage.db.QueryRow(
		"SELECT id, email, password_hash, first_name, last_name, create_date, update_date, friends_chat_only, "+
			"is_deleted FROM user_data WHERE email = $1", email).Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
		&u.FirstName,
		&u.LastName,
		&u.CreateDate,
		&u.UpdateDate,
		&u.FriendsChatOnly,
		&u.IsDeleted,
	); err != nil {
		return nil, err
	}
	return u, nil
}
