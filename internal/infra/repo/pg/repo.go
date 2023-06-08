package pg

import (
	"context"

	"github.com/lib/pq"

	"github.com/jobquestvault/platform-go-challenge/internal/domain/model"
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
