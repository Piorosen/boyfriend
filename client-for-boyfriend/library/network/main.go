package main

/*
#include <stdlib.h>
#include <stdint.h>

typedef struct {
	uint8_t* data;
	int32_t shape[4];
} inference_request;

typedef struct {
	uint8_t* data;
	int32_t shape;
} inference_response;
*/
import "C"

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/Piorosen/boyfriend/client-for-boyfriend/library/network/network"
	"github.com/Piorosen/boyfriend/client-for-boyfriend/library/network/service"
)

var triton_host string
var client = service.NewClient()

//export set_environment
func set_environment(host *C.char, port int) {
	triton_host = fmt.Sprintf("%s:%d", ConvertC2Go(host), port)
}

func server_open() {
	client.Open(triton_host)
}

func server_close() {
	client.Close()
}

//export server_live_request
func server_live_request() *C.char {
	result := client.ServerLiveRequest()
	return ConvertGo2C(result.String())
}

//export server_ready_request
func server_ready_request() *C.char {
	result := client.ServerReadyRequest()
	return ConvertGo2C(result.String())
}

//export model_metadata_request
func model_metadata_request(model_name, version *C.char) *C.char {
	name := ConvertC2Go(model_name)
	v := ConvertC2Go(version)
	result := client.ModelMetadataRequest(name, v)
	return ConvertGo2C(result.String())
}

//export inference
func inference(name, version *C.char, request *C.inference_request, response *C.inference_response) {
	shape := 1
	for _, i := range request.shape {
		shape *= int(i)
	}
	n := ConvertC2Go(name)
	v := ConvertC2Go(version)

	data := C.GoBytes(unsafe.Pointer(request.data), C.int(shape))
	resp := client.ModelInferRequest(data, n, v)
	response.shape = 1000
	response.data = (*C.uint8_t)(C.CBytes(resp.RawOutputContents[0]))
}

//export inference_demo
func inference_demo(name, version, path_jpg, path_labels *C.char) *C.char {
	n := ConvertC2Go(name)
	v := ConvertC2Go(version)
	jpg := ConvertC2Go(path_jpg)
	label := ConvertC2Go(path_labels)

	// fileName := "resources/mug.jpg"
	floatArray, err := network.ImageToFloatArray(jpg)
	if err != nil {
		log.Fatalf("Error processing image: %v", err)
	}
	labels, err := network.LoadLabels(label)

	input := network.Preprocess(floatArray)
	inferResponse := client.ModelInferRequest(input, n, v)

	outputs := network.Postprocess(inferResponse.RawOutputContents[0])
	outputs = network.Softmax(outputs)

	result := "\nChecking Inference Outputs\n--------------------------\n"
	for i := 0; i < 1000; i++ {
		if outputs[i]*100 > 5 {
			result += fmt.Sprintf("%d : %s : %.1f %%\n", i, labels[i], outputs[i]*100)
		}
	}
	return ConvertGo2C(result)
}

func main() {}
