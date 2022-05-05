package service

import (
	"errors"
	"time"
	"unicode/utf16"

	"github.com/Moonlight-Zhao/go-project-example/repository"
	idworker "github.com/gitstliu/go-id-worker"
)

var idGen *idworker.IdWorker

func init() {
	idGen = &idworker.IdWorker{}
	idGen.InitIdWorker(1, 1)
}

func PublishPost(topicId int64, content string) (int64, error) {
	return NewPublishPostFlow(topicId, content).Do()
}

func NewPublishPostFlow(topicId int64, content string) *PublishPostFlow {
	return &PublishPostFlow{
		content: content,
		topicId: topicId,
	}
}

type PublishPostFlow struct {
	content string
	topicId int64

	postId int64
}

func (f *PublishPostFlow) Do() (int64, error) {
	if err := f.checkParam(); err != nil {
		return 0, err
	}
	if err := f.publish(); err != nil {
		return 0, err
	}
	return f.postId, nil
}

func (f *PublishPostFlow) checkParam() error {
	if len(utf16.Encode([]rune(f.content))) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

func (f *PublishPostFlow) publish() error {
	post := &repository.Post{
		ParentId:   f.topicId,
		Content:    f.content,
		CreateTime: time.Now().Unix(),
	}
	id, err := idGen.NextId()
	if err != nil {
		return err
	}
	post.Id = id
	if err := repository.NewPostDaoInstance().InsertPost(post); err != nil {
		return err
	}
	f.postId = post.Id
	return nil
}
