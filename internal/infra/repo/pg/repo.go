package pg

import (
	"context"
	"database/sql"
	"fmt"
	"math"

	"github.com/google/uuid"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/model"
	"github.com/jobquestvault/platform-go-challenge/internal/sys/errors"
)

func (ar *AssetRepo) GetAssets(ctx context.Context, page, size int) ([]model.Asset[model.Favable], int, error) {
	db, ok := ar.PgDB()
	if !ok {
		return nil, 0, NoConnectionErr
	}

	offset := (page - 1) * size

	query := `
	SELECT 
		coalesce(a.asset_id, c.id, i.id, u.id) AS asset_id,
-- 		coalesce(a.asset_type, 'chart', 'insight', 'audience') AS asset_type,
    CASE
        WHEN c.id IS NOT NULL THEN 'chart'
        WHEN i.id IS NOT NULL THEN 'insight'
        WHEN u.id IS NOT NULL THEN 'audience'
        ELSE NULL
    END AS asset_type,
		coalesce(a.name, c.name, i.name, u.name) AS name,
		a.description,
		c.title,
		c.x_axis_title,
		c.y_axis_title,
		c.data,
		i.text,
		i.topic,
		u.gender,
		u.birth_country,
		u.age_group,
		u.hours_spent_on_social,
		u.num_purchases_last_month
	FROM
		ak.assets a
	FULL OUTER JOIN
		ak.charts c ON a.asset_id = c.id AND a.asset_type = 'chart'
	FULL OUTER JOIN
		ak.insights i ON a.asset_id = i.id AND a.asset_type = 'insight'
	FULL OUTER JOIN
		ak.audiences u ON a.asset_id = u.id AND a.asset_type = 'audience'
	ORDER BY
		asset_id
	LIMIT $1
	OFFSET $2
`

	rows, err := db.QueryContext(ctx, query, size, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	assets := make([]model.Asset[model.Favable], 0)

	for rows.Next() {
		var (
			assetID                       string
			assetType                     string
			name                          sql.NullString
			description                   sql.NullString
			chartTitle                    sql.NullString
			chartXAxisTitle               sql.NullString
			chartYAxisTitle               sql.NullString
			chartData                     []uint8
			insightText                   sql.NullString
			insightTopic                  sql.NullString
			audienceGender                sql.NullString
			audienceBirthCountry          sql.NullString
			audienceAgeGroup              sql.NullString
			audienceHoursSpentOnSocial    sql.NullInt64
			audienceNumPurchasesLastMonth sql.NullInt64
		)

		err := rows.Scan(
			&assetID,
			&assetType,
			&name,
			&description,
			&chartTitle,
			&chartXAxisTitle,
			&chartYAxisTitle,
			&chartData,
			&insightText,
			&insightTopic,
			&audienceGender,
			&audienceBirthCountry,
			&audienceAgeGroup,
			&audienceHoursSpentOnSocial,
			&audienceNumPurchasesLastMonth,
		)
		if err != nil {
			return nil, 0, err
		}

		switch assetType {
		case "chart":
			chart := model.Chart{
				ID:         model.ID{ID: assetID, Name: name.String},
				Title:      chartTitle.String,
				XAxisTitle: chartXAxisTitle.String,
				YAxisTitle: chartYAxisTitle.String,
				Data:       chartData,
			}
			assets = append(assets, model.NewAsset(assetID, name.String, description.String, assetType, chart))
		case "insight":
			insight := model.Insight{
				ID:    model.ID{ID: assetID, Name: name.String},
				Text:  insightText.String,
				Topic: insightTopic.String,
			}
			assets = append(assets, model.NewAsset(assetID, name.String, description.String, assetType, insight))
		case "audience":
			audience := model.Audience{
				ID:                    model.ID{ID: assetID, Name: name.String},
				Gender:                audienceGender.String,
				BirthCountry:          audienceBirthCountry.String,
				AgeGroup:              audienceAgeGroup.String,
				HoursSpentOnSocial:    int(audienceHoursSpentOnSocial.Int64),
				NumPurchasesLastMonth: int(audienceNumPurchasesLastMonth.Int64),
			}
			assets = append(assets, model.NewAsset(assetID, name.String, description.String, assetType, audience))
		}
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	countQuery := `
		SELECT COUNT(*)
		FROM ak.assets
	`

	var totalCount int
	err = db.QueryRowContext(ctx, countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(size)))

	return assets, totalPages, nil
}

func (ar *AssetRepo) GetFaved(ctx context.Context, userID string, page, size int) ([]model.Asset[model.Favable], int, error) {
	db, ok := ar.PgDB()
	if !ok {
		return nil, 0, NoConnectionErr
	}

	offset := (page - 1) * size

	query := `
	SELECT 
		coalesce(a.asset_id, c.id, i.id, u.id) AS asset_id,
    CASE
        WHEN c.id IS NOT NULL THEN 'chart'
        WHEN i.id IS NOT NULL THEN 'insight'
        WHEN u.id IS NOT NULL THEN 'audience'
        ELSE NULL
    END AS asset_type,
		coalesce(a.name, c.name, i.name, u.name) AS name,
		a.description,
		c.title,
		c.x_axis_title,
		c.y_axis_title,
		c.data,
		i.text,
		i.topic,
		u.gender,
		u.birth_country,
		u.age_group,
		u.hours_spent_on_social,
		u.num_purchases_last_month
	FROM
		ak.assets a
	LEFT JOIN 
		ak.charts c ON a.asset_id = c.id AND a.asset_type = 'chart'
	LEFT JOIN
		ak.insights i ON a.asset_id = i.id AND a.asset_type = 'insight'
	LEFT JOIN
		ak.audiences u ON a.asset_id = u.id AND a.asset_type = 'audience'
	WHERE a.user_id = $1
	ORDER BY
		asset_id
	LIMIT $2
	OFFSET $3
`

	rows, err := db.QueryContext(ctx, query, userID, size, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	assets := make([]model.Asset[model.Favable], 0)

	for rows.Next() {
		var (
			assetID                       string
			assetType                     string
			name                          sql.NullString
			description                   sql.NullString
			chartTitle                    sql.NullString
			chartXAxisTitle               sql.NullString
			chartYAxisTitle               sql.NullString
			chartData                     []uint8
			insightText                   sql.NullString
			insightTopic                  sql.NullString
			audienceGender                sql.NullString
			audienceBirthCountry          sql.NullString
			audienceAgeGroup              sql.NullString
			audienceHoursSpentOnSocial    sql.NullInt64
			audienceNumPurchasesLastMonth sql.NullInt64
		)

		err := rows.Scan(
			&assetID,
			&assetType,
			&name,
			&description,
			&chartTitle,
			&chartXAxisTitle,
			&chartYAxisTitle,
			&chartData,
			&insightText,
			&insightTopic,
			&audienceGender,
			&audienceBirthCountry,
			&audienceAgeGroup,
			&audienceHoursSpentOnSocial,
			&audienceNumPurchasesLastMonth,
		)
		if err != nil {
			return nil, 0, err
		}

		switch assetType {
		case "chart":
			chart := model.Chart{
				ID:         model.ID{ID: assetID, Name: name.String},
				Title:      chartTitle.String,
				XAxisTitle: chartXAxisTitle.String,
				YAxisTitle: chartYAxisTitle.String,
				Data:       chartData,
			}
			assets = append(assets, model.NewAsset(assetID, name.String, description.String, assetType, chart))
		case "insight":
			insight := model.Insight{
				ID:    model.ID{ID: assetID, Name: name.String},
				Text:  insightText.String,
				Topic: insightTopic.String,
			}
			assets = append(assets, model.NewAsset(assetID, name.String, description.String, assetType, insight))
		case "audience":
			audience := model.Audience{
				ID:                    model.ID{ID: assetID, Name: name.String},
				Gender:                audienceGender.String,
				BirthCountry:          audienceBirthCountry.String,
				AgeGroup:              audienceAgeGroup.String,
				HoursSpentOnSocial:    int(audienceHoursSpentOnSocial.Int64),
				NumPurchasesLastMonth: int(audienceNumPurchasesLastMonth.Int64),
			}
			assets = append(assets, model.NewAsset(assetID, name.String, description.String, assetType, audience))
		}
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	countQuery := `
		SELECT COUNT(*)
		FROM ak.assets
	`

	var totalCount int
	err = db.QueryRowContext(ctx, countQuery).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(size)))

	return assets, totalPages, nil
}

func (ar *AssetRepo) AddFav(ctx context.Context, userID, assetType, ID string) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionErr
	}

	tableName, err := ar.assetTable(assetType)
	if err != nil {
		return UnsupportedAssetErr
	}

	assetName := fmt.Sprintf("COALESCE((SELECT name FROM %s WHERE id = $3), $4 || '-' || $5 || '-' || RIGHT($5, 8))", tableName)
	assetDesc := fmt.Sprintf("CONCAT('⭐ ', COALESCE((SELECT name FROM %s WHERE id = $3), $4 || '-' || $5 || '-' || RIGHT($5, 8)))", tableName)

	query := `
		INSERT INTO assets (id, user_id, asset_id, asset_type, name, description)
		VALUES ($1, $2, $3, $4::text, ` + assetName + `, ` + assetDesc + `)
	`

	_, err = db.ExecContext(ctx, query, uuid.New().String(), userID, ID, assetType, ID)
	if err != nil {
		return err
	}

	return nil
}

func (ar *AssetRepo) UpdateFav(ctx context.Context, userID, assetType, ID, name, description string) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionErr
	}

	query := `
		UPDATE assets
		SET name = $1, description = $2
		WHERE user_id = $3 AND asset_type = $4 AND asset_id = $5
	`

	_, err := db.ExecContext(ctx, query, name, description, userID, assetType, ID)
	if err != nil {
		return err
	}

	return nil
}

func (ar *AssetRepo) RemoveFav(ctx context.Context, userID, assetType, ID string) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionErr
	}

	query := `
		DELETE FROM assets
		WHERE user_id = $1 AND asset_type = $2 AND asset_id = $3
	`

	_, err := db.ExecContext(ctx, query, userID, assetType, ID)
	if err != nil {
		return err
	}

	return nil
}

// Helper methods
func (ar *AssetRepo) validateAssetType(assetType string) error {
	validAssetTypes := map[string]bool{
		"chart":    true,
		"insight":  true,
		"audience": true,
	}

	if !validAssetTypes[assetType] {
		msg := fmt.Sprintf("invalid asset type: %s", assetType)
		return errors.NewError(msg)
	}

	return nil
}

func (ar *AssetRepo) assetTable(assetType string) (table string, err error) {
	switch assetType {
	case "chart":
		return "charts", nil
	case "insight":
		return "insights", nil
	case "audience":
		return "audiences", nil
	default:
		msg := fmt.Sprintf("invalid asset type: %s", assetType)
		return table, errors.NewError(msg)
	}
}
