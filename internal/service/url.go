package service

import (
	"errors"
	"github.com/Shobhit150/url_shortner/internal/repository"
	"github.com/Shobhit150/url_shortner/internal/utils"
)
func Shorten(longURL string) (string, error) {
	slug := utils.GenerateSlug()

	err := repository.Save(slug, longURL) 

	if err != nil {
		return "", err
	}
	return slug, nil
}

func Resolve(slug string) (string, error){
	return repository.Find(slug)
}