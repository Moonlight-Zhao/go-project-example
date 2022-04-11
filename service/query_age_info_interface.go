package service

import (
	"time"
)

type QueryPageFlow interface {
	checkParam(ctx *QueryPageCtx) error
	prepareInfo(ctx *QueryPageCtx) error
	packPageInfo(ctx *QueryPageCtx) error
}

type QueryPageFlowA struct {
}

func (q *QueryPageFlowA) checkParam(ctx *QueryPageCtx) error {
	panic("implement me")
}

func (q *QueryPageFlowA) prepareInfo(ctx *QueryPageCtx) error {
	panic("implement me")
}

func (q *QueryPageFlowA) packPageInfo(ctx *QueryPageCtx) error {
	panic("implement me")
}

type QueryPageFlowB struct {
}

func (q *QueryPageFlowB) checkParam(ctx *QueryPageCtx) error {
	panic("implement me")
}

func (q *QueryPageFlowB) prepareInfo(ctx *QueryPageCtx) error {
	panic("implement me")
}

func (q *QueryPageFlowB) packPageInfo(ctx *QueryPageCtx) error {
	panic("implement me")
}

type PageResult struct {

}

type QueryPageCtx struct {
	*PageResult
}

type QueryPage struct {

}

func (*QueryPage) Do(ctx *QueryPageCtx) (*PageResult,error) {
	flow := getQueryPageFlow()
	if err := flow.checkParam(ctx); err != nil {
		return nil,err
	}
	if err := flow.prepareInfo(ctx); err != nil {
		return nil,err
	}
	if err := flow.packPageInfo(ctx); err != nil {
		return nil,err
	}
	return ctx.PageResult,nil
}

func getQueryPageFlow() QueryPageFlow {
	if time.Now().Unix()%2 == 1 {
		return &QueryPageFlowA{}
	}
	return &QueryPageFlowB{}
}
