package article

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gunzgo2mars/go-graceful-with-context/src/model"
	"github.com/redis/go-redis/v9"
)

const (
	CREATE_ARTICLE_KEY = "ARTICLE|%s"
)

type IArticleRepository interface {
	SetArticleInfo(ctx context.Context, schema *model.CacheArticleInfoSchema) error
}

type articleRepository struct {
	cacheClient *redis.Client
}

func New(cacheClient *redis.Client) IArticleRepository {
	return &articleRepository{
		cacheClient: cacheClient,
	}
}

func (r *articleRepository) SetArticleInfo(ctx context.Context, schema *model.CacheArticleInfoSchema) error {
	time.Sleep(time.Second * 3)
	fmt.Println("Caching article info")
	fmt.Println("Work Done at layer repository")

	if err := r.cacheClient.HMSet(ctx, fmt.Sprintf(CREATE_ARTICLE_KEY, strconv.Itoa(schema.ID)), schema).Err(); err != nil {
		return err
	}

	return nil
}
