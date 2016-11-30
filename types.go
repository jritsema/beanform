package main

// TfState represents a .tfstate file
type TfState struct {
	Version          int    `json:"version"`
	TerraformVersion string `json:"terraform_version"`
	Serial           int    `json:"serial"`
	Lineage          string `json:"lineage"`
	Modules          []struct {
		Path    []string `json:"path"`
		Outputs struct {
		} `json:"outputs"`
		Resources map[string]BeanstalkEnvironment `json:"resources"`
		DependsOn []interface{}                   `json:"depends_on"`
	} `json:"modules"`
}

// BeanstalkEnvironment represents a BeanstalkEnvironment
type BeanstalkEnvironment struct {
	Type      string                      `json:"type"`
	DependsOn []interface{}               `json:"depends_on"`
	Primary   PrimaryBeanstalkEnvironment `json:"primary"`
	Deposed   []interface{}               `json:"deposed"`
	Provider  string                      `json:"provider"`
}

// PrimaryBeanstalkEnvironment represents a primary beanstalk environment
type PrimaryBeanstalkEnvironment struct {
	ID         string            `json:"id"`
	Attributes map[string]string `json:"attributes"`
	Meta       struct {
		SchemaVersion string `json:"schema_version"`
	} `json:"meta"`
	Tainted bool `json:"tainted"`
}

// BeanstalkHCL is an in-memory representation of a Beanstalk app in HCL
type BeanstalkHCL struct {
	Source                                   string
	Application                              string
	Description                              string
	Environment                              string
	SolutionStack                            string
	undefined                                string
	EvaluationPeriods                        string
	ListenerProtocol                         string
	LoadBalancerPortProtocol                 string
	AvailabilityZones                        string
	RollingUpdateEnabled                     string
	AppSource                                string
	AssociatePublicIPAddress                 string
	IgnoreHealthCheck                        string
	CrossZone                                string
	RootVolumeType                           string
	ListenerEnabled                          string
	CustomAvailabilityZones                  string
	BatchSizeType                            string
	LoadBalancerSSLPortProtocol              string
	NotificationTopicARN                     string
	LowerThreshold                           string
	SSLCertificateIDListener                 string
	BlockDeviceMappings                      string
	ELBScheme                                string
	UnhealthyThreshold                       string
	InstanceProtocol                         string
	BatchSize                                string
	MinInstancesInService                    string
	LaunchType                               string
	Interval                                 string
	IamInstanceProfile                       string
	Period                                   string
	UpperBreachScaleIncrement                string
	LogPublicationControl                    string
	LoadBalancerType                         string
	InstancePort                             string
	InstancePortCloudFormation               string
	Unit                                     string
	Enable32bitApplications                  string
	RollingUpdateType                        string
	ApplicationHealthcheckURL                string
	ConnectionSettingIdleTimeout             string
	PolicyNames                              string
	RootVolumeSize                           string
	EnvironmentType                          string
	ConnectionDrainingTimeout                string
	NotificationEndpoint                     string
	ELBSubnets                               string
	ExternalExtensionsS3Key                  string
	EnvironmentVariables                     string
	RootVolumeIOPS                           string
	MeasureName                              string
	Timeout                                  string
	RollingUpdateTimeout                     string
	CommandTimeout                           string
	HealthCheckTimeout                       string
	HealthCheckSuccessThreshold              string
	Statistic                                string
	MonitoringInterval                       string
	DeploymentPolicy                         string
	InstanceType                             string
	NotificationTopicName                    string
	ExternalExtensionsS3Bucket               string
	Cooldown                                 string
	SecurityGroups                           string
	SecurityGroupsELB                        string
	ServiceRole                              string
	NotificationProtocol                     string
	LowerBreachScaleIncrement                string
	ImageID                                  string
	UpperThreshold                           string
	SSLCertificateID                         string
	AutomaticallyTerminateUnhealthyInstances string
	MinSize                                  string
	PauseTime                                string
	VPCId                                    string
	ConnectionDrainingEnabled                string
	BreachDuration                           string
	Target                                   string
	LaunchTimeout                            string
	EC2KeyName                               string
	DefaultSSHPort                           string
	MaxBatchSize                             string
	SSHSourceRestriction                     string
	HealthyThreshold                         string
	TargetRuntime                            string
	MaxSize                                  string
	LoadBalancerHTTPSPort                    string
	LoadBalancerHTTPPort                     string
	Subnets                                  string
	RollbackLaunchOnFailure                  string
}
