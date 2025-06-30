package database

import (
	"capital-game-go/internal/game"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection() (*sql.DB, error) {
	var db *sql.DB
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)

	for i := 0; i < 5; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			fmt.Println("Error al abrir la conexión, reintentando en 5 segundos...")
			time.Sleep(5 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {

			fmt.Println("¡Conexión exitosa a la base de datos!")
			return db, nil
		}

		fmt.Println("No se pudo hacer ping a la base de datos, reintentando en 5 segundos...")
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("no se pudo conectar a la base de datos después de varios intentos: %w", err)
}
func InitializeSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS LEADERBOARD (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		points INT NOT NULL
	);`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("error al inicializar el esquema: %w", err)
	}

	fmt.Println("Esquema de la base de datos inicializado")
	return nil
}

func SaveScore(db *sql.DB, name string, points int) error {

	query := "INSERT INTO LEADERBOARD (name, points) VALUES (?, ?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		return fmt.Errorf("error al preparar la consulta de inserción: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(name, points)
	if err != nil {
		return fmt.Errorf("error al ejecutar la inserción en la base de datos: %w", err)
	}

	return nil
}

func GetLeaderboard(db *sql.DB) ([]game.PlayerScore, error) {

	query := "SELECT id, name, points FROM LEADERBOARD ORDER BY points DESC"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al consultar el leaderboard: %w", err)
	}
	defer rows.Close()

	var scores []game.PlayerScore

	for rows.Next() {
		var score game.PlayerScore

		if err := rows.Scan(&score.ID, &score.Name, &score.Points); err != nil {
			return nil, fmt.Errorf("error al escanear la fila del leaderboard: %w", err)
		}

		scores = append(scores, score)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error durante la iteración de filas del leaderboard: %w", err)
	}

	return scores, nil
}
