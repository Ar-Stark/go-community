package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//数据存储在本地文件中，建立索引，提升读取性能

var (
	topicIndexMap map[int64]*Topic
	postIndexMap  map[int64][]*Post
)

func Init(filePath string) error {
	err := initTopicIndexMap(filePath)
	if err != nil {
		return err
	}
	err = initPostIndexMap(filePath)
	if err != nil {
		return err
	}
	return nil
}

func initTopicIndexMap(filePath string) error {
	file, err := os.Open(filePath + "topic")
	if err != nil {
		return fmt.Errorf("open topic file failed: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	topicTmpMap := make(map[int64]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		err = json.Unmarshal([]byte(text), &topic)
		if err != nil {
			return fmt.Errorf("unmarshal topic failed: %v", err)
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}

func initPostIndexMap(filePath string) error {
	file, err := os.Open(filePath + "post")
	if err != nil {
		return fmt.Errorf("open post file failed: %v", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	postTmpMap := make(map[int64][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		err = json.Unmarshal([]byte(text), &post)
		if err != nil {
			return fmt.Errorf("unmarshal post failed: %v", err)
		}
		posts, ok := postTmpMap[post.ParentId]
		if !ok {
			posts = make([]*Post, 0)
		}
		posts = append(posts, &post)
		postTmpMap[post.ParentId] = posts
	}
	postIndexMap = postTmpMap
	return nil
}
