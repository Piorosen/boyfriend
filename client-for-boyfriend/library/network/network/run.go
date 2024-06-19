package network

import (
	"fmt"
	"os"
	"strconv"

	_ "image/jpeg" // Register JPEG format
)

type Flags struct {
	ModelName    string
	ModelVersion string
	BatchSize    int
	URL          string
}

const (
	inputSize  = 3 * 224 * 224
	outputSize = 1000
)

func get_environment() (Flags, error) {
	name := os.Getenv("MODEL_NAME_FROM_TRITON")
	version := os.Getenv("MODEL_VERSION_FROM_TRITON")
	batch := os.Getenv("MODEL_BATCH_SIZE_FROM_TRITON")
	host := os.Getenv("INFERENCE_SERVER_TO_TRITON")

	if name == "" {
		return Flags{}, fmt.Errorf("MODEL_NAME_FROM_TRITON is not set")
	}
	if version == "" {
		return Flags{}, fmt.Errorf("MODEL_VERSION_FROM_TRITON is not set")
	}
	if batch == "" {
		return Flags{}, fmt.Errorf("MODEL_BATCH_SIZE_FROM_TRITON is not set")
	}
	if host == "" {
		return Flags{}, fmt.Errorf("INFERENCE_SERVER_TO_TRITON is not set")
	}
	num_batch, err := strconv.ParseInt(batch, 10, 32)
	if err != nil {
		return Flags{}, fmt.Errorf("BATCHSIZE ERROR %v", err)
	}
	return Flags{
		ModelName:    name,
		ModelVersion: version,
		BatchSize:    int(num_batch),
		URL:          host,
	}, nil
}
