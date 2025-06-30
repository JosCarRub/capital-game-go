package database

import "database/sql"

// InitializeSchema crea la tabla LEADERBOARD si no existe.
func InitializeSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS LEADERBOARD (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		points INT NOT NULL
	);`

	_, err := db.Exec(schema)
	return err
}
