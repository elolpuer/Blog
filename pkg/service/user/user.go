package user

import (
	"database/sql"

	"github.com/elolpuer/blog/pkg/models"
)

//GetAll ...
func GetAll(db *sql.DB) ([]*models.SessionUser, error) {
	rows, err := db.Query("SELECT id, username FROM users")
	if err != nil {
		return nil, err
	}
	var users = make([]*models.SessionUser, 0)
	defer rows.Close()
	for rows.Next() {
		user := new(models.SessionUser)
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
