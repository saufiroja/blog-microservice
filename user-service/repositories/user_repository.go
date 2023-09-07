package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/saufiroja/blog-microservice/user-service/interfaces"
	"github.com/saufiroja/blog-microservice/user-service/models/dto"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FindAllUsers(pagination *dto.Pagination) ([]dto.FindAllUsersDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var users []dto.FindAllUsersDTO

	// query
	query := `SELECT id, name, email, created_at FROM users ORDER BY id DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, query, pagination.Limit, pagination.Offset)
	if err != nil {
		return nil, err
	}

	// scan rows
	for rows.Next() {
		var user dto.FindAllUsersDTO
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) CountAllUsers() (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var count int32

	// db begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	// query
	query := `SELECT COUNT(*) FROM users`
	err = tx.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}

		return 0, err
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}

		return 0, err
	}

	return count, nil
}

func (r *userRepository) InsertUser(user *dto.InsertUserDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// db begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// query
	query := `INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = tx.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	// commit transaction
	return tx.Commit()
}
