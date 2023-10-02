package models

var (
	Users = make([]*User, 0)
	Posts = make([]*Post, 0)
	Comments = make([]*Comment, 0)

	blockHeight = uint64(152)
)

func GetUsers() []*User {
	return Users
}

func GetPosts() []*Post {
	return Posts
}