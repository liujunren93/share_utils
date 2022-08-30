package redis

import (
	"context"
	"time"
)

func (c *Client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) {
	if c.Mode == 0 {
		c.Client.Set(ctx, key, value, expiration)
	}
	c.ClusterClient.Set(ctx, key, value, expiration)
}
