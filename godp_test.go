package godp

import (
	"os"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestSerde(t *testing.T) {
	type Sample struct {
		Name string
	}
	s := Sample{"test"}

	err := SerializeData(s, "test")
	if err != nil {
		t.Error(err)
	}

	s1, err := DeserializeData[Sample]("test")

	if err != nil {
		t.Error(err)
	}

	if s.Name != s1.Name {
		t.Error("Deserialization failed")
	}

	err = os.Remove("test")
	if err != nil {
		t.Error(err)
	}
}

func TestClient(t *testing.T) {
	_, err := FetchPackagesFromAPI[APIResponse]("https://rdb.altlinux.org/api/export/branch_binary_packages", "p10")
	if err != nil {
		t.Error(err)
	}
}
