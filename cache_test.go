package wechat

import (
	"context"
	"testing"
	"time"
)

func TestMemCache_Set(t *testing.T) {
	type args struct {
		key        string
		value      interface{}
		expiration time.Duration
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				key:        "key",
				value:      "test",
				expiration: DefaultCacheExpiration,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := []MemCacheOptFunc{}
			mc, err := NewMemCache(options...)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			mc.Set(ctx, tt.args.key, tt.args.value, tt.args.expiration)
			value, err := mc.Get(ctx, tt.args.key)
			if err != nil || value.(string) != tt.args.value.(string) {
				t.Log(value, err)
				t.FailNow()
			}
			t.Log("value", value)
			mc.Delete(ctx, tt.args.key)
			if _, err := mc.Get(ctx, tt.args.key); err != nil && err != ErrCacheKeyNotExist {
				t.Log("value should be deleted")
				t.FailNow()
			}
		})
	}
}
