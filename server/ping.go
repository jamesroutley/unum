package server

import (
	"context"

	"github.com/jamesroutley/unum/unumpb"
)

func (s Server) Ping(ctx context.Context, req *unumpb.PingRequest) (*unumpb.PingResponse, error) {
	text := req.Text

	return &unumpb.PingResponse{
		Text: text,
	}, nil
}
