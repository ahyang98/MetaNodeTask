package task3

import (
	"context"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func initBlockDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:ahyang@tcp(192.168.252.128:30306)/blog"))
	if err != nil {
		panic(err)
	}
	return db
}

func initTables(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

type User struct {
	Id        int64          `gorm:"primaryKey, autoIncrement"`
	Email     sql.NullString `gorm:"unique"`
	Password  sql.NullString
	Ctime     int64
	Utime     int64
	Phone     sql.NullString `gorm:"unique"`
	PostCount int64          `bson:"post_count, omitempty"`
}

type Post struct {
	Id            int64  `gorm:"primaryKey, autoIncrement" bson:"id,omitempty"`
	Title         string `gorm:"type=varchar(4096)" bson:"title,omitempty"`
	Content       string `gorm:"type=BLOB" bson:"content,omitempty"`
	AuthorId      int64  `gorm:"index" bson:"author_id,omitempty"`
	Status        uint8  `bson:"status,omitempty"`
	Ctime         int64  `bson:"ctime,omitempty"`
	Utime         int64  `bson:"utime,omitempty"`
	CommentStatus string `bson:"comment_status"`
}

func (p *Post) BeforeCreate(db *gorm.DB) error {
	return db.Model(&User{}).Where("id = ?", p.AuthorId).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1)).Error
}

type Comment struct {
	Id       int64  `gorm:"primaryKey, autoIncrement" bson:"id,omitempty"`
	Content  string `gorm:"type=BLOB" bson:"content,omitempty"`
	PostId   int64  `gorm:"index" bson:"post_id,omitempty"`
	AuthorId int64  `gorm:"index" bson:"author_id,omitempty"`
	Status   uint8  `bson:"status,omitempty"`
	Ctime    int64  `bson:"ctime,omitempty"`
	Utime    int64  `bson:"utime,omitempty"`
}

const PostNoComment = "无评论"

func (c *Comment) AfterDelete(db *gorm.DB) error {
	var count int64
	db.Model(Comment{}).Where("post_id = ?", c.PostId).Count(&count)
	if count == 0 {
		return db.Model(Post{}).Where("id = ?", c.PostId).Update("comment_status", PostNoComment).Error
	}
	return nil
}

func TestGorm() {
	db := initBlockDb()
	err := initTables(db)
	if err != nil {
		panic(err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	user := &User{
		Id: 1,
	}
	posts, comments, err := queryPostAndComment(ctx, db, user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("user %d posts:%+v\ncomments%+v\n", user.Id, posts, comments)

	err, hotPost := queryHotPost(ctx, db)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("hot post:%+v\n", hotPost)
	post := &Post{
		Title:    "web5",
		Content:  "web5",
		AuthorId: user.Id,
	}

	db.First(&user, user.Id)
	fmt.Printf("user %+v\n", user)
	db.Create(post)
	db.First(&user, user.Id)
	fmt.Printf("user %+v\n", user)
	comment := &Comment{
		Content:  "web5 comment",
		PostId:   post.Id,
		AuthorId: 1,
	}
	db.Create(comment)
	db.First(post, post.Id)
	fmt.Printf("post %+v\n", post)
	db.Delete(comment)
	db.First(post, post.Id)
	fmt.Printf("post %+v\n", post)
}

func queryPostAndComment(ctx context.Context, db *gorm.DB, user *User) ([]Post, []Comment, error) {
	var (
		posts    []Post
		comments []Comment
	)

	postsQuery := db.WithContext(ctx).Model(&Post{}).Where("author_id = ?", user.Id)

	err := postsQuery.Find(&posts).Error
	if err != nil {
		return posts, comments, err
	}

	postIds := postsQuery.Select("id")

	err = db.WithContext(ctx).Model(&Comment{}).Where("post_id in (?)", postIds).Find(&comments).Error
	return posts, comments, err
}

func queryHotPost(ctx context.Context, db *gorm.DB) (error, Post) {
	var post Post
	err := db.WithContext(ctx).Model(&Post{}).
		Select("posts.*, count(comments.id) as comment_count").
		Joins("left join comments on posts.id=comments.post_id").
		Group("posts.id").
		Order("comment_count desc").
		First(&post).Error
	return err, post
}
