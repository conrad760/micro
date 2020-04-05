package data

import "testing"

func TestNameIsRequired(t *testing.T) {
	p := &Product{
		Name:  "conrad",
		Price: 100,
		SKU:   "abc123",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
