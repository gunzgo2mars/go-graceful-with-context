package service

import (
	"context"
	"fmt"
	"time"

	articleRepo "github.com/gunzgo2mars/go-graceful-with-context/src/repository/article"
)

type IService interface {
	CreateNewArticleInfo(ctx context.Context) error
}

type service struct {
	articleRepo articleRepo.IArticleRepository
}

func New(articleRepo articleRepo.IArticleRepository) IService {
	return &service{
		articleRepo: articleRepo,
	}
}

func (s *service) CreateNewArticleInfo(ctx context.Context) error {

	time.Sleep(time.Second * 1)
	fmt.Println("Work Done At layer service")

	if err := s.articleRepo.SetArticleInfo(ctx); err != nil {
		return err
	}
	return nil
}
