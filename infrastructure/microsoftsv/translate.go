package microsoftsv

import (
	"context"
	"fmt"
	"translation-app/service"
	"translation-app/service/entity"
)

type microsoftTranslateAPI struct {
}

func New() service.OnlineService{
	return microsoftTranslateAPI{}
}

func (ms microsoftTranslateAPI) Translate(ctx context.Context, orgText, source, dest string) (entity.Translation, error){
	fmt.Println("Use service of microsoft")

	return entity.Translation{}, nil
}