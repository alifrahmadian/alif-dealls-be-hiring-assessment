package repositories

import (
	"database/sql"

	"github.com/alifrahmadian/alif-dealls-be-hiring-assessment/internal/models"
	e "github.com/alifrahmadian/alif-dealls-be-hiring-assessment/pkg/errors"
)

type UserRepository interface {
	GetUser(username string) (*models.User, error)
	GetUserByID(userID int64) (*models.User, error)
}

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository{
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) GetUser(username string) (*models.User, error) {
	query := `
		SELECT id, username, password, role_id, base_salary, created_at, updated_at FROM users
		WHERE username = $1
	`

	user := &models.User{}

	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.RoleID, &user.BaseSalary, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserByID(userID int64) (*models.User, error) {
	query := `
	SELECT id, username, password, role_id, base_salary, created_at, updated_at FROM users
	WHERE id = $1
`

	user := &models.User{}	

	err := r.DB.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Password, &user.RoleID, &user.BaseSalary, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, e.ErrUserNotFound
		}

		return nil, err
	}

	return user, nil
}

