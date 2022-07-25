package server

import (
	"context"
)

func (c *GRPCServer) Get(ctx context.Context, in *NilBody) (*NilBody, error) {
	return nil, nil
}
