package queue

import (
	"context"

	"github.com/hibiken/asynq"
)

type Client struct {
	c *asynq.Client
}

// NewClient builds an Asynq client. Caller closes via Close().
func NewClient(redisAddr, redisPassword string, db int) *Client {
	return &Client{
		c: asynq.NewClient(asynq.RedisClientOpt{
			Addr: redisAddr, Password: redisPassword, DB: db,
		}),
	}
}

func (c *Client) Close() error { return c.c.Close() }

func (c *Client) Enqueue(ctx context.Context, t *asynq.Task, opts ...asynq.Option) error {
	_, err := c.c.EnqueueContext(ctx, t, opts...)
	return err
}
