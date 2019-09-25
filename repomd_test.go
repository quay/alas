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
		repo, err := repomdRoot.Repo(test, "")
		assert.NoError(t, err)
		assert.NotEmpty(t, repo)
	}
}

func Test_RepoMD_FQDN(t *testing.T) {
	tests := []struct {
		repoType     RepoType
		expectedFQDN string
	}{
		{
			PrimaryDB,
			"http://test-mirror/repodata/primary.sqlite.bz2",
		}, {
			OtherDB,
			"http://test-mirror/repodata/other.sqlite.bz2",
		}, {
			GroupGZ,
			"http://test-mirror/repodata/comps.xml.gz",
		}, {
			Group,
			"http://test-mirror/repodata/comps.xml",
		}, {
			FileLists,
			"http://test-mirror/repodata/filelists.sqlite.bz2",
		}, {
			UpdateInfo,
			"http://test-mirror/repodata/updateinfo.xml.gz",
		},
	}

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
		repo, err := repomdRoot.Repo(test.repoType, "http://test-mirror/")
		assert.NoError(t, err)
		assert.NotEmpty(t, repo)
		assert.Equal(t, test.expectedFQDN, repo.Location.Href)
	}
}
