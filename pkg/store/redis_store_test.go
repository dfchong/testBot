package store

import "testing"

func TestRedisStore_SetGet(t *testing.T){
	store := NewRedisStore("localhsot:6379", "")
	
}