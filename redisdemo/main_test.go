package redisdemo

import (
	"context"
	"testing"
)

var redisClient = NewClient()

func TestGetFromRedis(t *testing.T) {
	ctx := context.Background()

	key1, value1 := "cat-01", "hudi"
	key2, value2 := "cat-02", "jiujiu"

	SaveToRedis(redisClient, key1, value1, ctx)
	SaveToRedis(redisClient, key2, value2, ctx)

	tests := []struct {
		name string
		key  string
		want string
	}{
		{"find_01", key1, value1},
		{"find_02", key2, value2},
		{"not_find_01", "water", ""},
		{"not_find_02", "book", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFromRedis(redisClient, tt.key, ctx); got != tt.want {
				t.Errorf("GetFromRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}
