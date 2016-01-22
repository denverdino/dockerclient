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

// Address represents an IP address
type Address struct {
	Addr      string
	PrefixLen int
}
