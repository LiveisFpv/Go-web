package queries

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

func (q *Queries) GetAllAchievementCategory(ctx context.Context, filters map[string]string, rowCount, page int, search string) ([]*domain.AchievementCategory, int, error) {
	getAllCategory := `SELECT 
		id_category,
		achivments_type_category,
		score_category
	FROM category WHERE 1=1`
	countQuery := `SELECT COUNT(*) FROM category WHERE 1=1`

	var args []interface{}
	argCount := 1

	// Handle global search if provided
	if search != "" {
		getAllCategory += fmt.Sprintf(` AND (
			achivments_type_category ILIKE $%d OR 
			score_category::text ILIKE $%d
		)`, argCount, argCount)
		countQuery += fmt.Sprintf(` AND (
			achivments_type_category ILIKE $%d OR 
			score_category::text ILIKE $%d
		)`, argCount, argCount)
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
			getAllCategory += fmt.Sprintf(" AND %s ILIKE $%d", key, argCount)
			countQuery += fmt.Sprintf(" AND %s ILIKE $%d", key, argCount)
			args = append(args, "%"+value+"%")
			argCount++
		}
	}

	// Handle sorting
	if sortField != "" {
		// Validate sort field to prevent SQL injection
		validFields := map[string]string{
			"id_category":              "id_category",
			"achivments_type_category": "achivments_type_category",
			"score_category":           "score_category",
		}
		if field, ok := validFields[sortField]; ok {
			order := "ASC"
			if sortOrder == "desc" {
				order = "DESC"
			}
			getAllCategory += fmt.Sprintf(" ORDER BY %s %s", field, order)
		}
	}

	// Add pagination
	offset := (page - 1) * rowCount
	getAllCategory += fmt.Sprintf(" LIMIT %d OFFSET %d", rowCount, offset)

	rows, err := q.pool.Query(ctx, getAllCategory, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("can't query achievement categories: %w", err)
	}
	defer rows.Close()

	categories := []*domain.AchievementCategory{}
	for rows.Next() {
		category := &domain.AchievementCategory{}
		err := rows.Scan(
			&category.Id_category,
			&category.Achivments_type_category,
			&category.Score_category)
		if err != nil {
			return nil, 0, fmt.Errorf("can't scan achievement categories: %w", err)
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	// Get total count without pagination
	var count int
	err = q.pool.QueryRow(ctx, countQuery, args...).Scan(&count)
	if err != nil {
		return nil, 0, fmt.Errorf("can't count achievement categories: %w", err)
	}

	return categories, count, nil
}

func (q *Queries) CreateAchievementCategory(ctx context.Context, achivments_type_category string, score_category uint8) (*domain.AchievementCategory, error) {
	sqlStatement := `INSERT INTO category (achivments_type_category, score_category) VALUES ($1, $2) RETURNING id_category`
	var id uint64
	err := q.pool.QueryRow(ctx, sqlStatement, achivments_type_category, score_category).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("can't create achievement category: %w", err)
	}
	category, err := q.FindAchievementCategoryByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("can't find achievement category: %w", err)
	}
	return category, err
}

func (q *Queries) FindAchievementCategoryByID(ctx context.Context, id uint64) (*domain.AchievementCategory, error) {
	row := q.pool.QueryRow(ctx, `SELECT 
		id_category,
		achivments_type_category,
		score_category
	FROM category WHERE id_category = $1`, id)

	category := &domain.AchievementCategory{}
	err := row.Scan(&category.Id_category,
		&category.Achivments_type_category,
		&category.Score_category)
	if err != nil {
		return nil, fmt.Errorf("can't scan achievement category: %w", err)
	}

	return category, nil
}

func (q *Queries) UpdateAchievementCategoryByID(ctx context.Context, id_category uint64, achivments_type_category string, score_category uint8) (*domain.AchievementCategory, error) {
	sqlStatement := `UPDATE category SET achivments_type_category=$1, score_category=$2 WHERE id_category=$3 RETURNING id_category`
	err := q.pool.QueryRow(ctx, sqlStatement, achivments_type_category, score_category, id_category).Scan(&id_category)
	if err != nil {
		return nil, fmt.Errorf("can't update achievement category: %w", err)
	}
	category, err := q.FindAchievementCategoryByID(ctx, id_category)
	if err != nil {
		return nil, fmt.Errorf("can't find achievement category: %w", err)
	}
	return category, err
}

func (q *Queries) DeleteAchievementCategoryByID(ctx context.Context, id_category uint64) error {
	sqlStatement := `DELETE FROM category WHERE id_category=$1`
	_, err := q.pool.Exec(ctx, sqlStatement, id_category)
	if err != nil {
		return fmt.Errorf("can't delete achievement category: %w", err)
	}
	return nil
}

func (q *Queries) DeleteAchievementCategories(ctx context.Context, ids []uint64) error {
	sqlStatement := `DELETE FROM category WHERE id_category = ANY($1)`
	_, err := q.pool.Exec(ctx, sqlStatement, ids)
	if err != nil {
		return fmt.Errorf("can't delete achievement categories: %w", err)
	}
	return nil
}
