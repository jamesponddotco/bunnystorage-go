package bunnystorage

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
)

// Endpoint represents the primary storage region of a storage zone.
type Endpoint int

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
	default:
		return "https://storage.bunnycdn.com"
	}
}

// IsValid returns true if the endpoint is a valid Bunny.net endpoint.
func (e Endpoint) IsValid() bool {
	switch e {
	case EndpointFalkenstein, EndpointNewYork, EndpointLosAngeles, EndpointSingapore, EndpointSydney, EndpointLondon, EndpointStockholm, EndpointSaoPaulo, EndpointJohannesburg:
		return true
	default:
		return false
	}
}
