// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2017-2020 Datadog, Inc.

// +build windows

package containers

import (
	"github.com/DataDog/datadog-agent/pkg/util/containers/metrics"
)

// SetCgroups has to be called when creating the Container, in order to
// be able to enable FillCgroupMetrics and FillNetworkMetrics to works
func (c *Container) SetCgroups(cgroup *metrics.ContainerCgroup) error {
	return nil
}

// FillCgroupLimits fills the resource limits for a Container, based on the
// associated cgroups. This can be called once if the limits are assumed constant.
func (c *Container) FillCgroupLimits() error {
	return nil
}

// FillCgroupMetrics fills the performance metrics for a Container, based on the
// associated cgroups. Network metrics are handled by FillNetworkMetrics
func (c *Container) FillCgroupMetrics() error {
	return nil
}

// FillNetworkMetrics fills the network metrics for a Container,
// based on the associated cgroups.
func (c *Container) FillNetworkMetrics(networks map[string]string) error {
	return nil
}
