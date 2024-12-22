package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

func (r *PgRepository) Get(ctx context.Context, id string) ([]byte, float64, error) {
	query := `SELECT data, execution_time FROM pj.sessions WHERE key = $1`

	row := r.pool.QueryRow(ctx, query, id)

	var (
		data []byte
		dur  float64
	)

	if err := row.Scan(&data, &dur); err != nil {
		if err == pgx.ErrNoRows {
			log.Println("No data found for id:", id)
			return nil, 0, fmt.Errorf("no data found for id: %s", id)
		}
		log.Println("Error scanning row:", err)
		return nil, 0, fmt.Errorf("error scanning row: %w", err)
	}

	return data, dur, nil
}
