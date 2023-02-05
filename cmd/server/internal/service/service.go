package service

import (
	"context"

	"github.com/ea3hsp/mclog/pkg/mclog"
)

type mclogService struct {
	mclog.UnimplementedLogServiceServer
}

func (mclogService) Log(context.Context, *mclog.LogReq) (*mclog.LogRes, error) {

}
