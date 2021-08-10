package domain

import "errors"

type Post struct {
	Description string
}


func NewPost(Description string) (*Post, error){
	if len(Description) < 10 {
		return nil, errors.New("description should have the length of at least 10 characters")
	}
	return &Post{Description: Description}, nil
}




