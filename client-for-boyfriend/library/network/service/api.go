package service

import (
	"context"
	"log"
	"time"

	triton "github.com/Piorosen/boyfriend/client-for-boyfriend/library/network/grpc-client"
)

func (client *Client) ServerLiveRequest() *triton.ServerLiveResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverLiveRequest := triton.ServerLiveRequest{}
	// Submit ServerLive request to server
	serverLiveResponse, err := client.client.ServerLive(ctx, &serverLiveRequest)
	if err != nil {
		log.Fatalf("Couldn't get server live: %v", err)
	}

	return serverLiveResponse
}

func (client *Client) ServerReadyRequest() *triton.ServerReadyResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serverReadyRequest := triton.ServerReadyRequest{}
	// Submit ServerReady request to server
	serverReadyResponse, err := client.client.ServerReady(ctx, &serverReadyRequest)
	if err != nil {
		log.Fatalf("Couldn't get server ready: %v", err)
	}
	return serverReadyResponse
}

func (client *Client) ModelMetadataRequest(modelName string, modelVersion string) *triton.ModelMetadataResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create status request for a given model
	modelMetadataRequest := triton.ModelMetadataRequest{
		Name:    modelName,
		Version: modelVersion,
	}
	// Submit modelMetadata request to server
	modelMetadataResponse, err := client.client.ModelMetadata(ctx, &modelMetadataRequest)
	if err != nil {
		log.Fatalf("Couldn't get server model metadata: %v", err)
	}
	return modelMetadataResponse
}

func (client *Client) ModelInferRequest(rawInput []byte, modelName string, modelVersion string) *triton.ModelInferResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create request input tensors
	inferInputs := []*triton.ModelInferRequest_InferInputTensor{
		{
			Name:     "data_0",
			Datatype: "FP32",
			Shape:    []int64{3, 224, 224},
		},
	}

	// Create request input output tensors
	inferOutputs := []*triton.ModelInferRequest_InferRequestedOutputTensor{
		{
			Name: "fc6_1",
			// Parameters: map[string]*triton.InferParameter{
			// 	"classification": {ParameterChoice: &triton.InferParameter_Int64Param{Int64Param: 2}},
			// },
		},
	}

	// Create inference request for specific model/version
	modelInferRequest := triton.ModelInferRequest{
		ModelName:    modelName,
		ModelVersion: modelVersion,
		Inputs:       inferInputs,
		Outputs:      inferOutputs,
	}

	modelInferRequest.RawInputContents = append(modelInferRequest.RawInputContents, rawInput)

	// Submit inference request to server
	modelInferResponse, err := client.client.ModelInfer(ctx, &modelInferRequest)
	if err != nil {
		log.Fatalf("Error processing InferRequest: %v", err)
	}
	return modelInferResponse
}

func (client *Client) RawModelInferRequest(rawInput []byte, modelName string, modelVersion string) *triton.ModelInferResponse {
	// Create context for our request with 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// inferRequestHeader := &triton.inferRequestHeader{
	// 	Input: []*trtis.InferRequestHeader_Input{
	// 		&trtis.InferRequestHeader_Input{
	// 			Name: "INPUT",
	// 		},
	// 	},
	// 	Output: []*trtis.InferRequestHeader_Output{
	// 		&trtis.InferRequestHeader_Output{
	// 			Name: "OUTPUT",
	// 		},
	// 	},
	// 	BatchSize: uint32(batchSize),
	// }
	// Create request input tensors
	inferInputs := []*triton.ModelInferRequest_InferInputTensor{
		{
			Name:     "data_0",
			Datatype: "FP32",
			Shape:    []int64{3, 224, 224},
		},
	}

	// Create request input output tensors
	inferOutputs := []*triton.ModelInferRequest_InferRequestedOutputTensor{
		{
			Name: "fc6_1",
			// Parameters: map[string]*triton.InferParameter{
			// 	"classification": {ParameterChoice: &triton.InferParameter_Int64Param{Int64Param: 2}},
			// },
		},
	}

	// Create inference request for specific model/version
	modelInferRequest := triton.ModelInferRequest{
		ModelName:    modelName,
		ModelVersion: modelVersion,
		Inputs:       inferInputs,
		Outputs:      inferOutputs,
	}

	modelInferRequest.RawInputContents = append(modelInferRequest.RawInputContents, rawInput)

	// Submit inference request to server
	modelInferResponse, err := client.client.ModelInfer(ctx, &modelInferRequest)
	if err != nil {
		log.Fatalf("Error processing InferRequest: %v", err)
	}
	return modelInferResponse
}
