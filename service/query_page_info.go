package service

import (
	"errors"
	"go-community/repository"
	"sync"
)

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic    *repository.Topic
	postList []*repository.Post
}

func NewQueryPageInfoFlow(topicId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topicId,
	}
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	err := f.checkParam()
	if err != nil {
		return nil, err
	}
	err = f.prepareInfo()
	if err != nil {
		return nil, err
	}
	err = f.packPageInfo()
	if err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id should greater than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	//并发获取topic和post信息
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		f.topic = repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
	}()
	go func() {
		defer wg.Done()
		f.postList = repository.NewPostDaoInstance().QueryPostsByParentId(f.topicId)
	}()
	wg.Wait()
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.postList,
	}
	return nil
}
