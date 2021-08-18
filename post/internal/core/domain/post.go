package domain

import "errors"

type Post struct {
	PostId string
	Description string
}


func NewPost(PostId string, Description string) (*Post, error){
	if len(Description) < 10 {
		return nil, errors.New("description should have the length of at least 10 characters")
	}
	return &Post{PostId: PostId, Description: Description}, nil
}


func NewPostWithoutId(Description string) (*Post, error){
	if len(Description) < 10 {
		return nil, errors.New("description should have the length of at least 10 characters")
	}
	return &Post{Description: Description}, nil
}


