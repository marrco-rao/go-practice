package main

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey;column:id"`
	Username  string `gorm:"column:username;unique;size:64"`
	Age       int    `gorm:"column:age"`
	Email     string `gorm:"column:email;size:255"`
	PostCount int    `gorm:"column:post_count"`
	Posts     []Post `gorm:"foreignKey:userID"`
}

func (User) TableName() string {
	return "users"
}

type Post struct {
	ID           uint      `gorm:"primaryKey;column:id"`
	UserID       uint      `gorm:"column:user_id;size:64"`
	Title        string    `gorm:"column:title;size:128"`
	Content      string    `gorm:"column:content"`
	CommentState int       `gorm:"column:comment_state"`
	Comments     []Comment `gorm:"foreignKey:postID"`
}

func (Post) TableName() string {
	return "posts"
}

// post 勾子
func (p *Post) AfterCreate(tx *gorm.DB) error {
	fmt.Println("AfterCreate......")
	err := updatePostCount(tx, p.UserID)
	if err != nil {
		return err
	}
	return nil
}
func updatePostCount(tx *gorm.DB, userId uint) error {
	fmt.Println("updatePostCount start.")
	var postCount int64
	rs := tx.Table("posts").Where("user_id = ?", userId).Count(&postCount)
	if rs.Error != nil {
		fmt.Println("post count err:", rs.Error)
		return rs.Error
	}
	err := tx.Table("users").Where("id = ?", userId).Update("post_count", postCount).Error
	if err != nil {
		return err
	}
	fmt.Println("updatePostCount finished....")
	return nil
}

type Comment struct {
	ID             uint   `gorm:"primaryKey;column:id"`
	PostId         uint   `gorm:"column:post_id"`
	CommentContent string `gorm:"column:comment_content;size:512"`
}

func (Comment) TableName() string {
	return "comments"
}

func (c *Comment) AfterCreate(tx *gorm.DB) error {
	fmt.Println("AfterCreate Comment......")
	// comment_state : 1-有评论，0-无评论
	err := tx.Table("posts").Where("id", c.PostId).Update("comment_state", 1).Error
	if err != nil {
		fmt.Println("update comment_state error. ", err)
		return err
	}
	return nil
}
func (c *Comment) AfterDelete(tx *gorm.DB) error {
	fmt.Println("AfterDeleteComment......")
	err := updateCommentStatus(tx, c.PostId)
	if err != nil {
		return err
	}
	return nil
}
func updateCommentStatus(tx *gorm.DB, postId uint) error {
	var commentCount int64
	rs := tx.Table("comments").Where("post_id", postId).Count(&commentCount)
	if rs.Error != nil {
		fmt.Println("comments count err:", rs.Error)
		return rs.Error
	}
	var e error
	if commentCount == 0 {
		e = tx.Table("posts").Where("id", postId).Update("comment_state", 0).Error
	} else {
		e = tx.Table("posts").Where("id", postId).Update("comment_state", 1).Error
	}
	if e != nil {
		return e
	}
	return nil
}
