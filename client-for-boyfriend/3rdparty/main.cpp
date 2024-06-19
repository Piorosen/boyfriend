#include <iostream>
#include <fstream>
#include <vector>
#include <cmath>
#include <portaudio.h>
#include <sndfile.h>

#define SAMPLE_RATE 44100
#define SECONDS_TO_RECORD 5
#define FRAMES_PER_BUFFER 512
#define THRESHOLD 0.1

struct AudioData {
    std::vector<float> buffer;
    bool isRecording = false;
};

static int audioCallback(const void* inputBuffer, void* outputBuffer,
                         unsigned long framesPerBuffer,
                         const PaStreamCallbackTimeInfo* timeInfo,
                         PaStreamCallbackFlags statusFlags,
                         void* userData) {
    auto* data = static_cast<AudioData*>(userData);
    const auto* in = static_cast<const float*>(inputBuffer);

    if (inputBuffer != nullptr) {
        for (unsigned int i = 0; i < framesPerBuffer; ++i) {
            if (std::abs(in[i]) > THRESHOLD) {
                data->isRecording = true;
            }
            if (data->isRecording) {
                data->buffer.push_back(in[i]);
                if (data->buffer.size() >= SAMPLE_RATE * SECONDS_TO_RECORD) {
                    return paComplete;
                }
            }
        }
    }
    return paContinue;
}

void writeWavFile(const std::string& filename, const std::vector<float>& buffer) {
    SF_INFO sfInfo;
    sfInfo.channels = 1;
    sfInfo.samplerate = SAMPLE_RATE;
    sfInfo.format = SF_FORMAT_WAV | SF_FORMAT_PCM_16;

    SNDFILE* outFile = sf_open(filename.c_str(), SFM_WRITE, &sfInfo);
    if (outFile == nullptr) {
        std::cerr << "Failed to open file for writing: " << sf_strerror(outFile) << std::endl;
        return;
    }

    sf_write_float(outFile, buffer.data(), buffer.size());
    sf_close(outFile);
}

void playWavFile(const std::string& filename) {
    SF_INFO sfInfo;
    SNDFILE* inFile = sf_open(filename.c_str(), SFM_READ, &sfInfo);
    if (inFile == nullptr) {
        std::cerr << "Failed to open file for reading: " << sf_strerror(inFile) << std::endl;
        return;
    }

    std::vector<float> buffer(sfInfo.frames);
    sf_read_float(inFile, buffer.data(), sfInfo.frames);
    sf_close(inFile);

    PaStream* stream;

    Pa_OpenDefaultStream(&stream, 0, 1, paFloat32, sfInfo.samplerate, FRAMES_PER_BUFFER, nullptr, nullptr);
    Pa_StartStream(stream);

    Pa_WriteStream(stream, buffer.data(), buffer.size());

    Pa_StopStream(stream);
    Pa_CloseStream(stream);
}

int main() {
    Pa_Initialize();

    AudioData data;
    PaStream* stream;
    Pa_OpenDefaultStream(&stream, 1, 0, paFloat32, SAMPLE_RATE, FRAMES_PER_BUFFER, audioCallback, &data);
    Pa_StartStream(stream);

    std::cout << "Listening for audio threshold...\n";
    while (Pa_IsStreamActive(stream)) {
        Pa_Sleep(100);
    }

    Pa_StopStream(stream);
    Pa_CloseStream(stream);


    if (data.buffer.empty()) {
        std::cout << "No audio threshold exceeded, no recording made.\n";
        return 0;
    }

    std::string filename = "recorded.wav";
    writeWavFile(filename, data.buffer);
    std::cout << "Recording saved to " << filename << "\n";

    std::cout << "Playing back recording...\n";
    playWavFile(filename);
    Pa_Terminate();
    return 0;
}
