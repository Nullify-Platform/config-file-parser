package models

type AttackSurface struct {
	// global only
	EnableDNSTraversal bool     `yaml:"enable_dns_traversal,omitempty"`
	DomainNames        []string `yaml:"domain_names,omitempty"`
	IgnoreDomainNames  []string `yaml:"ignore_domain_names,omitempty"`
}
