package handler

import (
	"context"
	"time"

	"github.com/Shobhit150/url_shortner/internal/service"
	urlshortenerpb "github.com/Shobhit150/url_shortner/proto"
)

// Implements urlshortenerpb.URLShortenerServer
type URLShortenerServer struct {
	urlshortenerpb.UnimplementedURLShortenerServer
}

func (s *URLShortenerServer) Shorten(ctx context.Context, req *urlshortenerpb.ShortenRequest) (*urlshortenerpb.ShortenResponse, error) {
	// Parse expires_at if provided
	var expiresAt *time.Time
	if req.ExpiresAt != "" {
		parsed, err := time.Parse(time.RFC3339, req.ExpiresAt)
		if err != nil {
			return nil, err
		}
		expiresAt = &parsed
	}

	slug, err := service.Shorten(req.LongUrl, req.CustomSlug, expiresAt)
	if err != nil {
		return nil, err
	}

	resp := &urlshortenerpb.ShortenResponse{
		Slug:      slug,
		// ExpiresAt: req.ExpiresAt, // (Optional: can set actual DB value if returned by service)
	}
	return resp, nil
}

func (s *URLShortenerServer) Redirect(ctx context.Context, req *urlshortenerpb.RedirectRequest) (*urlshortenerpb.RedirectResponse, error) {
	// Pass all analytics fields to the service
	longURL, expiresAt, err := service.Redirect(
		req.Slug,
		req.IpAddress,
		req.UserAgent,
		req.Referrer,
	)
	if err != nil {
		return nil, err
	}
	resp := &urlshortenerpb.RedirectResponse{
		LongUrl:   longURL,
		ExpiresAt: "",
	}
	if expiresAt != nil {
		resp.ExpiresAt = expiresAt.Format(time.RFC3339)
	}
	return resp, nil
}

func (s *URLShortenerServer) GetStats(ctx context.Context, req *urlshortenerpb.StatsRequest) (*urlshortenerpb.StatsResponse, error) {
	clicks, expiresAt, err := service.GetClicks(req.Slug)
	if err != nil {
		return nil, err
	}
	resp := &urlshortenerpb.StatsResponse{
		Clicks:    int32(clicks),
		ExpiresAt: "",
	}
	if expiresAt != nil {
		resp.ExpiresAt = expiresAt.Format(time.RFC3339)
	}
	return resp, nil
}
