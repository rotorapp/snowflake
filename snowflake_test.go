package snowflake

import (
	"testing"
)

func TestAdjacentSnowflakesDontCollide(t *testing.T) {
	Init(0)

	ids := make(map[Snowflake]bool)

	for i := 0; i < 1000000; i++ {
		id := New()
		if ids[id] {
			t.Fatalf("Collision detected for ID: %d", id)
		}
		ids[id] = true
	}
}

func TestGeneratorIDBounds(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for out-of-bounds generator ID")
		}
	}()

	Init(5000) // This should panic
}

func TestSnowflakeComponents(t *testing.T) {
	Init(1234)

	sf := New()

	if sf.GeneratorID() != 1234 {
		t.Errorf("Expected GeneratorID to be 1234, got %d", sf.GeneratorID())
	}

	if sf.Increment() != 0 {
		t.Errorf("Expected Increment to be 0, got %d", sf.Increment())
	}

	sf2 := New()
	if sf2.Increment() != 1 {
		t.Errorf("Expected Increment to be 1, got %d", sf2.Increment())
	}
}
