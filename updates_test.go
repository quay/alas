package alas

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Updates_Parse(t *testing.T) {
	path := filepath.Join("testdata", "test_updateinfo.xml")
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open test data: %v", err)
	}

	updates := &Updates{}
	err = xml.NewDecoder(f).Decode(updates)
	if err != nil {
		t.Fatalf("failed to parse updateinfo test data into struct: %v", err)
	}

	assert.NotNil(t, updates)
	assert.NotEmpty(t, updates.Updates)
}
