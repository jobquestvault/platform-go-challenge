package pg

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lib/pq"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/model"
	"github.com/jobquestvault/platform-go-challenge/internal/domain/port"
)

// Charts

// GetAllCharts retrieves all charts from the database
func (ar *AssetRepo) GetAllCharts(ctx context.Context) (charts []model.Chart, err error) {
	db, ok := ar.PgDB()
	if !ok {
		return charts, NoConnectionError
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM ak.charts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var chart model.Chart
		err := rows.Scan(&chart.ID, &chart.Title, &chart.XAxisTitle, &chart.YAxisTitle, pq.Array(&chart.Data), chart.Faved())
		if err != nil {
			return nil, err
		}
		charts = append(charts, chart)
	}

	return charts, nil
}

// GetChartByID retrieves a specific chart by its ID from the database
func (ar *AssetRepo) GetChartByID(ctx context.Context, id string) (chart model.Chart, err error) {
	db, ok := ar.PgDB()
	if !ok {
		return chart, NoConnectionError
	}

	err = db.QueryRowContext(ctx, "SELECT * FROM ak.charts WHERE id = $1", id).Scan(&chart.ID, &chart.Title, &chart.XAxisTitle, &chart.YAxisTitle, pq.Array(&chart.Data), chart.Faved())
	if err != nil {
		return chart, err
	}
	return chart, nil
}

// CreateChart creates a new chart in the database
func (ar *AssetRepo) CreateChart(ctx context.Context, chart model.Chart) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionError
	}

	_, err := db.ExecContext(ctx, "INSERT INTO ak.charts (id, title, x_axis_title, y_axis_title, data, favorite) VALUES ($1, $2, $3, $4, $5, $6)", chart.ID, chart.Title, chart.XAxisTitle, chart.YAxisTitle, pq.Array(chart.Data), chart.Faved())
	return err
}

// UpdateChart updates an existing chart in the database
func (ar *AssetRepo) UpdateChart(ctx context.Context, chart model.Chart) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionError
	}

	_, err := db.ExecContext(ctx, "UPDATE ak.charts SET title = $2, x_axis_title = $3, y_axis_title = $4, data = $5, favorite = $6 WHERE id = $1", chart.ID, chart.Title, chart.XAxisTitle, chart.YAxisTitle, pq.Array(chart.Data), chart.Faved())
	return err
}

// DeleteChart deletes a chart from the database by its ID
func (ar *AssetRepo) DeleteChart(ctx context.Context, id string) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionError
	}

	_, err := db.ExecContext(ctx, "DELETE FROM ak.charts WHERE id = $1", id)
	return err
}

// Insights

// GetAllInsights retrieves all insights from the database
func (ar *AssetRepo) GetAllInsights(ctx context.Context) (insight []model.Insight, err error) {
	db, ok := ar.PgDB()
	if !ok {
		return insight, NoConnectionError
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM ak.insights")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var insights []model.Insight
	for rows.Next() {
		var insight model.Insight
		err := rows.Scan(&insight.ID, &insight.Text, &insight.Topic, insight.Faved())
		if err != nil {
			return nil, err
		}
		insights = append(insights, insight)
	}

	return insights, nil
}

// GetInsightByID retrieves a specific insight by its ID from the database
func (ar *AssetRepo) GetInsightByID(ctx context.Context, id string) (insight model.Insight, err error) {
	db, ok := ar.PgDB()
	if !ok {
		return insight, NoConnectionError
	}

	err = db.QueryRowContext(ctx, "SELECT * FROM ak.insights WHERE id = $1", id).Scan(&insight.ID, &insight.Text, &insight.Topic, insight.Faved())
	if err != nil {
		return insight, err
	}
	return insight, nil
}

// CreateInsight creates a new insight in the database
func (ar *AssetRepo) CreateInsight(ctx context.Context, insight model.Insight) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionError
	}

	_, err := db.ExecContext(ctx, "INSERT INTO ak.insights (id, text, topic, favorite) VALUES ($1, $2, $3, $4)", insight.ID, insight.Text, insight.Topic, insight.Faved())
	return err
}

// UpdateInsight updates an existing insight in the database
func (ar *AssetRepo) UpdateInsight(ctx context.Context, insight model.Insight) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionError
	}

	_, err := db.ExecContext(ctx, "UPDATE ak.insights SET text = $2, topic = $3, favorite = $4 WHERE id = $1", insight.ID, insight.Text, insight.Topic, insight.Faved())
	return err
}

// DeleteInsight deletes an insight from the database by its ID
func (ar *AssetRepo) DeleteInsight(ctx context.Context, id string) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionError
	}

	_, err := db.ExecContext(ctx, "DELETE FROM ak.insights WHERE id = $1", id)
	return err
}

// Audiences

// GetAllAudiences retrieves all audiences from the database
func (ar *AssetRepo) GetAllAudiences(ctx context.Context) (audiences []model.Audience, err error) {
	db, ok := ar.PgDB()
	if !ok {
		return nil, NoConnectionError
	}

	rows, err := db.QueryContext(ctx, "SELECT * FROM ak.audiences")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var audience model.Audience
		err := rows.Scan(&audience.ID, &audience.Gender, &audience.BirthCountry, &audience.AgeGroup, &audience.HoursSpentOnSocial, &audience.NumPurchasesLastMth, audience.Faved())
		if err != nil {
			return nil, err
		}
		audiences = append(audiences, audience)
	}

	return audiences, nil
}

// GetAudienceByID retrieves a specific audience by its ID from the database
func (ar *AssetRepo) GetAudienceByID(ctx context.Context, id string) (audience model.Audience, err error) {
	db, ok := ar.PgDB()
	if !ok {
		return audience, NoConnectionError
	}

	err = db.QueryRowContext(ctx, "SELECT * FROM ak.audiences WHERE id = $1", id).Scan(&audience.ID, &audience.Gender, &audience.BirthCountry, &audience.AgeGroup, &audience.HoursSpentOnSocial, &audience.NumPurchasesLastMth, audience.Faved())
	if err != nil {
		return model.Audience{}, err
	}
	return audience, nil
}

// CreateAudience creates a new audience in the database
func (ar *AssetRepo) CreateAudience(ctx context.Context, audience model.Audience) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionError
	}

	_, err := db.ExecContext(ctx, "INSERT INTO ak.audiences (id, gender, birth_country, age_group, hours_spent_on_social, num_purchases_last_mth, favorite) VALUES ($1, $2, $3, $4, $5, $6, $7)", audience.ID, audience.Gender, audience.BirthCountry, audience.AgeGroup, audience.HoursSpentOnSocial, audience.NumPurchasesLastMth, audience.Faved())
	return err
}

// UpdateAudience updates an existing audience in the database
func (ar *AssetRepo) UpdateAudience(ctx context.Context, audience model.Audience) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionError
	}

	_, err := db.ExecContext(ctx, "UPDATE ak.audiences SET gender = $2, birth_country = $3, age_group = $4, hours_spent_on_social = $5, num_purchases_last_mth = $6, favorite = $7 WHERE id = $1", audience.ID, audience.Gender, audience.BirthCountry, audience.AgeGroup, audience.HoursSpentOnSocial, audience.NumPurchasesLastMth, audience.Faved())
	return err
}

// DeleteAudience deletes an audience from the database
func (ar *AssetRepo) DeleteAudience(ctx context.Context, id string) error {
	db, ok := ar.PgDB()
	if !ok {
		return NoConnectionError
	}

	_, err := db.ExecContext(ctx, "DELETE FROM ak.audiences WHERE id = $1", id)
	return err
}

// GetAssets retrieves all assets from the database
func (ar *AssetRepo) GetAssetss(ctx context.Context, status ...port.AssetStatus) (assets []model.Asset[model.Favable], err error) {
	db, ok := ar.PgDB()
	if !ok {
		return assets, NoConnectionError
	}

	query := `
		SELECT
			c.id AS chart_id, c.title AS chart_title, c.x_axis_title, c.y_axis_title, c.data,
			i.id AS insight_id, i.text AS insight_text, i.topic,
			a.id AS audience_id, a.gender, a.birth_country, a.age_group, a.hours_spent_on_social, a.num_purchases_last_mth,
			f.favorite
		FROM ak.charts c
		LEFT JOIN ak.insights i ON c.id = i.insight_id
		LEFT JOIN ak.audiences a ON c.id = a.audience_id
		LEFT JOIN favorites f ON f.entity_id = COALESCE(c.id, i.id, a.id) 
	`

	if len(status) > 0 {
		switch status[0] {
		case port.Faved:
			query = query + "WHERE f.favorite == TRUE"
		case port.NotFaved:
			query = query + "WHERE f.favorite == FALSE"
		}
	}

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result rows
	for rows.Next() {
		// Variables to hold the retrieved data
		var (
			chartID, chartName, chartTitle, xAxisTitle, yAxisTitle, insightID, insightName, insightText, insightTopic,
			audienceID, audienceName, gender, birthCountry, ageGroup string
			data                                    []uint8
			hoursSpentOnSocial, numPurchasesLastMth int
			favorite                                bool
		)

		// Scan the row values into the variables
		err := rows.Scan(
			&chartID, &chartName, &chartTitle, &xAxisTitle, &yAxisTitle, pq.Array(&data),
			&insightID, &insightName, &insightText, &insightTopic,
			&audienceID, &audienceName, &gender, &birthCountry, &ageGroup, &hoursSpentOnSocial, &numPurchasesLastMth,
			&favorite,
		)
		if err != nil {
			return nil, err
		}

		// Construct the appropriate fa struct based on the retrieved values
		var fa model.FavableAsset
		switch {
		case chartID != "":
			fa = model.Chart{
				ID:         model.ID{ID: chartID, Name: chartName},
				Title:      chartTitle,
				XAxisTitle: xAxisTitle,
				YAxisTitle: yAxisTitle,
				Data:       data,
			}
		case insightID != "":
			fa = model.Insight{
				ID:    model.ID{ID: insightID, Name: insightName},
				Text:  insightText,
				Topic: insightTopic,
			}
		case audienceID != "":
			fa = model.Audience{
				ID:                  model.ID{ID: audienceID, Name: audienceName},
				Gender:              gender,
				BirthCountry:        birthCountry,
				AgeGroup:            ageGroup,
				HoursSpentOnSocial:  hoursSpentOnSocial,
				NumPurchasesLastMth: numPurchasesLastMth,
			}
		}

		// Create the Asset object and set the favorite field
		asset := model.Asset[model.Favable]{
			ID: model.ID{
				ID:   fa.GetID(),
				Name: fa.GetName(),
			},
			Data: fa,
		}

		// Set the favorite field based on the retrieved value
		//if favorite {
		//	asset.favorite = true
		//}

		// Append the asset to the slice
		assets = append(assets, asset)
	}

	// Check for any errors during the iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Return the retrieved assets
	return assets, nil
}

// AddFav mark an asset as faved
func (ar *AssetRepo) AddFav(ctx context.Context, asset *model.Asset[model.Favable]) error {
	return nil
}

// RemoveFav mark an asset as not faved
func (ar *AssetRepo) RemoveFav(ctx context.Context, asset *model.Asset[model.Favable]) error {
	return nil
}

// UpdateFav updates asset name
func (ar *AssetRepo) UpdateFav(ctx context.Context, asset *model.Asset[model.Favable]) error {
	return nil
}

func (ar *AssetRepo) GetAssets(ctx context.Context, userID string, status ...port.AssetStatus) (assets []model.Asset[model.Favable], err error) {
	db, ok := ar.PgDB()
	if !ok {
		return assets, NoConnectionError
	}

	var favCondition string
	if len(status) > 0 {
		switch status[0] {
		case port.Faved:
			favCondition = " AND favorite = TRUE "
		case port.NotFaved:
			favCondition = " AND (favorite = FALSE OR favorite IS NULL) "
		}
	}

	query := fmt.Sprintf(`
SELECT id, name, user_id, 'chart' AS type, title, x_axis_title, y_axis_title, data, favorite, NULL, NULL, NULL, NULL, NULL, NULL, NULL
FROM ak.charts
WHERE user_id = $1%s
UNION ALL
SELECT id, name, user_id, 'insight' AS type, NULL, NULL, NULL, NULL, favorite, text, topic, NULL, NULL, NULL, NULL, NULL
FROM ak.insights
WHERE user_id = $1%s
UNION ALL
SELECT id, name, user_id, 'audience' AS type, NULL, NULL, NULL, NULL, favorite, NULL, NULL, gender, birth_country, age_group, NULL, NULL
FROM ak.audiences
WHERE user_id = $1%s`, favCondition, favCondition, favCondition)

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		return assets, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return assets, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id                  string
			name                string
			userID              string
			assetType           string
			favorite            sql.NullBool
			title               sql.NullString
			xaxisTitle          sql.NullString
			yaxisTitle          sql.NullString
			data                []uint8
			text                sql.NullString
			topic               sql.NullString
			gender              sql.NullString
			birthCountry        sql.NullString
			ageGroup            sql.NullString
			hoursSpentOnSocial  sql.NullInt64
			numPurchasesLastMth sql.NullInt64
		)

		err = rows.Scan(&id, &name, &userID, &assetType, &title, &xaxisTitle, &yaxisTitle, &data, &favorite, &text, &topic, &gender, &birthCountry, &ageGroup, &hoursSpentOnSocial, &numPurchasesLastMth)
		if err != nil {
			return assets, err
		}

		switch assetType {
		case "chart":
			chart := model.Chart{
				ID:         model.ID{ID: id, Name: name},
				Title:      title.String,
				XAxisTitle: xaxisTitle.String,
				YAxisTitle: yaxisTitle.String,
				Data:       data,
				Favorite:   model.Favorite(favorite.Bool),
			}
			assets = append(assets, model.NewAsset(id, name, assetType, chart))

		case "insight":
			insight := model.Insight{
				ID:       model.ID{ID: id, Name: name},
				Text:     text.String,
				Topic:    topic.String,
				Favorite: model.Favorite(favorite.Bool),
			}
			assets = append(assets, model.NewAsset(id, name, assetType, insight))

		case "audience":
			audience := model.Audience{
				ID:                  model.ID{ID: id, Name: name},
				Gender:              gender.String,
				BirthCountry:        birthCountry.String,
				AgeGroup:            ageGroup.String,
				HoursSpentOnSocial:  int(hoursSpentOnSocial.Int64),
				NumPurchasesLastMth: int(numPurchasesLastMth.Int64),
				Favorite:            model.Favorite(favorite.Bool),
			}
			assets = append(assets, model.NewAsset(id, name, assetType, audience))
		}
	}

	if err = rows.Err(); err != nil {
		return assets, err
	}

	return assets, nil
}
