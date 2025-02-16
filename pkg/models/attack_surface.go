package models

type AttackSurface struct {
	// global only
	Enable               bool                       `yaml:"enable"`
	EnableDNSEnumeration bool                       `yaml:"enable_dns_enumeration"`
	AWSIntegration       *AWSIntegration            `yaml:"aws_integration"`
	Hosts                []string                   `yaml:"hosts,omitempty"`
	IncludeOnly          []AttackSurfaceIncludeOnly `yaml:"include_only,omitempty"`
	Ignore               []AttackSurfaceIgnore      `yaml:"ignore,omitempty"`
}

type AWSIntegration struct {
	Enable           bool      `yaml:"enable"`
	PrimaryAccountID string    `yaml:"primary_account_id,omitempty"`
	PrimaryRegion    string    `yaml:"primary_region,omitempty"`
	TargetRegions    *[]string `yaml:"target_regions,omitempty"`
	TargetAccounts   *[]string `yaml:"target_accounts,omitempty"`
}

type AttackSurfaceIncludeOnly struct {
	Hosts []string                      `yaml:"hosts,omitempty"`
	HTTP  *HTTPAttackSurfaceIncludeOnly `yaml:"http,omitempty"`
}

type HTTPAttackSurfaceIncludeOnly struct {
	Methods []string `yaml:"methods,omitempty"`
	Paths   []string `yaml:"paths,omitempty"`
}

type AttackSurfaceIgnore struct {
	// empty fields are equivalent to *
	Hosts              []string                 `yaml:"hosts,omitempty"`
	TransportProtocols []string                 `yaml:"transport_protocols,omitempty"`
	Ports              []string                 `yaml:"ports,omitempty"`
	HTTP               *HTTPAttackSurfaceIgnore `yaml:"http,omitempty"`
}

type HTTPAttackSurfaceIgnore struct {
	Methods []string `yaml:"methods,omitempty"`
	Paths   []string `yaml:"paths,omitempty"`
}
