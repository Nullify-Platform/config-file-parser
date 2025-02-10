package models

type AttackSurface struct {
	// global only
	Enable               bool     `yaml:"enable"`
	EnableDNSEnumeration bool     `yaml:"enable_dns_enumeration"`
	DomainNames          []string `yaml:"domain_names,omitempty"`
	IgnoreDomainNames    []string `yaml:"ignore_domain_names,omitempty"`
	PathPrefixes         []string `yaml:"path_prefixes,omitempty"`
	IgnoreMethods        []string `yaml:"ignore_methods,omitempty"`
	IgnorePorts          []int    `yaml:"ignore_ports,omitempty"`
	Schemes              []string `yaml:"schemes"`
}
