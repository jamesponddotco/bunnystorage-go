package bunnystorage

import "net/url"

// The following endpoints are available for use with the Edge Storage API.
const (
	EndpointFalkenstein Endpoint = iota + 1
	EndpointNewYork
	EndpointLosAngeles
	EndpointSingapore
	EndpointSydney
	EndpointLondon
	EndpointStockholm
	EndpointSaoPaulo
	EndpointJohannesburg
	EndpointLocalhost // For testing purposes only
)

// Endpoint represents the primary storage region of a storage zone.
type Endpoint int

// Parse parses a string representation of an endpoint into an Endpoint. If the
// string is not a valid endpoint, EndpointFalkenstein is returned.
func Parse(s string) Endpoint {
	uri, err := url.Parse(s)
	if err != nil {
		return EndpointFalkenstein
	}

	switch uri.Host {
	case "storage.bunnycdn.com":
		return EndpointFalkenstein
	case "ny.storage.bunnycdn.com":
		return EndpointNewYork
	case "la.storage.bunnycdn.com":
		return EndpointLosAngeles
	case "sg.storage.bunnycdn.com":
		return EndpointSingapore
	case "syd.storage.bunnycdn.com":
		return EndpointSydney
	case "uk.storage.bunnycdn.com":
		return EndpointLondon
	case "se.storage.bunnycdn.com":
		return EndpointStockholm
	case "br.storage.bunnycdn.com":
		return EndpointSaoPaulo
	case "jh.storage.bunnycdn.com":
		return EndpointJohannesburg
	case "localhost:62769":
		return EndpointLocalhost
	default:
		return EndpointFalkenstein
	}
}

// String returns the string representation of the endpoint.
func (e Endpoint) String() string {
	switch e {
	case EndpointFalkenstein:
		return "https://storage.bunnycdn.com"
	case EndpointNewYork:
		return "https://ny.storage.bunnycdn.com"
	case EndpointLosAngeles:
		return "https://la.storage.bunnycdn.com"
	case EndpointSingapore:
		return "https://sg.storage.bunnycdn.com"
	case EndpointSydney:
		return "https://syd.storage.bunnycdn.com"
	case EndpointLondon:
		return "https://uk.storage.bunnycdn.com"
	case EndpointStockholm:
		return "https://se.storage.bunnycdn.com"
	case EndpointSaoPaulo:
		return "https://br.storage.bunnycdn.com"
	case EndpointJohannesburg:
		return "https://jh.storage.bunnycdn.com"
	case EndpointLocalhost:
		return "http://localhost:62769"
	default:
		return "https://storage.bunnycdn.com"
	}
}

// IsValid returns true if the endpoint is a valid Bunny.net endpoint.
func (e Endpoint) IsValid() bool {
	switch e {
	case EndpointFalkenstein, EndpointNewYork, EndpointLosAngeles, EndpointSingapore, EndpointSydney, EndpointLondon, EndpointStockholm, EndpointSaoPaulo, EndpointJohannesburg, EndpointLocalhost:
		return true
	default:
		return false
	}
}
