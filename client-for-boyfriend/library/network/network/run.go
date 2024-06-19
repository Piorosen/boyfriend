package network

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "image/jpeg" // Register JPEG format

	"github.com/Piorosen/boyfriend/client-for-boyfriend/library/network/service"
	"github.com/subosito/gotenv"
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

func entrypoint() {
	gotenv.Load()
	env, err := get_environment()
	if err != nil {
		log.Fatalf("Error load environemnt: %v", err)
	}

	client := service.NewClient()
	client.Open(env.URL)
	defer client.Close()

	serverLiveResponse := client.ServerLiveRequest()
	fmt.Printf("Triton Health - Live: %v\n", serverLiveResponse.Live)

	serverReadyResponse := client.ServerReadyRequest()
	fmt.Printf("Triton Health - Ready: %v\n", serverReadyResponse.Ready)

	modelMetadataResponse := client.ModelMetadataRequest(env.ModelName, env.ModelVersion)
	fmt.Println(modelMetadataResponse)

	fileName := "resources/mug.jpg"
	floatArray, err := ImageToFloatArray(fileName)
	if err != nil {
		log.Fatalf("Error processing image: %v", err)
	}
	input := Preprocess(floatArray)
	// /* We use a simple model that takes 2 input tensors of 16 integers
	// each and returns 2 output tensors of 16 integers each. One
	// output tensor is the element-wise sum of the inputs and one
	// output is the element-wise difference. */
	inferResponse := client.ModelInferRequest(input, env.ModelName, env.ModelVersion)
	// /* We expect there to be 2 results (each with batch-size 1). Walk
	// over all 16 result elements and print the sum and difference
	// calculated by the model. */

	labels, err := loadLabels("./resources/densenet_labels.txt")
	outputs := Postprocess(inferResponse.RawOutputContents[0])
	outputs = softmax(outputs)

	if err != nil {
		log.Fatalf("Error loading labels from file: %v", err)
	}
	// outputs := string(inferResponse.RawOutputContents[0])

	fmt.Println("\nChecking Inference Outputs\n--------------------------")

	for i := 0; i < outputSize; i++ {
		if outputs[i]*100 > 5 {
			fmt.Printf("%d : %s : %.1f %%\n", i, labels[i], outputs[i]*100)
		}
	}
}
