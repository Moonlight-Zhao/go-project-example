package repository

import (
	"sync"
	"time"
)

type Topic struct {
	Id         int64     `gorm:"column:id"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}
type TopicDao struct {
}
var (
	topicDao *TopicDao
	topicOnce sync.Once
)
func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
