package gosensors

import "testing"

func TestInit(t *testing.T) {
	err := Init()
	if err != nil {
		t.Fatalf("Init() failed with error: %v", err)
	}

	Cleanup()
}
