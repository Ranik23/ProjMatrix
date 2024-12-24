package repository

import (
	"context"
	"fmt"
	"log"
)

func (r *PgRepository) Save(ctx context.Context, id string, time float64, data []byte) error {
	log.Println(id)
	log.Println("Saving data to the database")
	query := `INSERT INTO pj.sessions (key, data, execution_time) VALUES ($1, $2, $3)`
	_, err := r.pool.Exec(ctx, query, id, data, time)
	if err != nil {
		fmt.Errorf("error in saving data: %w", err)
	}
	return nil
}
