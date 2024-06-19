//
// Created by ChaCha on 6/19/24.
//

#ifndef RECORDER_H
#define RECORDER_H

#include <portaudio.h>


static int audioCallback(const void* inputBuffer, void* outputBuffer,
                         unsigned long framesPerBuffer,
                         const PaStreamCallbackTimeInfo* timeInfo,
                         PaStreamCallbackFlags statusFlags,
                         void* userData) {
    auto* data = static_cast<audio_data*>(userData);
    const auto* in = static_cast<const float*>(inputBuffer);
    auto max_size = static_cast<size_t>(data->sample_rate * data->duration);

    if (inputBuffer != nullptr) {
        for (unsigned int i = 0; i < framesPerBuffer; ++i) {
            if (std::abs(in[i]) > data->threshold) {
                data->is_recording = true;
                break;
            }
        }

        if (data->is_recording) {
            // if - => fin
            // + add framePerbuffer
            auto total_remain = max_size - data->buffer.size();
            auto insert_value = std::min(total_remain, framesPerBuffer);

            data->buffer.insert(data->buffer.end(), in + 0, in + insert_value);

            if (data->buffer.size() >= max_size) {
                return paComplete;
            }
        }
    }
    return paContinue;
}

#endif //RECORDER_H
