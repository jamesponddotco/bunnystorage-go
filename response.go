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
	GUID            string `json:"Guid,omitempty"`
	StorageZoneName string `json:"StorageZoneName,omitempty"`
	Path            string `json:"Path,omitempty"`
	ObjectName      string `json:"ObjectName,omitempty"`
	LastChanged     string `json:"LastChanged,omitempty"`
	UserID          string `json:"UserId,omitempty"`
	DateCreated     string `json:"DateCreated,omitempty"`
	IsDirectory     bool   `json:"IsDirectory,omitempty"`
	Length          int    `json:"Length,omitempty"`
	ServerID        int    `json:"ServerId,omitempty"`
	StorageZoneID   int    `json:"StorageZoneId,omitempty"`
}
