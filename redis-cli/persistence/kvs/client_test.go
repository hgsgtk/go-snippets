package kvs_test

import (
	"testing"
	"time"

	"github.com/higasgt/go-snippets/redis-cli/persistence/kvs"
	"github.com/higasgt/go-snippets/redis-cli/testhelper"
)

func TestSetToken(t *testing.T) {
	client := testhelper.NewMockRedis(t)

	if err := kvs.SetToken(client, "test", 1); err != nil {
		t.Fatalf("unexpected error while SetToken '%#v'", err)
	}
	actual, err := client.Get("TOKEN_test").Result()
	if err != nil {
		t.Fatalf("unexpected error while get value '%#v'", err)
	}

	if expected := "1"; expected != actual {
		t.Errorf("expected value '%s', actual value '%s'", expected, actual)
	}
}

func TestGetIDByToken(t *testing.T) {
	client := testhelper.NewMockRedis(t)
	if err := client.Set("TOKEN_test", 1, time.Second*1000).Err(); err != nil {
		t.Fatalf("unexpected error while set test data '%#v'", err)
	}

	actual, err := kvs.GetIDByToken(client, "test")
	if err != nil {
		t.Fatalf("unexpected error while GetIDByToken '%#v'", err)
	}
	if expected := 1; expected != actual {
		t.Errorf("expected value '%#v', actual value '%#v'", expected, actual)
	}
}
