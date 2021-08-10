package post

import (
	"github.com/pooriaPoorsarvi/event-driven-design-golang/post/internal/core/domain"
	"github.com/pooriaPoorsarvi/event-driven-design-golang/post/internal/core/service"
	"github.com/pooriaPoorsarvi/event-driven-design-golang/post/internal/repositories"
)

var repo = &repositories.MemoryRepository{Posts: []domain.Post{}}
var PostService = service.NewPostService(repo)














