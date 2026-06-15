package models

// Artist represents the structural layout of a single musical act.
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	Dates        string   `json:"dates"`
	Relations    string   `json:"relations"`
}

// LocationsResponse represents the index wrapper returned by the /locations endpoint
type LocationsResponse struct {
	Index []Location `json:"index"`
}

// Location handles mapping individual geographical gig historical tracks
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// DatesResponse represents the index wrapper returned by the /dates endpoint
type DatesResponse struct {
	Index []Date `json:"index"`
}

// Date holds arrays of past/upcoming event date strings
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// RelationsResponse represents the index wrapper returned by the /lation endpoint
type RelationsResponse struct {
	Index []Relation `json:"index"`
}

// Relation explixitly maps a string location value directly to its corresponding concert schedule arrays
type Relation struct {
	ID        int      `json:"id"`
	Relations []string `json:"relations"`
}

// ArtistFull aggregates relational indices into a composite data transfer object
type ArtistFull struct {
	Artist    Artist
	Locations []string
	Dates     []string
	Relations map[string][]string
}