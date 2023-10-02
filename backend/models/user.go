package models

import "fmt"

type AccountID uint64

var usersCounter = uint64(0)

type User struct {
	AccID        AccountID
	Account_desc string
	Tags         []string
	Following    []AccountID
	Followers    []AccountID
}

func NewUser(accountDescription string) *User {
	usersCounter++
	return &User{
		AccID:        AccountID(usersCounter),
		Account_desc: accountDescription,
		Tags:         make([]string, 0),
		Following:    make([]AccountID, 0),
		Followers:    make([]AccountID, 0),
	}
}

func (u *User) Follow(accountToFollow AccountID) error {
	index := -1
	usersPool := GetUsers()
	for i := range usersPool {
		if accountToFollow == usersPool[i].AccID {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("cannot find user with account id '%v'", accountToFollow)
	}

	u.Following = append(u.Following, accountToFollow)
	usersPool[index].Followers = append(usersPool[index].Followers, u.AccID)

	return nil
}

func (u *User) LikePost(postID uint64) error {
	index := -1
	postsPool := GetPosts()

	for i := range postsPool {
		if postsPool[i].Id == postID {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("post with id '%v' does not exist", postID)
	}

	like := &Like{
		AccID:       u.AccID,
		BlockHeight: blockHeight,
	}

	postsPool[index].Likes = append(postsPool[index].Likes, like)

	return nil
}

func (u *User) LikeComment(postID, commentOwnerID uint64) error {
	pIndex := -1
	postsPool := GetPosts()

	for i := range postsPool {
		if postsPool[i].Id == postID {
			pIndex = i
			break
		}
	}

	if pIndex == -1 {
		return fmt.Errorf("post with id '%v' does not exist", postID)
	}

	cIndex := -1
	for i, comment := range postsPool[pIndex].Comments {
		if comment.AccID == AccountID(commentOwnerID) {
			cIndex = i
			break
		}
	}

	if cIndex == -1 {
		return fmt.Errorf("comment from the user '%v' does not exist under the post with id '%v'", commentOwnerID, postID)
	}

	comment := postsPool[pIndex].Comments[cIndex]
	comment.Likes++

	return nil
}

func (u *User) WritePost(content []byte) error {
	post := NewPost(u.AccID, content)

	Posts = append(GetPosts(), post)

	return nil
}

func (u *User) WriteComment(postID AccountID, comment string) error {
	com := &Comment{
		AccID:       u.AccID,
		Text:        comment,
		BlockHeight: blockHeight,
		Likes:       0,
	}

	Comments = append(Comments, com)

	return nil
}

func (u *User) GetRecommendedPosts() ([]*Post, error) {
	return Posts, nil
}
