package store

import (
	"context"
	"testing"
)

func TestRedisStore_SetGet(t *testing.T){
	store := NewRedisStore("localhsot:6379", "")
	ctx := context.Background()

	key := "test-key"
	value := "test-value"

	if err := store.Set(ctx, key, value); err != nil {
		t.Fatalf("Set failed : %v", err)
	}

	got, err := store.Get(ctx, key)
	if err != nil {
		t.Fatalf("Get failed %v ", err)
	}

	if got != value{
		t.Errorf("预期值 %s, 实际值 %s", value, got)
	}

}