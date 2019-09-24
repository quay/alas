package alas

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RepoMD_Parse(t *testing.T) {
	path := filepath.Join("testdata", "test_repomd.xml")
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open test data: %v", err)
	}

	repomdRoot := &RepoMD{}
	err = xml.NewDecoder(f).Decode(repomdRoot)
	if err != nil {
		t.Fatalf("failed to parse repomd test data into struct: %v", err)
	}

	assert.NotNil(t, repomdRoot)
	assert.Len(t, repomdRoot.RepoList, 6)
}

func Test_RepoMD_GetRepo(t *testing.T) {
	tests := []RepoType{PrimaryDB, OtherDB, GroupGZ, Group, FileLists, UpdateInfo}

	path := filepath.Join("testdata", "test_repomd.xml")
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open test data: %v", err)
	}

	repomdRoot := &RepoMD{}
	err = xml.NewDecoder(f).Decode(repomdRoot)
	if err != nil {
		t.Fatalf("failed to parse repomd test data into struct: %v", err)
	}

	for _, test := range tests {
		repo, err := repomdRoot.Repo(test)
		assert.NoError(t, err)
		assert.NotEmpty(t, repo)
	}

}
