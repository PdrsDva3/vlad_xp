package user

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"vlad_xp/internal/models"
	"vlad_xp/internal/repository"
)

type User struct {
	db *sqlx.DB
}

func InitUserRepo(db *sqlx.DB) repository.UserRepo {
	return User{
		db: db,
	}
}

func (u User) GetPwdByEmail(ctx context.Context, email string) (string, error) {
	var hashedPwd string
	row := u.db.QueryRowContext(ctx, `SELECT hashed_pwd FROM users WHERE users.email = $1`,
		email)

	err := row.Scan(&hashedPwd)
	if err != nil {
		return "", err
	}

	return hashedPwd, nil
}

func (u User) Get(ctx context.Context, id int) (*models.GetUser, error) {
	var user models.GetUser

	row := u.db.QueryRowContext(ctx, `SELECT username, surname,  phone FROM users WHERE users.id = $1`,
		id)

	err := row.Scan(&user.Name, user.Surname, &user.Phone)
	if err != nil {
		return nil, err
	}

	user.ID = id

	return &user, nil
}

func (u User) Delete(ctx context.Context, id int) error {
	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	res, err := tx.ExecContext(ctx, `DELETE FROM users WHERE users.id = $1;`, id)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return err
		}
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	if count != 1 {
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return errors.New("count error")
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (u User) Create(ctx context.Context, user models.CreateUser) (int, error) {
	tx, err := u.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, err
	}

	var id int

	row := tx.QueryRowContext(ctx, `INSERT INTO users (email, hashed_pwd) values ($1, $2) RETURNING id`, user.Email, user.HashedPassword)
	err = row.Scan(&id)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return 0, rbErr
		}
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}
