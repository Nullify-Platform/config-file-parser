package models

type AttackSurface struct {
	// global only
	Enable               bool                       `yaml:"enable"`
	EnableDNSEnumeration bool                       `yaml:"enable_dns_enumeration"`
	AWSIntegration       *AWSIntegration            `yaml:"aws_integration,omitempty"`
	IPAddresses          []string                   `yaml:"ip_addresses,omitempty"`
	DomainNames          []string                   `yaml:"domain_names,omitempty"`
	IncludeOnly          []AttackSurfaceIncludeOnly `yaml:"include_only,omitempty"`
	Ignore               []AttackSurfaceIgnore      `yaml:"ignore,omitempty"`
}

type AWSIntegration struct {
	EnableAWSIntegration bool      `yaml:"enable_aws_integration"`
	PrimaryAccountID     string    `yaml:"primary_account_id"`
	PrimaryRegion        string    `yaml:"primary_region"`
	TargetRegions        *[]string `yaml:"target_regions,omitempty"`
	TargetAccounts       *[]string `yaml:"target_accounts,omitempty"`
}

type AttackSurfaceIncludeOnly struct {
	DomainNames []string                      `yaml:"domain_names,omitempty"`
	HTTP        *HTTPAttackSurfaceIncludeOnly `yaml:"http,omitempty"`
}

type HTTPAttackSurfaceIncludeOnly struct {
	Paths []string `yaml:"paths,omitempty"`
}

type AttackSurfaceIgnore struct {
	// empty fields are equivalent to *
	IPAddresses        []string                 `yaml:"ip_addresses,omitempty"`
	DomainNames        []string                 `yaml:"domain_names,omitempty"`
	TransportProtocols []string                 `yaml:"transport_protocols,omitempty"`
	Ports              []string                 `yaml:"ports,omitempty"`
	HTTP               *HTTPAttackSurfaceIgnore `yaml:"http,omitempty"`
}

type HTTPAttackSurfaceIgnore struct {
	Methods []string `yaml:"methods,omitempty"`
	Paths   []string `yaml:"paths,omitempty"`
}
