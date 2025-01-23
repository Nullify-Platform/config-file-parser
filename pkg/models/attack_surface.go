package models

type AttackSurface struct {
	// global only
	Enable             bool     `yaml:"enable,omitempty"`
	EnableDNSTraversal bool     `yaml:"enable_dns_traversal,omitempty"`
	DomainNames        []string `yaml:"domain_names,omitempty"`
	IgnoreDomainNames  []string `yaml:"ignore_domain_names,omitempty"`
}
