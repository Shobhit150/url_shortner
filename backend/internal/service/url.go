package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	// "fmt"

	"github.com/Shobhit150/url_shortner/internal/cache"
	"github.com/Shobhit150/url_shortner/internal/kafka"
	"github.com/Shobhit150/url_shortner/internal/repository"
	"github.com/Shobhit150/url_shortner/internal/utils"
)





func Shorten(longURL string, CustomSlug string, expiresAt *time.Time) (string, error) {
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
	if expiresAt == nil {
		t := time.Now().Add(30 * 24 * time.Hour)
		expiresAt = &t
	}
	
	err := repository.Save(slug,longURL, expiresAt)
	if err != nil {
		return "", err
	}
	return slug, nil
}

func Redirect(slug, ip, userAgent, referrer string) (string, *time.Time, error){
	ctx := context.Background()

	
	
	longURL, err := cache.GetSlug(ctx, slug)

	if err == nil {
		err = kafka.PublishLinkClick(slug, ip, userAgent, referrer)
		if err != nil {
			fmt.Println("Error publishing to Kafka:", err)
		} else {
			fmt.Println("Published click to Kafka for slug:", slug)
		}

		return longURL,nil, nil
	}

	// fmt.Println("this is slug in after redirect", slug)
	longURL,  expiresAt, err := repository.Find(slug)

	

	if err != nil {
		// fmt.Println("error is here")
		return "",expiresAt,  err
	}
	if expiresAt != nil && time.Now().After(*expiresAt) {
		return "",expiresAt, errors.New("URL expired")
	}

	_ = cache.SetSlug(ctx, slug, longURL, expiresAt)

	err = kafka.PublishLinkClick(slug, ip, userAgent, referrer)
	if err != nil {
		fmt.Println("Error publishing to Kafka:", err)
	} else {
		fmt.Println("Published click to Kafka for slug:", slug)
	}
	return longURL, expiresAt, nil
}
func GetClicks(slug string) (int, *time.Time, error){
	click, time, err := repository.GetClickCount(slug)
	return click, time, err
}