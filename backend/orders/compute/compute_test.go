package orders

import (
	"testing"
)

func TestConsume(t *testing.T) {
	// Given

	// When
	computed, err := ComputeStats("9c98d18da506527d1292128b6205a08a94c3be204ef75e0a810fad400edac0dfd9fb97265ad44e722e3c9ad32267944bf01f2465ae7775f62ec87905f663d5e8c93c8d185e242f7391d4c16af8a5990827356a5954ca16c620da3ce0132b037665cbcda92fe24dfdcd8d50d8582c6b0b")

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	t.Logf("computed: %+v\n", computed)

}
