package models

type GetCMDClassSchema struct {
	Result struct {
		Name       string `json:"name"`
		Attributes []struct {
			IsMandatory  string `json:"is_mandatory"`
			DefaultValue string `json:"default_value"`
			Label        string `json:"label"`
			Type         string `json:"type"`
			IsDisplay    string `json:"is_display"`
			Element      string `json:"element"`
		} `json:"attributes"`
	} `json:"result"`
}

type GetTableItem struct {
	Result []map[string]interface{} `json:"result"`
}

/*type LinkValue struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}
type TableIface interface {
	GetSysClassName() string
}

func (t *GetTableWin) GetSysClassName() string {
	return t.SysClassName
}

func (t *GetTableLin) GetSysClassName() string {
	return t.SysClassName
}

type GetTableWin struct {
	SysClassName     string    `json:"sys_class_name"`
	Name             string    `json:"name,omitempty"`
	CPUManufacturer  string    `json:"cpu_manufacturer"`
	SysUpdatedOn     string    `json:"sys_updated_on"`
	SysCreatedBy     string    `json:"sys_created_by"`
	RAM              string    `json:"ram"`
	CPUSpeed         string    `json:"cpu_speed"`
	OwnedBy          string    `json:"owned_by"`
	DiskSpace        string    `json:"disk_space"`
	ObjectID         string    `json:"object_id"`
	DNSDomain        string    `json:"dns_domain"`
	ShortDescription string    `json:"short_description"`
	OsDomain         string    `json:"os_domain"`
	CPUCount         string    `json:"cpu_count"`
	Manufacturer     LinkValue `json:"manufacturer"`
	StartDate        string    `json:"start_date"`
	SerialNumber     string    `json:"serial_number"`
	Asset            LinkValue `json:"asset"`
	CPUCoreCount     string    `json:"cpu_core_count"`
	SysUpdatedBy     string    `json:"sys_updated_by"`
	SysCreatedOn     string    `json:"sys_created_on"`
	CPUType          string    `json:"cpu_type"`
	AssetTag         string    `json:"asset_tag"`
	ChassisType      string    `json:"chassis_type"`
	// Probably Unique Identifier.
	SysID     string    `json:"sys_id"`
	Company   LinkValue `json:"company"`
	Os        string    `json:"os"`
	IPAddress string    `json:"ip_address"`
	ModelID   LinkValue `json:"model_id"`
	Category  string    `json:"category"`
	HostName  string    `json:"host_name"`
	LeaseID   string    `json:"lease_id"`
}
*/
type GetTableLin struct {
	FirewallStatus    string `json:"firewall_status"`
	OsAddressWidth    string `json:"os_address_width"`
	AttestedDate      string `json:"attested_date"`
	OperationalStatus string `json:"operational_status"`

	OsServicePack string `json:"os_service_pack"`

	CPUCoreThread      string `json:"cpu_core_thread"`
	CPUManufacturer    string `json:"cpu_manufacturer"`
	SysUpdatedOn       string `json:"sys_updated_on"`
	DiscoverySource    string `json:"discovery_source"`
	FirstDiscovered    string `json:"first_discovered"`
	DueIn              string `json:"due_in"`
	UsedFor            string `json:"used_for"`
	GlAccount          string `json:"gl_account"`
	InvoiceNumber      string `json:"invoice_number"`
	SysCreatedBy       string `json:"sys_created_by"`
	RAM                string `json:"ram"`
	WarrantyExpiration string `json:"warranty_expiration"`
	CPUName            string `json:"cpu_name"`
	CPUSpeed           string `json:"cpu_speed"`
	OwnedBy            string `json:"owned_by"`
	CheckedOut         string `json:"checked_out"`
	//
	KernelRelease string `json:"kernel_release"`
	//
	Classification      string `json:"classification"`
	DiskSpace           string `json:"disk_space"`
	SysDomainPath       string `json:"sys_domain_path"`
	BusinessUnit        string `json:"business_unit"`
	ObjectID            string `json:"object_id"`
	MaintenanceSchedule string `json:"maintenance_schedule"`
	CostCenter          string `json:"cost_center"`
	AttestedBy          string `json:"attested_by"`
	DNSDomain           string `json:"dns_domain"`
	Assigned            string `json:"assigned"`
	LifeCycleStage      string `json:"life_cycle_stage"`
	PurchaseDate        string `json:"purchase_date"`
	CdSpeed             string `json:"cd_speed"`
	ShortDescription    string `json:"short_description"`
	Floppy              string `json:"floppy"`
	ManagedBy           struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"managed_by"`
	OsDomain       string `json:"os_domain"`
	CanPrint       string `json:"can_print"`
	LastDiscovered string `json:"last_discovered"`
	SysClassName   string `json:"sys_class_name"`
	CPUCount       string `json:"cpu_count"`
	Manufacturer   struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"manufacturer"`
	LifeCycleStageStatus string `json:"life_cycle_stage_status"`
	Vendor               struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"vendor"`
	ModelNumber   string `json:"model_number"`
	AssignedTo    string `json:"assigned_to"`
	StartDate     string `json:"start_date"`
	OsVersion     string `json:"os_version"`
	SerialNumber  string `json:"serial_number"`
	CdRom         string `json:"cd_rom"`
	SupportGroup  string `json:"support_group"`
	CorrelationID string `json:"correlation_id"`
	Unverified    string `json:"unverified"`
	Attributes    string `json:"attributes"`
	Asset         struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"asset"`
	CPUCoreCount     string `json:"cpu_core_count"`
	FormFactor       string `json:"form_factor"`
	SkipSync         string `json:"skip_sync"`
	MostFrequentUser string `json:"most_frequent_user"`
	AttestationScore string `json:"attestation_score"`
	SysUpdatedBy     string `json:"sys_updated_by"`
	SysCreatedOn     string `json:"sys_created_on"`
	CPUType          string `json:"cpu_type"`
	SysDomain        struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"sys_domain"`
	InstallDate       string `json:"install_date"`
	AssetTag          string `json:"asset_tag"`
	DrBackup          string `json:"dr_backup"`
	HardwareSubstatus string `json:"hardware_substatus"`
	Fqdn              string `json:"fqdn"`
	ChangeControl     string `json:"change_control"`
	InternetFacing    string `json:"internet_facing"`
	DeliveryDate      string `json:"delivery_date"`
	HardwareStatus    string `json:"hardware_status"`
	InstallStatus     string `json:"install_status"`
	SupportedBy       string `json:"supported_by"`
	Name              string `json:"name"`
	Subcategory       string `json:"subcategory"`
	DefaultGateway    string `json:"default_gateway"`
	ChassisType       string `json:"chassis_type"`
	Virtual           string `json:"virtual"`
	AssignmentGroup   string `json:"assignment_group"`
	ManagedByGroup    string `json:"managed_by_group"`
	SysID             string `json:"sys_id"`
	PoNumber          string `json:"po_number"`
	CheckedIn         string `json:"checked_in"`
	SysClassPath      string `json:"sys_class_path"`
	MacAddress        string `json:"mac_address"`
	Company           struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"company"`
	Justification string `json:"justification"`
	Department    string `json:"department"`
	Comments      string `json:"comments"`
	Cost          string `json:"cost"`
	Os            string `json:"os"`
	SysModCount   string `json:"sys_mod_count"`
	Monitor       string `json:"monitor"`
	IPAddress     string `json:"ip_address"`
	ModelID       struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"model_id"`
	DuplicateOf string `json:"duplicate_of"`
	SysTags     string `json:"sys_tags"`
	CostCc      string `json:"cost_cc"`
	OrderDate   string `json:"order_date"`
	Schedule    string `json:"schedule"`
	Environment string `json:"environment"`
	Due         string `json:"due"`
	Attested    string `json:"attested"`
	Location    struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"location"`
	Category   string `json:"category"`
	FaultCount string `json:"fault_count"`
	HostName   string `json:"host_name"`
	LeaseID    string `json:"lease_id"`
}
