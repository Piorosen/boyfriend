package network

import (
	"bufio"
	"encoding/binary"
	"log"
	"math"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

// TODO: https://github.com/sunhailin-Leo/triton-service-go
// TODO: https://github.com/triton-inference-server/server/issues/6463
// TOOD: https://github.com/triton-inference-server/client/blob/main/src/c%2B%2B/examples/image_client.cc
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

// Convert int32 input data into raw bytes (assumes Little Endian)
func Preprocess(inputs []float32) []byte {
	var inputBytes []byte
	// Temp variable to hold our converted int32 -> []byte
	bs := make([]byte, 4)
	for i := 0; i < inputSize; i++ {
		binary.LittleEndian.PutUint32(bs, math.Float32bits(inputs[i]))
		inputBytes = append(inputBytes, bs...)
	}

	return inputBytes
}

// Convert output's raw bytes into int32 data (assumes Little Endian)
func Postprocess(inferResponse []byte) []float32 {
	outputData0 := make([]float32, outputSize)
	for i := 0; i < outputSize; i++ {
		f := binary.LittleEndian.Uint32(inferResponse[i*4 : i*4+4])
		outputData0[i] = math.Float32frombits(f)
	}
	return outputData0
}

// LoadLabels loads labels from a file
func loadLabels(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var labels []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		labels = append(labels, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return labels, nil
}

// Softmax function to convert logits to probabilities
func softmax(logits []float32) []float32 {
	maxLogit := logits[0]
	for _, logit := range logits {
		if logit > maxLogit {
			maxLogit = logit
		}
	}

	expSum := float32(0)
	probs := make([]float32, len(logits))
	for i, logit := range logits {
		probs[i] = float32(math.Exp(float64(logit - maxLogit)))
		expSum += probs[i]
	}
	for i := range probs {
		probs[i] /= expSum
	}

	return probs
}
