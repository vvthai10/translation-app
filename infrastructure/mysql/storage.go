package mysql

import (
	"context"
	"fmt"
	"translation-app/service"
	"translation-app/service/entity"

	"gorm.io/gorm"
)

type mysqlRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) service.TranslateRepository {
	return mysqlRepo{db: db}
}

func (repo mysqlRepo) InsertTranslation(ctx context.Context, translation entity.Translation) error {
	fmt.Println("Insert translation from MySQL")
	return repo.db.Create(&translation).Error
}

func (repo mysqlRepo) GetTranslation(ctx context.Context, orgText, source, dest string) (entity.Translation, error) {
	fmt.Println("Get translation from MySQL")
	return entity.Translation{}, nil
}

func (repo mysqlRepo) FindHistories(ctx context.Context) ([]entity.Translation, error) {
	fmt.Println("Get histories from MySQL")
	return []entity.Translation{}, nil
}