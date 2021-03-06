// Copyright 2019 NetApp, Inc. All Rights Reserved.
package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/version"

	"github.com/netapp/trident/frontend/csi"
)

func TestSupportsFeature(t *testing.T) {

	var supportedTests = []struct {
		versionInfo version.Info
		expected    bool
	}{
		{version.Info{GitVersion: "v1.14.3"}, false},
		{version.Info{GitVersion: "v1.16.2"}, true},
		{version.Info{GitVersion: "garbage"}, false},
	}

	for _, tc := range supportedTests {
		plugin := Plugin{kubeVersion: &tc.versionInfo}
		supported := plugin.SupportsFeature(csi.ExpandCSIVolumes)
		if tc.expected {
			print("bluh\n")
			assert.True(t, supported, "Expected true")
		} else {
			print("bar\n")
			assert.False(t, supported, "Expected false")
		}
	}
}
