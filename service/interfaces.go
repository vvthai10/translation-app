package service

import (
	"context"
	"translation-app/service/entity"
)

// Tương tác với Controler
type TranslationUseCase interface {
	Translate(ctx context.Context, orgText, source, dest string) (entity.Translation, error)
	FetchHistories(ctx context.Context) ([]entity.Translation, error)
}

// Tương tác với DB/...
type TranslateRepository interface{
	GetTranslation(ctx context.Context, orgText, source, dest string) (entity.Translation, error)
	FindHistories(ctx context.Context) ([]entity.Translation, error)
	InsertTranslation(ctx context.Context, translation entity.Translation) (error) 
}

// 
type OnlineService interface{
	Translate(ctx context.Context, orgText, source, dest string) (entity.Translation, error)
}
