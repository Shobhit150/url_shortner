package service

import (
	"context"
	"errors"
	// "fmt"

	"github.com/Shobhit150/url_shortner/internal/cache"
	"github.com/Shobhit150/url_shortner/internal/repository"
	"github.com/Shobhit150/url_shortner/internal/utils"
)
func Shorten(longURL string, CustomSlug string) (string, error) {
	var slug string = ""
	if(CustomSlug == "") {
		for {
			slug = utils.GenerateSlug()
			exists, err := repository.Exists(slug)
			if(err!=nil) {
				return "", err;
			}
			if(!exists) {
				break
			}
		}
		// slug := utils.GenerateSlug()

		// exists, _ := repository.Exists(slug)

		// if exists {
		// 	return "", errors.New("slug collision, try again")
		// }

		// err := repository.Save(slug, longURL) 

		// if err != nil {
		// 	return "", err
		// }
		// return slug, nil
	}else {
		exists, err := repository.Exists(CustomSlug)
		if(err != nil) {
			return "", err
		}
		if(exists) {
			return "", errors.New("custom slug is already is already taken")
		}
		slug = CustomSlug
	}
	err := repository.Save(slug,longURL)
	if err != nil {
		return "", err
	}
	return slug, nil

}

func Redirect(slug string) (string, error){
	ctx := context.Background()

	// fmt.Println("this is slug in before redirect", slug)
	longURL, err := cache.GetSlug(ctx, slug)

	if err == nil {
		return longURL, nil
	}

	// fmt.Println("this is slug in after redirect", slug)
	longURL, err = repository.Find(slug)

	if err != nil {
		// fmt.Println("error is here")
		return "", err
	}

	_ = cache.SetSlug(ctx, slug, longURL)
	return longURL, nil
}
func GetClicks(slug string) (int, error){
	return repository.GetClickCount(slug)
}