package logic

import (
	"context"

	"gozero_single/greet/internal/svc"
	"gozero_single/greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GreetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) GreetLogic {
	return GreetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetLogic) Greet() (*types.Response, error) {

	return &types.Response{
		Message: "hello",
		Name:    "lee",
		Age:     28,
	}, nil
}
