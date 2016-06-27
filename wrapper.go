package dockerclient

import (
	"fmt"
	"encoding/json"
)

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

// UpdateConfig holds the mutable attributes of a Container.
// Those attributes can be updated at runtime.
type UpdateConfig struct {
	// Applicable to all platforms
	CPUShares int64 `json:"CpuShares"` // CPU shares (relative weight vs. other containers)
	Memory    int64 // Memory limit (in bytes)

	// Applicable to UNIX platforms
	CgroupParent         string // Parent cgroup.
	BlkioWeight          uint16 // Block IO weight (relative weight vs. other containers)
	BlkioWeightDevice    []*WeightDevice
	BlkioDeviceReadBps   []*ThrottleDevice
	BlkioDeviceWriteBps  []*ThrottleDevice
	BlkioDeviceReadIOps  []*ThrottleDevice
	BlkioDeviceWriteIOps []*ThrottleDevice
	CPUPeriod            int64           `json:"CpuPeriod"` // CPU CFS (Completely Fair Scheduler) period
	CPUQuota             int64           `json:"CpuQuota"`  // CPU CFS (Completely Fair Scheduler) quota
	CpusetCpus           string          // CpusetCpus 0-2, 0,1
	CpusetMems           string          // CpusetMems 0-2, 0,1
	Devices              []DeviceMapping // List of devices to map inside the container
	DiskQuota            int64           // Disk limit (in bytes)
	KernelMemory         int64           // Kernel memory limit (in bytes)
	MemoryReservation    int64           // Memory soft limit (in bytes)
	MemorySwap           int64           // Total memory usage (memory + swap); set `-1` to enable unlimited swap
	MemorySwappiness     *int64          // Tuning container memory swappiness behaviour
	OomKillDisable       *bool           // Whether to disable OOM Killer or not
	PidsLimit            int64           // Setting pids limit for a container
	Ulimits              []*Ulimit // List of ulimits to be set in the container

	// Applicable to Windows
	CPUCount           int64  `json:"CpuCount"`   // CPU count
	CPUPercent         int64  `json:"CpuPercent"` // CPU percent
	IOMaximumIOps      uint64 // Maximum IOps for the container system drive
	IOMaximumBandwidth uint64 // Maximum IO in bytes per second for the container system drive
}

func (client *DockerClient) UpdateContainer(id string, updateConfig *UpdateConfig) error {
	data, err := json.Marshal(updateConfig)
	if err != nil {
		return err
	}
	uri := fmt.Sprintf("/%s/containers/%s/update", "v1.22", id)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	_, err = client.doRequest("POST", uri, data, headers)
	return err
}