package repositories

import (
	"authorization/core/domain/entities"
	"authorization/core/domain/repositories/abstraction"
	"database/sql"
)

type userRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) abstraction.UserRepository {
	return &userRepository{connection}
}

func (repository *userRepository) SelectByLogin(login string) (*entities.User, error) {
	connection := repository.connection
	loginSelect, err := connection.Prepare("select * from users where login = $1")
	if err != nil {
		return nil, err
	}

	cursor, err := loginSelect.Query(login)
	if err != nil {
		return nil, err
	}

	var user *entities.User
	for cursor.Next() {
		err = cursor.Scan(user)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (repository *userRepository) Insert(user *entities.User) error {
	connection := repository.connection
	insert, err := connection.Prepare("insert into users(login, email, password) values($1,$2,$3)")
	if err != nil {
		return err
	}

	_, err = insert.Query(user.Login, user.Email, user.Password)
	return err
}

func (repository *userRepository) SelectByEmail(email string) (*entities.User, error) {
	return nil, nil
}
