package models

type AttackSurface struct {
	// global only
	Enable               bool     `yaml:"enable"`
	EnableDNSEnumeration bool     `yaml:"enable_dns_enumeration"`
	DomainNames          []string `yaml:"domain_names,omitempty"`
	IgnoreDomainNames    []string `yaml:"ignore_domain_names,omitempty"`
}
