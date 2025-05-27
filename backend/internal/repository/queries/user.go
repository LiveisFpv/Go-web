package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

// TODO work with database
func (q *Queries) Login(ctx context.Context, login, password string) (*domain.User, error) {
	sqlStatement := `SELECT * FROM "user" WHERE user_login = $1`
	row := q.pool.QueryRow(ctx, sqlStatement, login)
	user := &domain.User{}
	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Student_id,
		&user.Role,
		&user.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("can't found the user: %w", err)
	}
	return user, nil
}
func (q *Queries) Register(ctx context.Context, email, login, password string) error {
	sqlStatement := `INSERT INTO "user" (user_login, user_email, user_password) 
		VALUES ($1, $2, $3)`
	_, err := q.pool.Exec(ctx, sqlStatement, login, email, password)
	return err
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	sqlStatement := `SELECT * FROM "user" WHERE user_email = $1`
	row := q.pool.QueryRow(ctx, sqlStatement, email)
	user := &domain.User{}
	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Student_id,
		&user.Role,
		&user.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("can't find user by email: %w", err)
	}
	return user, nil
}

func (q *Queries) DeleteUser(ctx context.Context, userId int64) error {
	sqlStatement := `DELETE FROM "user" WHERE user_id = $1`
	_, err := q.pool.Exec(ctx, sqlStatement, userId)
	if err != nil {
		return fmt.Errorf("can't delete the user: %w", err)
	}
	return nil
}

func (q *Queries) UpdateUser(ctx context.Context, user *domain.User) error {
	sqlStatement := `UPDATE "user" SET user_email = $1, user_login = $2, user_student_id = $3, user_role = $4 WHERE user_id = $5`
	_, err := q.pool.Exec(ctx, sqlStatement, user.Email, user.Login, user.Student_id, user.Role, user.Id)
	if err != nil {
		return fmt.Errorf("can't update the user: %w", err)
	}
	return nil
}

func (q *Queries) GetUserByID(ctx context.Context, userId int64) (*domain.User, error) {
	sqlStatement := `SELECT * FROM "user" WHERE user_id = $1`
	row := q.pool.QueryRow(ctx, sqlStatement, userId)
	user := &domain.User{}
	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Email,
		&user.Student_id,
		&user.Role,
		&user.Password,
	)
	if err != nil {
		return nil, fmt.Errorf("can't find user: %w", err)
	}
	return user, nil
}

func (q *Queries) GetAllUsers(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.User, int, error) {
	getAllUsers := `SELECT * FROM "user" WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM "user" WHERE 1=1`

	var args []interface{}
	argCount := 1

	// Handle global search if provided
	if search != "" {
		getAllUsers += fmt.Sprintf(` AND (
			user_id::text ILIKE $%d OR 
			user_login ILIKE $%d OR 
			user_email ILIKE $%d OR 
			user_role ILIKE $%d
		)`, argCount, argCount, argCount, argCount)
		countQuery += fmt.Sprintf(` AND (
			user_id::text ILIKE $%d OR 
			user_login ILIKE $%d OR 
			user_email ILIKE $%d OR 
			user_role ILIKE $%d
		)`, argCount, argCount, argCount, argCount)
		args = append(args, "%"+search+"%")
		argCount++
	}

	// Extract sorting parameters
	sortField := filters["sort"]
	sortOrder := filters["order"]
	delete(filters, "sort")
	delete(filters, "order")

	// Handle specific filters
	for key, value := range filters {
		if value != "" {
			getAllUsers += fmt.Sprintf(" AND %s ILIKE $%d", key, argCount)
			countQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, argCount)
			args = append(args, "%"+value+"%")
			argCount++
		}
	}

	// Handle sorting
	if sortField != "" {
		// Validate sort field to prevent SQL injection
		validFields := map[string]string{
			"user_id":    "user_id",
			"user_login": "user_login",
			"user_email": "user_email",
			"user_role":  "user_role",
		}
		if field, ok := validFields[sortField]; ok {
			order := "ASC"
			if sortOrder == "desc" {
				order = "DESC"
			}
			getAllUsers += fmt.Sprintf(" ORDER BY %s %s", field, order)
		}
	}

	// Add pagination
	offset := (page - 1) * rowCount
	getAllUsers += fmt.Sprintf(" LIMIT %d OFFSET %d", rowCount, offset)

	rows, err := q.pool.Query(ctx, getAllUsers, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("can't query users: %w", err)
	}
	defer rows.Close()

	users := []*domain.User{}
	for rows.Next() {
		user := &domain.User{}
		err := rows.Scan(
			&user.Id,
			&user.Login,
			&user.Email,
			&user.Student_id,
			&user.Role,
			&user.Password,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("can't scan users: %w", err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	// Get total count without pagination
	var count int
	err = q.pool.QueryRow(ctx, countQuery, args...).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("can't count users: %w", err)
	}

	return users, count, nil
}

func (q *Queries) CreateUser(ctx context.Context, email, login, password string, studentId *int, role string) (*domain.User, error) {
	sqlStatement := `INSERT INTO "user" (user_login, user_email, user_password, user_student_id, user_role) VALUES ($1, $2, $3, $4, $5) RETURNING user_id`
	var userId int64
	err := q.pool.QueryRow(ctx, sqlStatement, login, email, password, studentId, role).Scan(&userId)
	if err != nil {
		return nil, fmt.Errorf("can't create user: %w", err)
	}
	return q.GetUserByID(ctx, userId)
}

func (q *Queries) DeleteUsers(ctx context.Context, ids []int64) error {
	for _, id := range ids {
		err := q.DeleteUser(ctx, id)
		if err != nil {
			return fmt.Errorf("can't delete user %d: %w", id, err)
		}
	}
	return nil
}
