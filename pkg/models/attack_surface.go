package models

type AttackSurface struct {
	// global only
	Enable               bool                       `yaml:"enable"`
	EnableDNSEnumeration bool                       `yaml:"enable_dns_enumeration"`
	Hosts                []string                   `yaml:"hosts,omitempty"`
	IncludeOnly          []AttackSurfaceIncludeOnly `yaml:"include_only,omitempty"`
	Ignore               []AttackSurfaceIgnore      `yaml:"ignore,omitempty"`
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
