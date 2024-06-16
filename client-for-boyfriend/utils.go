package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"math"
	"os"
	"strings"
)

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
		// f := binary.LittleEndian.Uint32(inferResponse[i*4 : i*4+4])
		// outputData0[i] = math.Float32frombits(f)
		f := readFloat32(inferResponse[i*4 : i*4+4])
		outputData0[i] = f
	}
	return outputData0
}
func readFloat32(fourBytes []byte) float32 {
	buf := bytes.NewBuffer(fourBytes)
	var retval float32
	binary.Read(buf, binary.LittleEndian, &retval)
	return retval
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
