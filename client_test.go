package vindalu

import (
	"testing"
)

var (
	testUrl = "http://localhost:5454"
)

func Test_NewClient(t *testing.T) {
	c, err := NewClient(testUrl)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%#v", *c.Config)
}
