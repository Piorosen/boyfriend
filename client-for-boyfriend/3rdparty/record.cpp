#include <iostream>
#include <portaudio.h>
#include <sndfile.h>
#include <vector>
#include <cmath>

#define SAMPLE_RATE 44100
#define FRAMES_PER_BUFFER 512
#define THRESHOLD 0.5
#define RECORD_SECONDS 1

struct RecordData {
    std::vector<float> recordedSamples;
    bool recording;
    int samplesRecorded;
};

static int recordCallback(const void *inputBuffer, void *outputBuffer,
                          unsigned long framesPerBuffer,
                          const PaStreamCallbackTimeInfo*   timeInfo,
                          PaStreamCallbackFlags statusFlags,
                          void *userData) {
    RecordData *data = (RecordData*)userData;
    const float *in = (const float*)inputBuffer;

    if (!inputBuffer)
        return paContinue;

    if (!data->recording) {
        for (unsigned long i = 0; i < framesPerBuffer; i++) {
            if (std::fabs(in[i]) > THRESHOLD) {
                data->recording = true;
                data->samplesRecorded = 0;
                break;
            }
        }
    }

    if (data->recording) {
        for (unsigned long i = 0; i < framesPerBuffer; i++) {
            if (data->samplesRecorded < SAMPLE_RATE * RECORD_SECONDS) {
                data->recordedSamples.push_back(in[i]);
                data->samplesRecorded++;
            } else {
                return paComplete;
            }
        }
    }

    return paContinue;
}

void writeToWavFile(const std::vector<float> &samples, const char *filename) {
    SF_INFO sfinfo;
    sfinfo.frames = samples.size();
    sfinfo.samplerate = SAMPLE_RATE;
    sfinfo.channels = 1;
    sfinfo.format = SF_FORMAT_WAV | SF_FORMAT_PCM_16;

    SNDFILE *outfile = sf_open(filename, SFM_WRITE, &sfinfo);
    if (!outfile) {
        std::cerr << "Error opening output file: " << sf_strerror(outfile) << std::endl;
        return;
    }

    sf_write_float(outfile, samples.data(), samples.size());
    sf_close(outfile);
}

int main() {
    PaError err = Pa_Initialize();
    if (err != paNoError) {
        std::cerr << "PortAudio error: " << Pa_GetErrorText(err) << std::endl;
        return 1;
    }

    PaStream *stream;
    RecordData data;
    data.recording = false;
    data.samplesRecorded = 0;

    err = Pa_OpenDefaultStream(&stream,
                               1, // Input channels
                               0, // Output channels
                               paFloat32, // Sample format
                               SAMPLE_RATE,
                               FRAMES_PER_BUFFER,
                               recordCallback,
                               &data);
    if (err != paNoError) {
        std::cerr << "PortAudio error: " << Pa_GetErrorText(err) << std::endl;
        Pa_Terminate();
        return 1;
    }

    err = Pa_StartStream(stream);
    if (err != paNoError) {
        std::cerr << "PortAudio error: " << Pa_GetErrorText(err) << std::endl;
        Pa_CloseStream(stream);
        Pa_Terminate();
        return 1;
    }

    std::cout << "Listening for sound above threshold..." << std::endl;

    while (Pa_IsStreamActive(stream) == 1) {
        Pa_Sleep(100);
    }

    err = Pa_CloseStream(stream);
    if (err != paNoError) {
        std::cerr << "PortAudio error: " << Pa_GetErrorText(err) << std::endl;
    }

    Pa_Terminate();

    if (data.samplesRecorded > 0) {
        std::cout << "Writing to file..." << std::endl;
        writeToWavFile(data.recordedSamples, "output.wav");
        std::cout << "Recording saved to output.wav" << std::endl;
    } else {
        std::cout << "No recording was made." << std::endl;
    }

    return 0;
}
