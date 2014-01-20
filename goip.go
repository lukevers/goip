package goip

import (
	"net"
)

// HasNetmask returns true if the address specified matches the
// netmask specified. The address specified is a net.IP in this
// case. Use HasNetmaskString if the address is in string
// form. HasNetmask returns false if the address specified does not
// match the netmask specified.
func HasNetmask(netmask *net.IPNet, address net.IP) bool {
	return netmask.Contains(address)
}

// HasNetmaskString returns true if the address specified matches the
// netmask specified. The address specified is a string in this
// case. Use HasNetmask if the address is already in net.IP
// form. HasNetmaskString returns false if the address specified does
// not match the netmask specified.
func HasNetmaskString(netmask *net.IPNet, address string) bool {
	return netmask.Contains(net.ParseIP(address))
}
