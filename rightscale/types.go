package rightscale

import (
	"time"
)

type Resource struct {
	Namespace string      `json:"namespace"`
	Type      string      `json:"type"`
	Fields    interface{} `json:"fields"`
}

//type Server struct {
//	Name                     string                   `json:"name" required:"true"`
//	Description              string                   `json:"description,omitempty" required:"false"`
//	Cloud                    string                   `json:"cloud_href" required:"true"`
//	Datacenter               string                   `json:"datacenter_href,omitempty" required:"false"`
//	Image                    string                   `json:"image_href" required:"true"`
//	InstanceType             string                   `json:"instance_type_href,omitempty" required:"false"`
//	KernelImage              string                   `json:"kernel_image_href,omitempty" required:"false"`
//	Optimized                string                   `json:"optimized,omitempty" required:"false"`
//	RamdiskImage             string                   `json:"ramdisk_image_href,omitempty" required:"false"`
//	SecurityGroups           []string                 `json:"security_group_hrefs,omitempty" required:"false"`
//	SSHKey                   string                   `json:"ssh_key_href,omitempty" required:"false"`
//	Subnets                  []string                 `json:"subnet_hrefs,omitempty" required:"false"`
//	UserData                 string                   `json:"user_data,omitempty" required:"false"`
//	VolumeAttachments        []string                 `json:"volume_attachment_hrefs,omitempty" required:"false"`
//	CloudSpecificAttributes  *CloudSpecificAttributes `json:"cloud_specific_attributes,omitempty" required:"false"`
//	AssociatePublicIPAddress string                   `json:"associate_public_ip_address,omitempty" required:"false"`
//	IPForwardingEnabled      string                   `json:"ip_forwarding_enabled,omitempty" required:"false"`
//	PlacementGroup           string                   `json:"placement_group_href,omitempty" required:"false"`
//	Inputs                   string                   `json:"inputs,omitempty" required:"false"`
//	MultiCloudImage          string                   `json:"multi_cloud_image_href,omitempty" required:"false"`
//	ServerTemplate           string                   `json:"server_template_href" required:"true"`
//}
//
//type Instance struct {
//	Name                     string                   `json:"name" required:"true"`
//	Description              string                   `json:"description,omitempty" required:"false"`
//	Cloud                    string                   `json:"cloud_href" required:"true"`
//	Datacenter               string                   `json:"datacenter_href,omitempty" required:"false"`
//	Image                    string                   `json:"image_href" required:"true"`
//	InstanceType             string                   `json:"instance_type_href,omitempty" required:"false"`
//	KernelImage              string                   `json:"kernel_image_href,omitempty" required:"false"`
//	Optimized                string                   `json:"optimized,omitempty" required:"false"`
//	RamdiskImage             string                   `json:"ramdisk_image_href,omitempty" required:"false"`
//	SecurityGroups           []string                 `json:"security_group_hrefs,omitempty" required:"false"`
//	SSHKey                   string                   `json:"ssh_key_href,omitempty" required:"false"`
//	Subnets                  []string                 `json:"subnet_hrefs,omitempty" required:"false"`
//	UserData                 string                   `json:"user_data,omitempty" required:"false"`
//	VolumeAttachments        []string                 `json:"volume_attachment_hrefs,omitempty" required:"false"`
//	CloudSpecificAttributes  *CloudSpecificAttributes `json:"cloud_specific_attributes,omitempty" required:"false"`
//	AssociatePublicIPAddress string                   `json:"associate_public_ip_address,omitempty" required:"false"`
//	IPForwardingEnabled      string                   `json:"ip_forwarding_enabled,omitempty" required:"false"`
//	PlacementGroup           string                   `json:"placement_group_href,omitempty" required:"false"`
//}
//
//type CloudSpecificAttributes struct {
//	AdminUsername                    string `json:"admin_username,omitempty" required:"false"`
//	AutomaticInstanceStoreMapping    string `json:"automatic_instance_store_mapping,omitempty" required:"false"`
//	AvailabilitySet                  string `json:"availability_set,omitempty" required:"false"`
//	CreateBootVolume                 string `json:"create_boot_volume,omitempty" required:"false"`
//	CreateDefaultPortForwardingRules string `json:"create_default_port_forwarding_rules,omitempty" required:"false"`
//	DeleteBootVolume                 string `json:"delete_boot_volume,omitempty" required:"false"`
//	DiskGB                           int    `json:"disk_gb,omitempty" required:"false"`
//	EBSOptimized                     string `json:"ebs_optimized,omitempty" required:"false"`
//	IAMInstanceProfile               string `json:"iam_instance_profile,omitempty" required:"false"`
//	KeepAliveID                      string `json:"keep_alive_id,omitempty" required:"false"`
//	KeepAliveURL                     string `json:"keep_alive_url,omitempty" required:"false"`
//	LocalSSDCount                    string `json:"local_ssd_count,omitempty" required:"false"`
//	LocalSSDInterface                string `json:"local_ssd_interface,omitempty" required:"false"`
//	MaxSpotPrice                     string `json:"max_spot_price,omitempty" required:"false"`
//	MemoryMB                         int    `json:"memory_mb,omitempty" required:"false"`
//	Metadata                         string `json:"metadata,omitempty" required:"false"`
//	NumCores                         int    `json:"num_cores,omitempty" required:"false"`
//	PlacementTenancy                 string `json:"placement_tenancy,omitempty" required:"false"`
//	Preemptible                      string `json:"preemptible,omitempty" required:"false"`
//	PricingType                      string `json:"pricing_type,omitempty" required:"false"`
//	RootVolumePerformance            string `json:"root_volume_performance,omitempty" required:"false"`
//	RootVolumeSize                   string `json:"root_volume_size,omitempty" required:"false"`
//	RootVolumeTypeUID                string `json:"root_volume_type_uid,omitempty" required:"false"`
//	ServiceAccount                   string `json:"service_account,omitempty" required:"false"`
//}

type Deployment struct {
	Name        string `json:"name" required:"true"`
	Description string `json:"description,omitempty" required:"false"`
}

type ProviderConfiguration struct {
	client        *RsClient
	accountNumber int
	apiHostname   string
}

type ProcessMedia struct {
	Id           string                       `json:"id"`
	Href         string                       `json:"href"`
	Name         string                       `json:"name"`
	Source       string                       `json:"source"`
	Main         string                       `json:"main"`
	Application  string                       `json:"application"`
	Status       string                       `json:"status"`
	Outputs      []OutputMedia                `json:"outputs"`
	Tasks        []TaskMedia                  `json:"tasks"`
	CreatedAt    time.Time                    `json:"created_at"`
	UpdatedAt    time.Time                    `json:"updated_at"`
	FinishedAt   time.Time                    `json:"finished_at"`
	FinishReason string                       `json:"finish_reason"`
	Links        map[string]map[string]string `json:"links"`
}

type OutputMedia struct {
	Name  string         `json:"name"`
	Value ParameterMedia `json:"value"`
}

type TaskMedia struct {
	Id         string                       `json:"id"`
	Href       string                       `json:"href"`
	Name       string                       `json:"name"`
	Label      string                       `json:"label"`
	Process    *ProcessMedia                `json:"process"`
	Sequence   int                          `json:"sequence"`
	Callstack  []string                     `json:"callstack"`
	Status     string                       `json:"status"`
	Error      *ErrorMedia                  `json:"error"`
	CreatedAt  time.Time                    `json:"created_at"`
	UpdatedAt  time.Time                    `json:"updated_at"`
	FinishedAt time.Time                    `json:"finished_at"`
	Links      map[string]map[string]string `json:"links"`
}

type ParameterMedia struct {
	Kind  string      `json:"kind"`
	Value interface{} `json:"value"`
}

type ErrorMedia struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}
