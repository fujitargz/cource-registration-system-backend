package infra

import (
	"database/sql"
	"log"

	"github.com/fujitargz/cource-registration-system-backend/internal/domain"
)

type UserRepositoryInfra struct {
	*sql.DB
}

func NewUserRepositoryInfra(db *sql.DB) domain.UserRepository {
	return &UserRepositoryInfra{db}
}

func (r *UserRepositoryInfra) Save(ID string, passwordHash string, isAdmin bool) error {
	_, err := r.Exec("insert into users(id, password_hash, is_admin) values(?, ?, ?) on conflict(id) do update set password_hash=?", ID, passwordHash, isAdmin, passwordHash)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryInfra) Delete(ID string) error {
	_, err := r.Exec("delete from users where id is ?", ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryInfra) FindByID(ID string) (*domain.User, error) {
	user := &domain.User{}
	row := r.QueryRow("select id, password_hash, is_admin from users where id is ?", ID)
	err := row.Scan(&user.ID, &user.PasswordHash, &user.IsAdmin)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./trial-api.db")
	if err != nil {
		log.Println("failed to open database: ", err)
		return nil, err
	}
	_, err = db.Exec("create table if not exists users(id text primary key, password_hash text, is_admin text)")
	if err != nil {
		log.Println("failed to create table: ", err)
		return nil, err
	}
	return db, err
}
