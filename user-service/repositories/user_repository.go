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

func (r *userRepository) FindUsersByEmail(email string) (*dto.FindUsersDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user dto.FindUsersDTO

	query := `SELECT id, name, email, password, created_at FROM users WHERE email = $1`
	err := r.db.QueryRowContext(ctx,
		query,
		email,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindUsersByID(id string) (*dto.FindUsersDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user dto.FindUsersDTO

	query := `SELECT id, name, email, created_at FROM users WHERE id = $1`
	err := r.db.QueryRowContext(ctx,
		query,
		id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(id string, user *dto.UpdateUserDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// db begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// query
	query := `UPDATE users SET name = $1, email = $2, updated_at = $3 WHERE id = $4`
	_, err = tx.ExecContext(ctx, query, user.Name, user.Email, user.UpdatedAt, id)
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

func (r *userRepository) DeleteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// db begin transaction
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	// query
	query := `DELETE FROM users WHERE id = $1`
	_, err = tx.ExecContext(ctx, query, id)
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
