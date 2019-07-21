package models

import "database/sql"

type Users struct {
	ID       string
	Fullname string
}

func GetUser(db *sql.DB, query string) ([]*Users, error) {
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	user := make([]*Users, 0)

	for rows.Next() {
		rawUser := new(Users)
		err := rows.Scan(&rawUser.ID, &rawUser.Fullname)
		if err != nil {
			return nil, err
		}
		user = append(user, rawUser)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return user, nil
}
