package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/harvester/harvester-installer/pkg/util"
)

func TestCalcCosPersistentPartSize(t *testing.T) {
	testCases := []struct {
		diskSize       uint64
		persistentSize string
		output         uint64
		err            string
	}{
		{
			diskSize:       10,
			persistentSize: "",
			output:         0,
			err:            "disk too small: 10GB. Minimum 60GB is required",
		},
		{
			diskSize:       120,
			persistentSize: "",
			output:         43,
		},
		{
			diskSize:       140,
			persistentSize: "20%",
			output:         28,
		},
		{
			diskSize:       140,
			persistentSize: "5%",
			output:         0,
			err:            "partition size cannot be smaller than 25Gi",
		},
	}

	for _, testCase := range testCases {
		persistentSize, err := calcCosPersistentPartSize(testCase.diskSize, testCase.persistentSize)
		assert.Equal(t, testCase.output, persistentSize)
		if err != nil {
			assert.EqualError(t, err, testCase.err)
		}
	}
}

func TestConvertToCos_SSHKeysInYipNetworkStage(t *testing.T) {
	conf, err := LoadHarvesterConfig(util.LoadFixture(t, "harvester-config.yaml"))
	assert.NoError(t, err)

	yipConfig, err := ConvertToCOS(conf)
	assert.NoError(t, err)

	assert.Equal(t, yipConfig.Stages["network"][0].SSHKeys["rancher"], conf.OS.SSHAuthorizedKeys)
	assert.Nil(t, yipConfig.Stages["initramfs"][0].SSHKeys)
}
