package cotroller

import (
	"strconv"

	"github.com/Moonlight-Zhao/go-project-example/service"
)

type PageData struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
func QueryPageInfo(topicIdStr string) *PageData {
	//通过topicIdStr来获取Pageinfo 如果解析出错 直接err
	topicId, err := strconv.ParseInt(topicIdStr, 10, 64)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	//解析正常 通过server来获取PageInfo
	pageInfo, err := service.QueryPageInfo(topicId)
	if err != nil {
		return &PageData{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: pageInfo,
	}

}
