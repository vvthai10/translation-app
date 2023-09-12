package googlesv

import (
	"context"
	"fmt"
	"translation-app/service"
	"translation-app/service/entity"

	translator "github.com/Conight/go-googletrans"
)

type googleTraslateAPI struct {
}

func New() service.OnlineService{
	return googleTraslateAPI{}
}

func (api googleTraslateAPI) Translate(ctx context.Context, orgText, source, dest string) (entity.Translation, error){
	fmt.Println("Use service of google")
	conf := translator.Config{
		UserAgent:   []string{"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1"},
		ServiceUrls: []string{"translate.google.com"},
	}

	trans := translator.New(conf)

	result, err := trans.Translate(orgText, source, dest)
	if err != nil{
		return entity.Translation{}, err
	}
	return entity.NewTranslation(orgText, source, dest, result.Text), nil
}