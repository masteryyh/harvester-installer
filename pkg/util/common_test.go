package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePartitionSize(t *testing.T) {
	testCases := []struct {
		diskSizeBytes uint64
		partitionSize string
		output        uint64
		err           string
	}{
		{
			diskSizeBytes: 214748364800,
			partitionSize: "",
			output:        0,
			err:           "partition size cannot be empty",
		},
		{
			diskSizeBytes: 214748364800,
			partitionSize: "20%",
			output:        42949672960,
		},
		{
			diskSizeBytes: 1099511627776,
			partitionSize: "10%",
			output:        109951162777,
		},
		{
			diskSizeBytes: 214748364800,
			partitionSize: "0%",
			output:        0,
			err:           "percentage cannot below 0 or bigger than 100%",
		},
		{
			diskSizeBytes: 214748364800,
			partitionSize: "lol%",
			output:        0,
			err:           "invalid partition size",
		},
		{
			diskSizeBytes: 150323855360,
			partitionSize: "1%",
			output:        0,
			err:           "partition size cannot be smaller than 25Gi",
		},
		{
			diskSizeBytes: 214748364800,
			partitionSize: "107%",
			output:        0,
			err:           "percentage cannot below 0 or bigger than 100%",
		},
		{
			diskSizeBytes: 214748364800,
			partitionSize: "100Gi",
			output:        107374182400,
		},
		{
			diskSizeBytes: 214748364800,
			partitionSize: "153600Mi",
			output:        161061273600,
		},
		{
			diskSizeBytes: 214748364800,
			partitionSize: "100Mi",
			output:        0,
			err:           "partition size cannot be smaller than 25Gi",
		},
		{
			diskSizeBytes: 214748364800,
			partitionSize: "100",
			output:        0,
			err:           "invalid disk space unit",
		},
		{
			diskSizeBytes: 214748364800,
			partitionSize: "500Gi",
			output:        0,
			err:           "partition size cannot be bigger than the whole disk",
		},
		{
			diskSizeBytes: 214748364800,
			partitionSize: "lolMi",
			output:        0,
			err:           "invalid partition size",
		},
	}

	for _, tCase := range testCases {
		actual, err := ParsePartitionSize(tCase.diskSizeBytes, tCase.partitionSize)
		assert.Equal(t, tCase.output, actual)
		if err != nil {
			assert.EqualError(t, err, tCase.err)
		}
	}
}
