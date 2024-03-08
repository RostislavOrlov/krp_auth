package repositories

import (
	"context"
	"github.com/IlyaZayats/auth/internal/dto"
	"github.com/IlyaZayats/auth/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) (*UserRepository, error) {
	return &UserRepository{
		db: db,
	}, nil
}

func (repo *UserRepository) Auth(user *dto.AuthRequest) (*entities.User, error) {
	q := "SELECT * FROM users WHERE email=$1 AND pswd=$2"
	row, err := repo.db.Query(context.Background(), q, user.Email, user.Password)
	if err != nil && err.Error() != "no row in result set" {
		return nil, err
	}

	defer row.Close()
	var usrDb entities.User
	for row.Next() {
		err = row.Scan(&usrDb.Id, &usrDb.LastName, &usrDb.FirstName,
			&usrDb.MiddleName, &usrDb.Email, &usrDb.Password,
			&usrDb.Passport, &usrDb.Inn, &usrDb.Snils, &usrDb.Birthday, &usrDb.Role)
		if err != nil {
			log.Fatalf("Unable to scan row: %v\n", err)
		}
	}

	return &usrDb, nil
}

func (repo *UserRepository) Register(user *dto.RegisterRequest, password string) (*entities.User, error) {
	q := "INSERT INTO users (lastname, firstname, middlename, email, pswd, passport, inn, snils, birthday, role)" + //(lastname, firstname, middlename, email, password, role)
		"VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *"

	row, err := repo.db.Query(context.Background(), q,
		user.LastName, user.FirstName, user.MiddleName,
		user.Email, password, user.Passport, user.Inn, user.Snils, user.Birthday, user.Role)
	if err != nil && err.Error() != "no row in result set" {
		return nil, err
	}

	defer row.Close()
	var usrDb entities.User
	for row.Next() {
		err = row.Scan(&usrDb.Id, &usrDb.LastName, &usrDb.FirstName,
			&usrDb.MiddleName, &usrDb.Email, &usrDb.Password,
			&usrDb.Passport, &usrDb.Inn, &usrDb.Snils, &usrDb.Birthday, &usrDb.Role)
		if err != nil {
			log.Fatalf("Unable to scan row: %v\n", err)
		}
	}
	return &usrDb, nil
}
