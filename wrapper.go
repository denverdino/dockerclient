package dockerclient

// Settings stores configuration details about the daemon network config
// TODO Windows. Many of these fields can be factored out.,
type NetworkSettings struct {
	IPAddress              string
	IPPrefixLen            int
	Gateway                string
	Bridge                 string
	Ports                  map[string][]PortBinding
	SandboxID              string
	HairpinMode            bool
	LinkLocalIPv6Address   string
	LinkLocalIPv6PrefixLen int
	Networks               map[string]*EndpointSettings
	SandboxKey             string
	SecondaryIPAddresses   []Address
	SecondaryIPv6Addresses []Address
	IsAnonymousEndpoint    bool
}

// EndpointSettings stores the network endpoint details
type EndpointSettings struct {
	EndpointID          string
	Gateway             string
	IPAddress           string
	IPPrefixLen         int
	IPv6Gateway         string
	GlobalIPv6Address   string
	GlobalIPv6PrefixLen int
	MacAddress          string
}

// Address represents an IP address
type Address struct {
	Addr      string
	PrefixLen int
}
