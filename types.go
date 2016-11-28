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
	Source                    string
	Application               string
	Description               string
	Environment               string
	SolutionStack             string
	VPCId                     string
	Subnets                   string
	ELBScheme                 string
	ApplicationHealthCheckURL string
	LoadBalancerType          string
	ServiceRole               string
	MinSize                   string
	MaxSize                   string
	InstanceType              string
	SecurityGroups            string
	IamInstanceProfile        string
	EC2KeyName                string
}
