package setup

import "context"

type Setup struct {
	ctx context.Context
}

func New(ctx context.Context) Setup {
	return Setup{
		ctx: ctx,
	}
}
