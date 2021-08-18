package repositories

import "github.com/pooriaAcademy/event-driven-design-golang/post/internal/core/domain"

type MemoryRepository struct {
	Posts []domain.Post
}


func (this * MemoryRepository) SavePost(post * domain.Post) error {
	this.Posts = append(this.Posts, *post)
	var a MemoryRepository
	a = MemoryRepository{Posts: []domain.Post{}}
	print(a)
	return nil
}

func (this * MemoryRepository) NewPostId(post * domain.Post) (string, error) {
	return "new id", nil
}





