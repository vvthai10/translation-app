package service

import (
	"context"
	"translation-app/service/entity"
)

type service struct {
	repository    TranslateRepository
	googleService OnlineService
}

func NewService(repository TranslateRepository, googleService OnlineService) service {
	return service{
		repository:    repository,
		googleService: googleService,
	}
}

func (s service) Translate(ctx context.Context, orgText, source, dest string) (entity.Translation, error){
	oldTrans, err := s.repository.GetTranslation(ctx, orgText, source, dest)

	if err == nil{
		return oldTrans, nil
	}

	newTrans, err := s.googleService.Translate(ctx, orgText, source, dest)

	if err!= nil{
		return entity.Translation{}, nil
	}

	go func(){
		_ = s.repository.InsertTranslation(ctx, newTrans)
	}()

	return newTrans, nil
}

func (s service) FetchHistories(ctx context.Context) ([]entity.Translation, error){
	return s.repository.FindHistories(ctx)
}