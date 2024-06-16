// Copyright (c) 2019-2020, NVIDIA CORPORATION. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//  * Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//  * Neither the name of NVIDIA CORPORATION nor the names of its
//    contributors may be used to endorse or promote products derived
//    from this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS ``AS IS'' AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
// PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL THE COPYRIGHT OWNER OR
// CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
// EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY
// OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package main

import (
	"flag"
	"fmt"
	"log"

	_ "image/jpeg" // Register JPEG format

	"github.com/Piorosen/boyfriend/client-for-boyfriend/service"
	"github.com/disintegration/imaging"
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

func parseFlags() Flags {
	var flags Flags
	// https://github.com/NVIDIA/triton-inference-server/tree/master/docs/examples/model_repository/simple
	flag.StringVar(&flags.ModelName, "m", "densenet_onnx", "Name of model being served. (Required)")
	flag.StringVar(&flags.ModelVersion, "x", "1", "Version of model. Default: Latest Version.")
	flag.IntVar(&flags.BatchSize, "b", 1, "Batch size. Default: 1.")
	flag.StringVar(&flags.URL, "u", "localhost:8001", "Inference Server URL. Default: localhost:8001")
	flag.Parse()
	return flags
}

// Function to read an image file, resize it to 224x224 pixels, and output a float array
func ImageToFloatArray(fileName string) ([]float32, error) {
	// Open a test image.
	src, err := imaging.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, 224, 224, imaging.Lanczos)
	width := 224
	height := 224
	channels := 3
	pixels := make([]float32, channels*width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := src.At(x, y).RGBA()
			// NCHW
			// Normalize to [-1, 1] and store in the array
			pixels[0*width*height+y*width+x] = (float32(r&0xff) / 127.5) - 1
			pixels[1*width*height+y*width+x] = (float32(g&0xff) / 127.5) - 1
			pixels[2*width*height+y*width+x] = (float32(b&0xff) / 127.5) - 1
			// NHWC
			// pixels[224*3*y+3*x+0] = (float32(r&0xff) / 127.5) - 1
			// pixels[224*3*y+3*x+1] = (float32(g&0xff) / 127.5) - 1
			// pixels[224*3*y+3*x+2] = (float32(b&0xff) / 127.5) - 1
		}
	}
	// fmt.Printf("%v", imaging.Save(src, "asdf.jpg"))
	return pixels, nil
}

func main() {
	fileName := "resources/mug.jpg"
	floatArray, err := ImageToFloatArray(fileName)
	if err != nil {
		log.Fatalf("Error processing image: %v", err)
	}

	FLAGS := parseFlags()
	fmt.Println("FLAGS:", FLAGS)

	client := service.NewClient()
	client.Open(FLAGS.URL)
	defer client.Close()

	serverLiveResponse := client.ServerLiveRequest()
	fmt.Printf("Triton Health - Live: %v\n", serverLiveResponse.Live)

	serverReadyResponse := client.ServerReadyRequest()
	fmt.Printf("Triton Health - Ready: %v\n", serverReadyResponse.Ready)

	modelMetadataResponse := client.ModelMetadataRequest(FLAGS.ModelName, FLAGS.ModelVersion)
	fmt.Println(modelMetadataResponse)
	input := Preprocess(floatArray)
	// /* We use a simple model that takes 2 input tensors of 16 integers
	// each and returns 2 output tensors of 16 integers each. One
	// output tensor is the element-wise sum of the inputs and one
	// output is the element-wise difference. */
	inferResponse := client.ModelInferRequest(input, FLAGS.ModelName, FLAGS.ModelVersion)
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
