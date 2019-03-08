package wechat

import (
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
				value:      "value",
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
			mc.Set(tt.args.key, tt.args.value, tt.args.expiration)
			if value, exist := mc.Get(tt.args.key); !exist || value.(string) != tt.args.value.(string) {
				t.Log(value, exist)
				t.FailNow()
			}
			mc.Delete(tt.args.key)
			if _, exist := mc.Get(tt.args.key); exist {
				t.Log("value should be deleted")
				t.FailNow()
			}
		})
	}
}
