package repository

import (
	"database/sql"
	"libretaxi/objects"
	"log"
)

type Repository struct {
	db *sql.DB
}

func (repo *Repository) FindUser(userId int64) *objects.User {
	user := &objects.User{}

	rows, err := repo.db.Query(`select "userId", "menuId" from users where "userId"=$1 limit 1`, userId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cnt := 0
	for rows.Next() {
		cnt++
		rows.Scan(&user.UserId, &user.MenuId)
	}

	if cnt == 0 {
		return nil
	}

	return user
}

func NewRepository(db *sql.DB) *Repository {
	repo := &Repository{db: db}
	return repo
}