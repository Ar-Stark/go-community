package repository

import "sync"

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

// 通过Dao对象访问数据
type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

// 单例模式
func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(func() {
		topicDao = &TopicDao{}
	})
	return topicDao
}

// 通过topic id获取topic
func (*TopicDao) QueryTopicById(id int64) *Topic {
	return topicIndexMap[id]
}
