package terraform_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/massdriver-cloud/airlock/pkg/terraform"
)

func TestSchemaToTf(t *testing.T) {
	type testData struct {
		name string
	}
	tests := []testData{
		{
			name: "default",
		},
		{
			name: "simple",
		},
		{
			name: "dependencies",
		},
		{
			name: "dynamics",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			want, err := os.ReadFile(filepath.Join("testdata/schemas", tc.name+".tf"))
			if err != nil {
				t.Fatalf("%d, unexpected error", err)
			}

			schemaFile, err := os.Open(filepath.Join("testdata/schemas", tc.name+".json"))
			if err != nil {
				t.Fatalf("%d, unexpected error", err)
			}

			got, err := terraform.SchemaToTf(schemaFile)
			if err != nil {
				t.Fatalf("%d, unexpected error", err)
			}

			if string(got) != string(want) {
				t.Fatalf("got %q want %q", string(got), string(want))
			}
		})
	}
}
