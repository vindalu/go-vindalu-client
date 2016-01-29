package vindalu

import (
	"testing"
)

var (
	testUrl  = "http://localhost:5454"
	testCred = "test"
)

func Test_NewClient(t *testing.T) {
	c, err := NewClient(testUrl)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%#v", *c.Config)
}

func Test_SetCredentials(t *testing.T) {
	c, _ := NewClient(testUrl)
	c.SetCredentials(testCred, testCred)
	if c.creds.Username != "test" {
		t.Fatalf("Cred user (%s) does not match test user (%s).", c.creds.Username, testCred)
	}
	t.Logf("%#v", c.creds)
}
