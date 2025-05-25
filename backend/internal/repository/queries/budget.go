package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

func (q *Queries) GetAllBudget(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.Budget, int, error) {
	getAllBudget := `SELECT 
		type_scholarship_budget,
		name_semester,
		size_budget,
		id_budget
	FROM budget WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM budget WHERE 1=1`

	var args []interface{}
	argCount := 1

	// Handle global search if provided
	if search != "" {
		getAllBudget += fmt.Sprintf(` AND (
			type_scholarship_budget ILIKE $%d OR 
			name_semester ILIKE $%d OR 
			size_budget::text ILIKE $%d
		)`, argCount, argCount, argCount)
		countQuery += fmt.Sprintf(` AND (
			type_scholarship_budget ILIKE $%d OR 
			name_semester ILIKE $%d OR 
			size_budget::text ILIKE $%d
		)`, argCount, argCount, argCount)
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
			getAllBudget += fmt.Sprintf(" AND %s ILIKE $%d", key, argCount)
			countQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, argCount)
			args = append(args, "%"+value+"%")
			argCount++
		}
	}

	// Handle sorting
	if sortField != "" {
		// Validate sort field to prevent SQL injection
		validFields := map[string]string{
			"type_scholarship_budget": "type_scholarship_budget",
			"name_semester":           "name_semester",
			"size_budget":             "size_budget",
			"id_budget":               "id_budget",
		}
		if field, ok := validFields[sortField]; ok {
			order := "ASC"
			if sortOrder == "desc" {
				order = "DESC"
			}
			getAllBudget += fmt.Sprintf(" ORDER BY %s %s", field, order)
		}
	}

	// Add pagination
	offset := (page - 1) * rowCount
	getAllBudget += fmt.Sprintf(" LIMIT %d OFFSET %d", rowCount, offset)

	rows, err := q.pool.Query(ctx, getAllBudget, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("can't query budgets: %w", err)
	}
	defer rows.Close()

	budgets := []*domain.Budget{}
	for rows.Next() {
		budget := &domain.Budget{}
		err := rows.Scan(
			&budget.Type_scholarship_budget,
			&budget.Name_semester,
			&budget.Size_budget,
			&budget.Id_budget)
		if err != nil {
			return nil, 0, fmt.Errorf("can't scan budgets: %w", err)
		}
		budgets = append(budgets, budget)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	// Получаем общее количество строк без пагинации
	var count int
	err = q.pool.QueryRow(ctx, countQuery, args...).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("can't count budgets: %w", err)
	}

	return budgets, count, nil
}

func (q *Queries) CreateBudget(ctx context.Context, type_scholarship_budget string, name_semester string, size_budget float64) (*domain.Budget, error) {
	sqlStatement := `INSERT INTO budget (type_scholarship_budget, name_semester, size_budget) VALUES ($1, $2, $3) RETURNING id_budget`
	var id int
	err := q.pool.QueryRow(ctx, sqlStatement, type_scholarship_budget, name_semester, size_budget).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("can't create budget: %w", err)
	}
	budget, err := q.FindBudgetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't find budget: %w", err)
	}
	return budget, err
}

func (q *Queries) FindBudgetByID(ctx context.Context, id int) (*domain.Budget, error) {
	row := q.pool.QueryRow(ctx, `SELECT 
		type_scholarship_budget,
		name_semester,
		size_budget,
		id_budget
	FROM budget WHERE id_budget = $1`, id)

	budget := &domain.Budget{}
	err := row.Scan(
		&budget.Type_scholarship_budget,
		&budget.Name_semester,
		&budget.Size_budget,
		&budget.Id_budget)
	if err != nil {
		return nil, fmt.Errorf("can't scan budget: %w", err)
	}

	return budget, nil
}

func (q *Queries) UpdateBudgetByID(ctx context.Context, id_budget int, type_scholarship_budget string, name_semester string, size_budget float64) (*domain.Budget, error) {
	sqlStatement := `UPDATE budget SET type_scholarship_budget=$1, name_semester=$2, size_budget=$3 WHERE id_budget=$4 RETURNING id_budget`
	err := q.pool.QueryRow(ctx, sqlStatement, type_scholarship_budget, name_semester, size_budget, id_budget).Scan(&id_budget)
	if err != nil {
		return nil, fmt.Errorf("can't update budget: %w", err)
	}
	budget, err := q.FindBudgetByID(ctx, id_budget)
	if err != nil {
		return nil, fmt.Errorf("can't find budget: %w", err)
	}
	return budget, err
}

func (q *Queries) DeleteBudgetByID(ctx context.Context, id_budget int) error {
	sqlStatement := `DELETE FROM budget WHERE id_budget=$1`
	_, err := q.pool.Exec(ctx, sqlStatement, id_budget)
	if err != nil {
		return fmt.Errorf("can't delete budget: %w", err)
	}
	return nil
}

func (q *Queries) DeleteBudgets(ctx context.Context, ids []int) error {
	sqlStatement := `DELETE FROM budget WHERE id_budget = ANY($1)`
	_, err := q.pool.Exec(ctx, sqlStatement, ids)
	if err != nil {
		return fmt.Errorf("can't delete budgets: %w", err)
	}
	return nil
}
