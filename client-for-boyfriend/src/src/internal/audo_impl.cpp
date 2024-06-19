//
// Created by ChaCha on 6/19/24.
//

#include <vector>

#include <internal/audio_impl.h>
#include <internal/recorder.h>

void audio_impl::setup(float threshold, float duration, float sample_rate, int frame_buffer) {
    received_data.threshold = threshold;
    received_data.duration = duration;
    received_data.sample_rate = sample_rate;
    received_data.frame_buffer = frame_buffer;
}

void audio_impl::start() {
    Pa_OpenDefaultStream(&input_stream, 1, 0, paFloat32, received_data.sample_rate, received_data.frame_buffer, audioCallback, &received_data);
}

void audio_impl::close() {
    Pa_StopStream(input_stream);
    Pa_CloseStream(input_stream);
}

void audio_impl::record(std::vector<float>& memory) {
    received_data.is_recording = true;
    Pa_StartStream(input_stream);

    this->received_data.buffer.reserve(static_cast<int>(this->received_data.sample_rate * this->received_data.duration));
    this->received_data.buffer.clear();

    while (Pa_IsStreamActive(input_stream)) {
        Pa_Sleep(100);
    }
    memory = std::move(this->received_data.buffer);
}
