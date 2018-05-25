package main

import (
	"fmt"

	"github.com/digitalocean/go-openvswitch/ovs"
)

// OVSSwitch is a bridge instance
type OVSSwitch struct {
	bridgeName string
	ovsclient  *ovs.Client
}

// NewOVSSwitch for creating a ovs bridge
func NewOVSSwitch(bridgeName string) (*OVSSwitch, error) {
	c := ovs.New(ovs.Sudo())
	if err := c.VSwitch.AddBridge(bridgeName); err != nil {
		return nil, fmt.Errorf("failed to add bridge: %v", err)
	}
	return &OVSSwitch{
		bridgeName: bridgeName,
		ovsclient:  c,
	}, nil
}

// ovs-vsctl add-port br0 eth0
func (sw *OVSSwitch) addPort(ifName string) error {
	if err := sw.ovsclient.VSwitch.AddPort(sw.bridgeName, ifName); err != nil {
		return fmt.Errorf("failed to add port: %v", err)
	}
	return nil
}

