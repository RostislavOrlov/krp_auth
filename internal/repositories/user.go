package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"krp_project/internal/dto"
	"krp_project/internal/entities"
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
	q := "SELECT * FROM Users WHERE email=$1 AND password=$2"
	row, err := repo.db.Query(context.Background(), q, user.Email, user.Password)
	if err != nil && err.Error() != "no row in result set" {
		return nil, err
	}

	defer row.Close()
	var usrDb *entities.User
	err = row.Scan(&usrDb.Id, &usrDb.LastName, &usrDb.FirstName, &usrDb.MiddleName, &usrDb.Email, &usrDb.Password)
	if err != nil {
		log.Fatalf("Unable to scan row: %v\n", err)
	}

	return usrDb, nil
}

func (repo *UserRepository) Register(user *dto.RegisterRequest) (*entities.User, error) {
	q := "INSERT INTO users(lastname, firstname, middlename, email, password) " +
		"VALUES($1, $2, $3, $4, $5) RETURNING *"

	row, err := repo.db.Query(context.Background(), q,
		user.LastName, user.FirstName, user.MiddleName,
		user.Email, user.Password)
	if err != nil && err.Error() != "no row in result set" {
		return nil, err
	}

	defer row.Close()
	var usrDb entities.User
	err = row.Scan(&usrDb.Id, &usrDb.LastName, &usrDb.FirstName, &usrDb.MiddleName, &usrDb.Email, &usrDb.Password)
	if err != nil {
		log.Fatalf("Unable to scan row: %v\n", err)
	}

	return &usrDb, nil
}

func (repo *UserRepository) UpdateRefreshToken() {

}
