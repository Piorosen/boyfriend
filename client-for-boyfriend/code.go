package main

//#cgo CXXFLAGS: -std=c++17
//#cgo LDFLAGS: -L./library -laudio
//#include "library/c_api.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type Audio struct {
	Threshold   float32
	Duration    float32
	SampleRate  float32
	FrameBuffer int
}

func NewAudio() *Audio {
	return &Audio{}
}

func (a *Audio) Demo() {
	C.demo()
}

func (a *Audio) Setup(threshold float32, duration float32, sampleRate float32, frameBuffer int) {
	a.Threshold = threshold
	a.Duration = duration
	a.SampleRate = sampleRate
	a.FrameBuffer = frameBuffer

	C.setup(C.float(threshold), C.float(duration), C.float(sampleRate), C.int(frameBuffer))
}

func (a *Audio) Init() {
	// Initialize audio system
	fmt.Println("Initializing audio system")
	C.init()
}

func (a *Audio) Start() {
	// Start audio processing
	fmt.Println("Starting audio processing")
	C.start()
}

func (a *Audio) Close() {
	// Close audio system
	fmt.Println("Closing audio system")
	C.close()
}

func (a *Audio) Terminate() {
	// Terminate audio system
	fmt.Println("Terminating audio system")
	C.terminate()
}

func (a *Audio) Play(memory []float32, sample_rate float32, channels int) {
	// Play audio from memory
	fmt.Println("Playing audio from memory")
	if sample_rate < 0 {
		a.SampleRate = sample_rate
	}
	// Audio playing logic
	C.play(C.float(sample_rate), C.int(channels))
}

func (a *Audio) Record() []float32 {
	// Record audio to memory
	fmt.Println("Recording audio to memory")
	// Audio recording logic
	var memory []float32

	ptr := (*C.float)(unsafe.Pointer(&memory))
	var length int

	C.record()
	memory = convertCFloatArray(ptr, length)

	return memory
}

func SaveToFile(file string, memory []float32, size int, sampleRate float32, channels int) error {
	if sampleRate == -1 {
		return errors.New("invalid sample rate")
	}
	// Save audio to file
	fmt.Printf("Saving audio to file: %s\n", file)
	// File saving logic

	return nil
}

// C.float 포인터를 Go 슬라이스로 변환하는 함수
func convertCFloatArray(ptr *C.float, length int) []float32 {
	// uintptr 타입의 포인터를 사용하여 슬라이스 헤더 생성
	sliceHeader := &struct {
		Data uintptr
		Len  int
		Cap  int
	}{
		Data: uintptr(unsafe.Pointer(ptr)),
		Len:  length,
		Cap:  length,
	}

	// 슬라이스 헤더를 float32 슬라이스로 변환
	return *(*[]float32)(unsafe.Pointer(sliceHeader))
}
