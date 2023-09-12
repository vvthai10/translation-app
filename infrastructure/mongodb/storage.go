package mongodb

import (
	"context"
	"database/sql"
	"fmt"
	"translation-app/service"
	"translation-app/service/entity"
)

type mongodbRepo struct {
	db *sql.DB
}

func NewMongoDB(db *sql.DB) service.TranslateRepository{
	fmt.Println("Create mongoDB")
	return mongodbRepo{db: db}
}

func (db mongodbRepo) GetTranslation(ctx context.Context, orgText, source, dest string) (entity.Translation, error){
	fmt.Println("Get translation from mongoDB")
	return entity.Translation{}, nil
}
func (db mongodbRepo) FindHistories(ctx context.Context) ([]entity.Translation, error){
	fmt.Println("Get histories from mongoDB")
	return []entity.Translation{}, nil
}
func (db mongodbRepo) InsertTranslation(ctx context.Context, translation entity.Translation) (error) {
	fmt.Println("Insert translation from mongoDB")
	return nil
}
