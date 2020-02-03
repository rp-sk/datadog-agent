// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2020 Datadog, Inc.

// +build windows

package metrics

// ScrapeAllCgroups returns ContainerCgroup for every container that's in a Cgroup.
// This version iterates on /{host/}proc to retrieve processes out of the namespace.
// We return as a map[containerID]Cgroup for easy look-up.
func ScrapeAllCgroups() (map[string]*ContainerCgroup, error) {
	return map[string]*ContainerCgroup{}, nil
}

// DetectNetworkDestinations lists all the networks available
// to a given PID and parses them in NetworkInterface objects
func DetectNetworkDestinations(pid int) ([]NetworkDestination, error) {
	return make([]NetworkDestination, 0), nil
}
