package models

type Err struct {
	Text_err string
	Code_err int
}

type Group struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location     string   `json:"location"`
	ConcertDate  string   `json:"concertDate"`
	Relations    string   `json:"relations"`
}

type Groups struct {
	Groups []Group
}
