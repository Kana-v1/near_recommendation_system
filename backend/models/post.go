package models

import "time"

type Post struct {
	Id             uint64
	CreatorID      AccountID
	BlockTimestamp time.Time
	Content        []byte
	Comments       []*Comment
	Likes          []*Like
}

type Comment struct {
	AccID       AccountID
	Text        string
	BlockHeight uint64
	Likes       uint32
}

type Like struct {
	AccID       AccountID
	BlockHeight uint64
}


var postsCounter = uint64(0)

func NewPost(accID AccountID, content []byte) *Post {
	postsCounter ++
	return &Post{
		Id: postsCounter ,
		CreatorID: accID,
		BlockTimestamp: time.Now(),
		Content: content,
		Comments: make([]*Comment, 0),
		Likes: make([]*Like, 0),
	}
}

