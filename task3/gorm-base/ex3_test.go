package main

import (
	"fmt"
	"testing"
)

// 日志暂时用fmt

func TestFindByUserId(t *testing.T) {
	InitDb()
	u, e := FindByUserId(1)
	if e != nil {
		fmt.Printf("FindByUserId error: %v", e)
	}
	if u != nil {
		if u.Posts != nil {
			for _, v := range u.Posts {
				fmt.Println("p:", v)
				if v.Comments != nil {
					for _, c := range v.Comments {
						fmt.Println("c:", c)
					}
				}
			}
		}
	}
}

func TestFindPostHasTopComments(t *testing.T) {
	InitDb()
	var post Post
	subQuery := db.Table("comments").Select("post_id").Group("post_id").Order("count(*) desc").Limit(1)
	// 这里 posts 别名不单独指定，直接用posts,避免跟gorm自身的语句转换不匹配
	err := db.Table("posts").
		Joins("join(?) t on posts.id=t.post_id", subQuery).First(&post).Error
	if err != nil {
		t.Errorf("FindPostHasTopComments error: %v", err)
	}
	t.Log(post)
}

func TestCreatePost(t *testing.T) {
	InitDb()
	var post Post = Post{
		UserID:  3,
		Title:   "文章添加测试3",
		Content: "文章添加测试3,内容......",
	}
	var uBase User
	db.Find(&uBase, post.UserID)
	t.Logf("user base. %v", uBase)
	e := db.Create(&post).Error
	if e != nil {
		fmt.Printf("createPost error: %v", e)
	}
	var u User
	db.Find(&u, post.UserID)
	t.Logf("user after count. %v", u)
	if u.PostCount != uBase.PostCount+1 {
		t.Errorf("post count error.")
	}
	t.Logf("add post info: %v", post)
}

func TestDeleteComment(t *testing.T) {
	InitDb()
	var comment Comment = Comment{
		PostId:         6,
		CommentContent: "测试评论1:文章添加测试1",
	}
	db.Create(&comment)

	// test delete
	t.Logf("start test delete comment")
	db.Delete(&comment)
	var modifyPost Post
	db.Find(&modifyPost, comment.PostId)
	if modifyPost.CommentState != 0 {
		t.Errorf("评论状态应该为0（无评论）,	实际为%d", modifyPost.CommentState)
	}
}
