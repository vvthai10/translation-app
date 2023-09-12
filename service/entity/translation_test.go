package entity_test

import (
	"testing"
	"translation-app/service/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewTranslation(t *testing.T){
	trans := entity.NewTranslation("Hello", "en", "vi", "Xin chào")
	assert.Equal(t, trans.OriginalText, "Hello")
}

func TestBookValidate(t *testing.T){
	type test struct{
		OriginalText string
		Source string
		Destination string 
		ResultText string 
		want error
	}

	tests := []test{
		{
			OriginalText: "Hello",
			Source: "en",
			Destination: "vi",
			ResultText: "Xin chào",
			want: nil,
		},
	}

	for _, tc := range tests{
		trans := entity.NewTranslation(tc.OriginalText, tc.Source, tc.Destination, tc.ResultText)
		assert.Equal(t, trans.OriginalText, tc.OriginalText)
	}
}