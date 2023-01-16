package service

import (
	"errors"
	"fmt"
	"github.com/Moonlight-Zhao/go-project-example/repository"
	"sync"
)

type TopicInfo struct {
	Topic *repository.Topic
	User  *repository.User
}

type PostInfo struct {
	Post *repository.Post
	User *repository.User
}

type PageInfo struct {
	TopicInfo *TopicInfo
	PostList  []*PostInfo
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

// 新建QueryPageInfoFlow结构体
func NewQueryPageInfoFlow(topId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topId,
	}
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo
	// 数据库对于DO
	topic   *repository.Topic
	posts   []*repository.Post
	userMap map[int64]*repository.User
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	// 校验参数
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	// 查询初始信息
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	// 组合页面信息
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	//根据topicId获取topic信息
	var wg sync.WaitGroup
	wg.Add(2)
	var topicErr, postErr error
	go func() {
		defer wg.Done()
		topic, err := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		if err != nil {
			topicErr = err
			return
		}
		f.topic = topic
	}()
	//根据topicId获取post列表
	go func() {
		defer wg.Done()
		posts, err := repository.NewPostDaoInstance().QueryPostByParentId(f.topicId)
		if err != nil {
			postErr = err
			return
		}
		f.posts = posts
	}()
	wg.Wait()
	if topicErr != nil {
		return topicErr
	}
	if postErr != nil {
		return postErr
	}
	//获取用户信息
	uids := []int64{f.topic.UserId}
	for _, post := range f.posts {
		uids = append(uids, post.UserId)
	}
	// 查询topic所有相关用户（topic用户和post用户）
	userMap, err := repository.NewUserDaoInstance().MQueryUserById(uids)
	if err != nil {
		return err
	}
	f.userMap = userMap
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	//topic info
	userMap := f.userMap
	topicUser, ok := userMap[f.topic.UserId]
	if !ok {
		return errors.New("has no topic user info")
	}
	//post list （组建post信息和发post用户信息）
	postList := make([]*PostInfo, 0)
	for _, post := range f.posts {
		postUser, ok := userMap[post.UserId]
		if !ok {
			return errors.New("has no post user info for " + fmt.Sprint(post.UserId))
		}
		postList = append(postList, &PostInfo{
			Post: post,
			User: postUser,
		})
	}
	f.pageInfo = &PageInfo{
		TopicInfo: &TopicInfo{
			Topic: f.topic,
			User:  topicUser,
		},
		PostList: postList,
	}
	return nil
}
