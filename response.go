package bunnystorage

import "net/http"

// Response represents a response from the BunnyCDN Storage API.
type Response struct {
	// Header contains the response headers.
	Header http.Header

	// Body contains the response body as a byte slice.
	Body []byte

	// Status is the HTTP status code of the response.
	Status int
}

// Object represents a file or directory in the BunnyCDN Storage API.
type Object struct {
	UserID          string `json:"UserId,omitempty"`
	ContentType     string `json:"ContentType,omitempty"`
	Path            string `json:"Path,omitempty"`
	ObjectName      string `json:"ObjectName,omitempty"`
	ReplicatedZones string `json:"ReplicatedZones,omitempty"`
	LastChanged     string `json:"LastChanged,omitempty"`
	StorageZoneName string `json:"StorageZoneName,omitempty"`
	Checksum        string `json:"Checksum,omitempty"`
	DateCreated     string `json:"DateCreated,omitempty"`
	GUID            string `json:"Guid,omitempty"`
	Length          int    `json:"Length,omitempty"`
	ServerID        int    `json:"ServerId,omitempty"`
	StorageZoneID   int    `json:"StorageZoneId,omitempty"`
	ArrayNumber     int    `json:"ArrayNumber,omitempty"`
	IsDirectory     bool   `json:"IsDirectory,omitempty"`
}
