package snowflake

import (
	"testing"
)

func TestAdjacentSnowflakesDontCollide(t *testing.T) {
	Init(0)

	ids := make(map[int64]bool)

	for i := 0; i < 1000000; i++ {
		id := New()
		if ids[id] {
			t.Fatalf("Collision detected for ID: %d", id)
		}
		ids[id] = true
	}
}
