package db_repository

func (pg *PgRepository) GetIsDatabaseUpAndRunning() (bool, error) {
	// Query to check if the database is up and running
	query := `SELECT 1`

	// Executing the query
	var result int
	err := pg.db.Raw(query).Row().Scan(&result)
	if err != nil {
		return false, err
	}

	return true, nil
}
