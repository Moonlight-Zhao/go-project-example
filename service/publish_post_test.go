package service

import (
	"bou.ke/monkey"
	"github.com/Moonlight-Zhao/go-project-example/repository"
	"github.com/Moonlight-Zhao/go-project-example/util"
	"github.com/stretchr/testify/assert"
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	if err := repository.Init(); err != nil {
		os.Exit(1)
	}
	if err := util.InitLogger(); err != nil {
		os.Exit(1)
	}
	m.Run()
}

func TestPublishPost(t *testing.T) {

	type args struct {
		topicId int64
		userId  int64
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "测试发布回帖",
			args: args{
				topicId: 1,
				userId:  2,
				content: "再次回帖",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := PublishPost(tt.args.topicId, tt.args.userId, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublishPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestPublishPostMockDao(t *testing.T) {
	var postDao *repository.PostDao
	monkey.PatchInstanceMethod(reflect.TypeOf(postDao), "CreatePost",
		func(_ *repository.PostDao, post *repository.Post) error {
			post.Id = 100
			return nil
		})
	defer monkey.UnpatchInstanceMethod(reflect.TypeOf(postDao), "CreatePost")
	do, _ := NewPublishPostFlow(1, 2, "mock测试").Do()
	assert.Equal(t, do, int64(100))
}
