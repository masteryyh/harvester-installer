package util

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	mebibyte = 1024 * 1024
	gibibyte = 1024 * mebibyte
)

func StringSliceContains(sSlice []string, s string) bool {
	for _, target := range sSlice {
		if target == s {
			return true
		}
	}
	return false
}

func DupStrings(src []string) []string {
	if src == nil {
		return nil
	}
	s := make([]string, len(src))
	copy(s, src)
	return s
}

func ParsePartitionSize(diskSizeBytes uint64, partitionSize string) (uint64, error) {
	if partitionSize == "" {
		return 0, fmt.Errorf("partition size cannot be empty")
	}

	if strings.HasSuffix(partitionSize, "%") {
		percentage := strings.TrimSuffix(partitionSize, "%")
		percentageFloat, err := strconv.ParseFloat(percentage, 64)

		if err != nil {
			return 0, fmt.Errorf("invalid partition size")
		}
		if percentageFloat <= 0 || percentageFloat > 100 {
			return 0, fmt.Errorf("percentage cannot below 0 or bigger than 100%%")
		}

		actualSize := uint64(float64(diskSizeBytes) * (percentageFloat / 100))
		if actualSize > diskSizeBytes {
			return 0, fmt.Errorf("partition size cannot be bigger than the whole disk")
		}
		if actualSize/gibibyte < 25 {
			return 0, fmt.Errorf("partition size cannot be smaller than 25Gi")
		}

		return actualSize, nil
	}

	unit := partitionSize[len(partitionSize)-2:]
	if unit != "Gi" && unit != "Mi" {
		return 0, fmt.Errorf("invalid disk space unit")
	}

	size := partitionSize[:len(partitionSize)-2]
	sizeInt, err := strconv.ParseUint(size, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid partition size")
	}

	var actualSize uint64
	switch unit {
	case "Gi":
		actualSize = sizeInt * gibibyte
	case "Mi":
		actualSize = sizeInt * mebibyte
	}

	if actualSize > diskSizeBytes {
		return 0, fmt.Errorf("partition size cannot be bigger than the whole disk")
	}
	if actualSize/gibibyte < 25 {
		return 0, fmt.Errorf("partition size cannot be smaller than 25Gi")
	}

	return actualSize, nil
}
