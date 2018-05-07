package k8splugin
//this file is referenced  file /contiv1.2.0/drivers/endpointstate.go

import (
	"encoding/json"
	"fmt"

	"github.com/contiv/netplugin/core"
	//"github.com/contiv/netplugin/netmaster/mastercfg"
)

const (
	// StateBasePath is the base path for all state operations.
	StateBasePath = "/contiv.io/"
	// StateConfigPath is the path to the root of the configuration state
	StateConfigPath = StateBasePath + "state/"
	// StateOperPath is the path for operational/runtime state
	StateOperPath = StateBasePath + "oper/"

	networkConfigPathPrefix  = StateConfigPath + "nets/"
	networkConfigPath        = networkConfigPathPrefix + "%s"
	endpointConfigPathPrefix = StateConfigPath + "eps/"
	endpointConfigPath       = endpointConfigPathPrefix + "%s"
	epGroupConfigPathPrefix  = StateConfigPath + "endpointGroups/"
	epGroupConfigPath        = epGroupConfigPathPrefix + "%s"
)

type cniEndpointState struct {
	core.CommonState
	NetID            string            `json:"netID"`
	EndpointID       string            `json:"endpointID"`
	ServiceName      string            `json:"serviceName"`
	EndpointGroupID  int               `json:"endpointGroupId"`

	/*
		resp.Tenant = tenant
		resp.Network = netw
		resp.Group = epg
		resp.EndpointID = pInfo.InfraContainerID
		resp.Name = pInfo.Name

		come from function getEPSpec()
		maybe include (type epSpec struct)
		get tenant id by split str
	*/
}

// Read the state for a given identifier.
// Don't use StateOperPath,use StateConfigPath instead
func (s *cniEndpointState) Read(id string) error {
	key := fmt.Sprintf(endpointConfigPath, id)
	return s.StateDriver.ReadState(key, s, json.Unmarshal)
}

// Write the state.
func (s *cniEndpointState) Write() error {
	key := fmt.Sprintf(endpointConfigPath, s.ID)
	return s.StateDriver.WriteState(key, s, json.Marshal)
}

// ReadAll reads all state objects for the endpoints.
func (s *cniEndpointState) ReadAll() ([]core.State, error) {
	return s.StateDriver.ReadAllState(endpointConfigPathPrefix, s, json.Unmarshal)
}

// WatchAll fills a channel on each state event related to endpoints.
func (s *cniEndpointState) WatchAll(rsps chan core.WatchState) error {
	return s.StateDriver.WatchAllState(endpointConfigPathPrefix, s, json.Unmarshal,
		rsps)
}


// Clear removes the state.
func (s *cniEndpointState) Clear() error {
	key := fmt.Sprintf(endpointConfigPath, s.ID)
	return s.StateDriver.ClearState(key)
}
