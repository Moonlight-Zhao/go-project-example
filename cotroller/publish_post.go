package cotroller

import (
	"github.com/Moonlight-Zhao/go-project-example/service"
	"strconv"
)


func PublishPost(topicIdStr, content string) *PageData {
	//参数转换
	topicId, _ := strconv.ParseInt(topicIdStr, 10, 64)
	//获取service层结果
	postId, err := service.PublishPost(topicId, content)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: map[string]int64{
			"post_id": postId,
		},
	}

}
