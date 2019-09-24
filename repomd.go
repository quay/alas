package alas

import "errors"

type RepoType string

var ErrRepoNotFound = errors.New("Repo not found")

const (
	PrimaryDB  RepoType = "primary_db"
	OtherDB    RepoType = "other_db"
	GroupGZ    RepoType = "group_gz"
	Group      RepoType = "group"
	FileLists  RepoType = "filelists_db"
	UpdateInfo RepoType = "updateinfo"
)

type RepoMD struct {
	XMLNS    string `xml:"xmlns,attr"`
	XMLRPM   string `xml:"xmlns rpm,attr"`
	Revision int    `xml:"revision"`
	RepoList []Repo `xml:"data"`
}

type Repo struct {
	Type            string   `xml:"type,attr"`
	Checksum        Checksum `xml:"checksum"`
	OpenChecksum    Checksum `xml:"open-checksum"`
	Location        Location `xml:"location"`
	Timestamp       int      `xml:"timestamp"`
	DatabaseVersion int      `xml:"database_version"`
	Size            int      `xml:"size"`
	OpenSize        int      `xml:"open-size"`
}

type Checksum struct {
	Sum  string `xml:",chardata"`
	Type string `xml:"type,attr"`
}

type Location struct {
	Href string `xml:"href,attr"`
}

func (md *RepoMD) Repo(t RepoType) (Repo, error) {
	var repo Repo
	for _, repo := range md.RepoList {
		if repo.Type == string(t) {
			return repo, nil
		}
	}
	return repo, ErrRepoNotFound
}
