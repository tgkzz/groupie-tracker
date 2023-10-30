package models

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

type DateLocation struct {
	DateLoc map[string][]string `json:"datesLocations"`
}

type Locations struct {
	Locate   []string `json:"locations"`
	DateLink string   `json:"dates"`
}

type Concerts struct {
	Location string
	Dates    []string
}

type ResultGroup struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	ConcertData  []Concerts
}

type Groups struct {
	Groups []Group
}

type Filter struct {
	CreationDateFrom int
	CreationDateTo   int
	FirstAlbumFrom   int
	FirstAlbumTo     int
	Members          []int
	Location         string
}
