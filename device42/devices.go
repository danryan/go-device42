package device42

import "time"

// DevicesService handles communication with the device-related
// methods of the Device42 API.
type DevicesService service

type Device struct {
	ID                      int                      `json:"id,omitempty"`
	DeviceID                int                      `json:"device_id,omitempty"`
	Aliases                 []string                 `json:"aliases,omitempty"`
	AssetNo                 string                   `json:"asset_no,omitempty"`
	Building                string                   `json:"building,omitempty"`
	Category                string                   `json:"category,omitempty"`
	CPUCore                 int                      `json:"cpucore,omitempty"`
	CPUCount                int                      `json:"cpucount,omitempty"`
	CPUSpeed                float32                  `json:"cpuspeed,omitempty"`
	Customer                string                   `json:"customer,omitempty"`
	CustomerID              int                      `json:"customer_id,omitempty"`
	CustomFields            CustomFields             `json:"custom_fields,omitempty"`
	DeviceExternalLinks     []DeviceExternalLink     `json:"device_external_links,omitempty"`
	DevicePurchaseLineItems []DevicePurchaseLineItem `json:"device_purchase_line_items,omitempty"`
	Groups                  string                   `json:"groups,omitempty"`
	HardDiskDetails         []HardDiskDetail         `json:"hdd_details,omitempty"`
	HardDiskCount           int                      `json:"hdcount,omitempty"`
	HardDiskRaid            string                   `json:"hddraid,omitempty"`
	HardDiskRaidType        string                   `json:"hddraid_type,omitempty"`
	HardDiskSize            string                   `json:"hddsize,omitempty"`
	HWDepth                 int                      `json:"hw_depth,omitempty"`
	HWModel                 string                   `json:"hw_model,omitempty"`
	HWSize                  float32                  `json:"hw_size,omitempty"`
	InService               bool                     `json:"in_service,omitempty"`
	IPAddresses             []DeviceIPAddress        `json:"ip_addresses,omitempty"`
	LastUpdated             time.Time                `json:"last_updated,omitempty"`
	MACAddresses            []DeviceMacAddress       `json:"mac_addresses,omitempty"`
	Manufacturer            string                   `json:"manufacturer,omitempty"`
	Name                    string                   `json:"name,omitempty"`
	Notes                   string                   `json:"notes,omitempty"`
	Orientation             int                      `json:"orientation,omitempty"`
	OS                      string                   `json:"os,omitempty"`
	OSArch                  int                      `json:"osarch,omitempty"`
	OSVer                   string                   `json:"osver,omitempty"`
	Rack                    string                   `json:"rack,omitempty"`
	RackID                  int                      `json:"rack_id,omitempty"`
	RAM                     float32                  `json:"ram,omitempty"`
	Room                    string                   `json:"room,omitempty"`
	Row                     string                   `json:"row,omitempty"`
	SerialNo                string                   `json:"serial_no,omitempty"`
	ServiceLevel            string                   `json:"service_level,omitempty"`
	StartAt                 float32                  `json:"start_at,omitempty"`
	Tags                    []string                 `json:"tags,omitempty"`
	Type                    string                   `json:"type,omitempty"`
	UCSManager              string                   `json:"ucs_manager,omitempty"`
	UUID                    string                   `json:"uuid,omitempty"`
	VirtualHostName         string                   `json:"virtual_host_name,omitempty"`
	Where                   int                      `json:"where,omitempty"`
	XPos                    int                      `json:"xpos,omitempty"`
}

type DeviceExternalLink map[string]string

type DevicePurchaseLineItem map[string]interface{}

type HardDiskDetail struct {
	Description string `json:"description,omitempty"`
	HardDisk    struct {
		Bytes          string `json:"bytes,omitempty"`
		Description    string `json:"description,omitempty"`
		HDID           int    `json:"hd_id,omitempty"`
		Location       string `json:"location,omitempty"`
		ManufacturerID int    `json:"manufacturer_id,omitempty"`
		Notes          string `json:"notes,omitempty"`
		PartNo         string `json:"partno,omitempty"`
		RPM            struct {
			ID   int    `json:"description,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"rpm,omitempty"`
		Size float32 `json:"size,omitempty"`
		Type struct {
			ID   int    `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"type,omitempty"`
		HardDiskCount int    `json:"hddcount,omitempty"`
		RaidGroup     string `json:"raid_group,omitempty"`
		RaidType      string `json:"raid_type,omitempty"`
	} `json:"hdd,omitempty"`
}

type DeviceIPAddress struct {
	IP         string `json:"ip,omitempty"`
	Label      string `json:"label,omitempty"`
	MACAddress string `json:"macaddress,omitempty"`
	Subnet     string `json:"subnet,omitempty"`
	SubnetID   int    `json:"subnet_id,omitempty"`
	Type       int    `json:"type,omitempty"`
}

type DeviceMacAddress struct {
	MAC      string `json:"mac,omitempty"`
	Port     string `json:"port,omitempty"`
	PortName string `json:"port_name,omitempty"`
	VLAN     int    `json:"vlan,omitempty"`
}
