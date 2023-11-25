package postgresql

import (
	"authorization/internal/domain/entities"
	"authorization/internal/domain/useCases"
	"context"
	"database/sql"
)

type userRepository struct {
	connection *sql.Conn
}

func NewUserRepository(connection *sql.Conn) useCases.UserRepository {
	return &userRepository{connection}
}

func (repository *userRepository) SelectByLogin(login string) (*entities.User, error) {
	connection := repository.connection
	loginSelect, err := connection.PrepareContext(context.TODO(), "select * from users where login = $1")
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
	insert, err := connection.PrepareContext(context.TODO(), "insert into users(login, email, password) values($1,$2,$3)")
	if err != nil {
		return err
	}

	_, err = insert.Query(user.Login, user.Email, user.Password)
	return err
}

func (repository *userRepository) SelectByEmail(email string) (*entities.User, error) {
	return nil, nil
}
