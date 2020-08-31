package testing_test

import (
	"testing"
)

func TestDeadlineConfirm(t *testing.T) {
	d, ok := t.Deadline()
	t.Logf("t.Deadline() = %v, timeout set = %v;", d, ok)
}
