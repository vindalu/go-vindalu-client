package vindalu

import (
	"testing"
)

var (
	testClient, _  = NewClient(testUrl)
	testAType      = "vserver"
	testId         = "test.vserver1"
	testData       = map[string]string{"status": "enabled", "environment": "production"}
	testUpdateData = map[string]string{"status": "disabled", "testfield1": "testval1"}
)

func Test_Create(t *testing.T) {
	id, err := testClient.Create(testAType, testId, testData)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%#v", id)
}

func Test_Get(t *testing.T) {

	ba, err := testClient.Get(testAType, testId, 0)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%v", ba)
}

func Test_Update(t *testing.T) {
	id, err := testClient.Update(testAType, testId, testUpdateData)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%#v", id)
}

func Test_GetVersions(t *testing.T) {

	versions, err := testClient.GetVersions(testAType, testId)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("Version: %d", len(versions))
}

func Test_GetVersionDiffs(t *testing.T) {

	versions, err := testClient.GetVersions(testAType, testId)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("Version diffs: %d", len(versions))
}

func Test_Get_version(t *testing.T) {

	ba, err := testClient.Get(testAType, testId, 1)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%v", ba)
}

func Test_Delete(t *testing.T) {
	id, err := testClient.Delete(testAType, testId)
	if err != nil {
		t.Fatalf("%s", err)
	}
	t.Logf("%#v", id)
}

func Test_GetTypes(t *testing.T) {
	_, err := testClient.GetTypes()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func Test_ListTypeProperties(t *testing.T) {
	_, err := testClient.ListTypeProperties(testAType)
	if err != nil {
		t.Fatalf("%s", err)
	}
}
