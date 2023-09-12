package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"translation-app/service"
	"translation-app/service/entity"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB() service.TranslateRepository{
	 db, err := sql.Open("postgres", "postgresql://root:password@localhost:5433/go-translate?sslmode=disable")
	if err != nil {
		return nil
	}

	return PostgresDB{db: db}
}

func (db PostgresDB) GetTranslation(ctx context.Context, orgText, source, dest string) (entity.Translation, error){
	fmt.Println("Get translation from PostgresDB")
	var id string
	translate := entity.Translation{}

	query := " SELECT id, original_text, source, destination, result_text FROM translate WHERE original_text = $1 AND source = $2 AND destination = $3"
	err := db.db.QueryRowContext(ctx, query, orgText, source, dest).Scan(&id, &translate.OriginalText, &translate.Source, &translate.Destination, &translate.ResultText)
	if err == nil{
		return translate, nil
	}

	return entity.Translation{}, err
}
func (db PostgresDB) FindHistories(ctx context.Context) ([]entity.Translation, error){
	fmt.Println("Get histories from PostgresDB")
	rows, err := db.db.Query("SELECT original_text, source, destination, result_text FROM translate")
	if err != nil{
		return []entity.Translation{}, err
	}
	defer rows.Close()

	var translations []entity.Translation
	for rows.Next() {
		var trans entity.Translation
		if err := rows.Scan(&trans.OriginalText, &trans.Source, &trans.Destination, &trans.ResultText); err != nil{
			return translations, err
		}
		translations = append(translations, trans)
	}
	if err = rows.Err(); err != nil {
		return translations, err
	}

	return translations, nil
}
func (db PostgresDB) InsertTranslation(ctx context.Context, translation entity.Translation) (error) {
	fmt.Println("Insert translation from PostgresDB")
	var lastInsertId int
	query := "INSERT INTO translate(original_text, source, destination, result_text) VALUES ($1, $2, $3, $4) returning id"
	err := db.db.QueryRowContext(ctx, query, translation.OriginalText, translation.Source, translation.Destination, translation.ResultText).Scan(&lastInsertId)
	if err != nil {
		return err
	}

	return nil
}