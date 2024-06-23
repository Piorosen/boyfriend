//
// Created by ChaCha on 6/19/24.
//


#ifndef AUDIO_IMPL_H
#define AUDIO_IMPL_H

#include <portaudio.h>

struct audio_data {
    std::vector<float> buffer;
    bool is_recording = false;

    float threshold;
    float duration;
    float sample_rate;
    int frame_buffer;
};

class audio_impl {
private:
    PaStream* input_stream;
    audio_data received_data;

public:
    void setup(float threshold, float duration, float sample_rate, int frame_buffer);
    void start();
    void close();
    void record(std::vector<float>& memory);
};


#endif

