package repositories

import (
	"context"
	"database/sql"
	mdldomain "rest-project/app/models/mdl-domain"
)

type UserRepo interface {
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*mdldomain.User, error)
	Save(tx *sql.Tx, u *mdldomain.User) error
}

type UserRepoImpl struct {
	DB *sql.DB
}

func NewUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{
		DB: db,
	}
}

// FindByEmail implements UserRepo.
func (u *UserRepoImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (*mdldomain.User, error) {
	query := "SELECT email FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	user := mdldomain.User{}
	if rows.Next() {
		err = rows.Scan(&user.Email)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, sql.ErrNoRows
}

// Save implements UserRepo.
func (*UserRepoImpl) Save(tx *sql.Tx, u *mdldomain.User) error {
	query := "INSERT INTO users (name, email, password, created_at) VALUES(?,?,?,?)"
	_, err := tx.Exec(query, u.Name, u.Email, u.Password, u.CreatedAt)
	return err
}
