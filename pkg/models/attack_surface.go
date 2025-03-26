package models

type AttackSurface struct {
	// global only
	Enable               bool                       `yaml:"enable"`
	EnableDNSEnumeration bool                       `yaml:"enable_dns_enumeration"`
	Hosts                []string                   `yaml:"hosts,omitempty"`
	IncludeOnly          []AttackSurfaceScopingRule `yaml:"include_only,omitempty"`
	Ignore               []AttackSurfaceScopingRule `yaml:"ignore,omitempty"`
}

type AttackSurfaceScopingRule struct {
	Hosts              []string                          `yaml:"hosts,omitempty"`
	TransportProtocols []string                          `yaml:"transport_protocols,omitempty"`
	Ports              []string                          `yaml:"ports,omitempty"`
	HTTP               *HTTPAttackSurfaceScopingRuleHTTP `yaml:"http,omitempty"`
}

type HTTPAttackSurfaceScopingRuleHTTP struct {
	Methods []string `yaml:"methods,omitempty"`
	Paths   []string `yaml:"paths,omitempty"`
}
