package service

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/gunzgo2mars/go-graceful-with-context/src/model"
	articleRepo "github.com/gunzgo2mars/go-graceful-with-context/src/repository/article"
)

type IService interface {
	CreateNewArticleInfo(ctx context.Context, request *model.RequestCreateArticleInfo) error
}

type service struct {
	articleRepo articleRepo.IArticleRepository
}

func New(articleRepo articleRepo.IArticleRepository) IService {
	return &service{
		articleRepo: articleRepo,
	}
}

func (s *service) CreateNewArticleInfo(ctx context.Context, request *model.RequestCreateArticleInfo) error {

	// Delay for testing graceful
	time.Sleep(time.Second * 3)
	fmt.Println("Work Done At layer service")

	if err := s.articleRepo.SetArticleInfo(ctx, mapRequestCreateArticleToSchema(request)); err != nil {
		return err
	}
	return nil
}

func mapRequestCreateArticleToSchema(request *model.RequestCreateArticleInfo) *model.CacheArticleInfoSchema {

	randomId := rand.IntN(100)
	fmt.Printf("Random ID: %d \n", randomId)
	return &model.CacheArticleInfoSchema{
		ID:     randomId,
		Author: request.Author,
		Title:  request.Text,
		Text:   request.Text,
	}
}
