package godp

import (
	"encoding/json"
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

func TestComparsion(t *testing.T) {
	sampleResponse := APIResponse{
		RequestArgs: map[string]interface{}{
			"query":    "example-query",
			"page":     1,
			"pageSize": 10,
		},
		Length: 2,
		Packages: []PackageInfo{
			{
				Name:      "example-package-1",
				Epoch:     0,
				Version:   "1.0.0",
				Release:   "1",
				Arch:      "x86_64",
				Disttag:   "f32",
				Buildtime: 1594640000,
				Source:    "example-source-1",
			},
			{
				Name:      "example-package-2",
				Epoch:     1,
				Version:   "2.1.4",
				Release:   "3",
				Arch:      "x86_64",
				Disttag:   "f33",
				Buildtime: 1605225600,
				Source:    "example-source-2",
			},
		},
	}

	sampleResponse1 := APIResponse{
		RequestArgs: map[string]interface{}{
			"query":    "example-query",
			"page":     1,
			"pageSize": 10,
		},
		Length: 2,
		Packages: []PackageInfo{
			{
				Name:      "example-package-1",
				Epoch:     0,
				Version:   "1.0.0",
				Release:   "1",
				Arch:      "x86_64",
				Disttag:   "f32",
				Buildtime: 1594640000,
				Source:    "example-source-1",
			},
			{
				Name:      "example-package-3",
				Epoch:     1,
				Version:   "2.1.3",
				Release:   "3",
				Arch:      "x86_64",
				Disttag:   "f33",
				Buildtime: 1605225600,
				Source:    "example-source-2",
			},
		},
	}

	result := ComparePackages(sampleResponse, sampleResponse1)

	rJson, err := json.Marshal(result)
	if err != nil {
		t.Error(err)
	}

	expected := `[{"arch":"x86_64","in_second_not_in_first":[{"name":"example-package-3","epoch":1,"version":"2.1.3","release":"3","arch":"x86_64","disttag":"f33","buildtime":1605225600,"source":"example-source-2"}],"in_first_not_in_second":[{"name":"example-package-2","epoch":1,"version":"2.1.4","release":"3","arch":"x86_64","disttag":"f33","buildtime":1605225600,"source":"example-source-2"}],"higher_in_first":null}]`
	// {
	// 	"arch":"x86_64",
	// 	"in_second_not_in_first":
	// 	[
	// 		{
	// 			"name":"example-package-3",
	// 			"epoch":1,
	// 			"version":"2.1.3",
	// 			"release":"3",
	// 			"arch":"x86_64",
	// 			"disttag":"f33",
	// 			"buildtime":1605225600,
	// 			"source":"example-source-2"
	// 		}
	// 	],
	// 	"in_first_not_in_second":
	// 	[
	// 		{
	// 			"name":"example-package-2",
	// 			"epoch":1,
	// 			"version":"2.1.4",
	// 			"release":"3",
	// 			"arch":"x86_64",
	// 			"disttag":"f33",
	// 			"buildtime":1605225600,
	// 			"source":"example-source-2"
	// 		}
	// 	],
	// 	"higher_in_first":null
	// }

	if string(rJson) != expected {
		t.Error("wrong result")
	}
}
