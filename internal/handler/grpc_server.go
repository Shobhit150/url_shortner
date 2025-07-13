package handler

import (
	"context"

	"github.com/Shobhit150/url_shortner/internal/service"
	urlshortenerpb "github.com/Shobhit150/url_shortner/proto"
)

// Implements urlshortenerpb.URLShortenerServer
type URLShortenerServer struct {
	urlshortenerpb.UnimplementedURLShortenerServer
}

func (s *URLShortenerServer) Shorten(ctx context.Context, req *urlshortenerpb.ShortenRequest) (*urlshortenerpb.ShortenResponse, error) {
	slug, err := service.Shorten(req.LongUrl, req.CustomSlug)
	if err != nil {
		return nil, err
	}
	return &urlshortenerpb.ShortenResponse{Slug: slug}, nil
}

func (s *URLShortenerServer) Redirect(ctx context.Context, req *urlshortenerpb.RedirectRequest) (*urlshortenerpb.RedirectResponse, error) {
	longURL, err := service.Redirect(req.Slug)
	if err != nil {
		return nil, err
	}
	return &urlshortenerpb.RedirectResponse{LongUrl: longURL}, nil
}

func (s *URLShortenerServer) GetStats(ctx context.Context, req *urlshortenerpb.StatsRequest) (*urlshortenerpb.StatsResponse, error) {
	clicks, err := service.GetClicks(req.Slug)
	if err != nil {
		return nil, err
	}
	return &urlshortenerpb.StatsResponse{Clicks: int32(clicks)}, nil
}
