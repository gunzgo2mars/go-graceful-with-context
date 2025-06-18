package article

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type IArticleRepository interface {
	SetArticleInfo(ctx context.Context) error
}

type articleRepository struct {
	cacheClient *redis.Client
}

func New(cacheClient *redis.Client) IArticleRepository {
	return &articleRepository{
		cacheClient: cacheClient,
	}
}

func (r *articleRepository) SetArticleInfo(ctx context.Context) error {
	time.Sleep(time.Second * 4)
	fmt.Println("Caching article info")
	fmt.Println("Work Done at layer repository")
	return nil
}
