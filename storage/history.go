package storage

import (
	"database/sql"
	"time"
)

type Row struct {
	Date           time.Time
	SourceCurrency string
	TargetCurrency string
	Amount         float64
	Result         float64
}

func List(user_id int64) ([]Row, error) {
	q := `SELECT date, source_currency, target_currency, amount, result FROM history WHERE user_id = $1`

	table, err := db.Query(q, user_id)
	if err != nil {
		return nil, err
	}
	res, err := parseTable(table)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func parseTable(rows *sql.Rows) ([]Row, error) {
	defer rows.Close()
	var table []Row
	for rows.Next() {
		var row Row
		err := rows.Scan(&row.Date, &row.SourceCurrency, &row.TargetCurrency, &row.Amount, &row.Result)
		if err != nil {
			return nil, err
		}
		table = append(table, row)
	}
	return table, nil
}

func Insert(user_id int64, source, target string, amount, result float64) error {
	q := `INSERT INTO history (user_id, source_currency, target_currency, amount, result) VALUES ($1, $2, $3, $4, $5) RETURNING date`
	_, err := db.Exec(q, user_id, source, target, amount, result)
	if err != nil {
		return err
	}
	return nil
}
