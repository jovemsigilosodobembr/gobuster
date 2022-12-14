package gobustervhost

import (
	"github.com/OJ/gobuster/v3/libgobuster"
)

// OptionsVhost is the struct to hold all options for this plugin
type OptionsVhost struct {
	libgobuster.HTTPOptions
	AppendDomain  bool
	ExcludeLength []int
	Domain        string
}

// NewOptionsVhost returns a new initialized OptionsVhost
func NewOptionsVhost() *OptionsVhost {
	return &OptionsVhost{}
}
