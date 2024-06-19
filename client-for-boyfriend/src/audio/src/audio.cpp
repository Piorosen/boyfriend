//
// Created by ChaCha on 6/19/24.
//


#include <audio/audio.h>
#include <internal/utils.h>
#include <internal/audio_impl.h>

// #define SAMPLE_RATE 44100
// #define SECONDS_TO_RECORD 5
// #define FRAMES_PER_BUFFER 512
// #define THRESHOLD 0.1

audio::audio() {
    this->impl_ = new audio_impl;
}
audio::~audio() {
    delete this->impl_;
}

void audio::setup(float threshold, float duration, float sample_rate, int frame_buffer) {
    this->impl_->setup(threshold, duration, sample_rate, frame_buffer);
    this->threshold = threshold;
    this->duration = duration;
    this->sample_rate = sample_rate;
    this->frame_buffer = frame_buffer;
}

void audio::init() {
    Pa_Initialize();
}

void audio::start() {
    this->impl_->start();
}
void audio::close() {
    this->impl_->close();
}
void audio::terminate() {
    Pa_Terminate();
}

bool audio::play(const std::string& file) {
    std::vector<float> memory;
    int32_t sample_rate;
    int32_t channels;

    if (!load_pcm_from_mpeg(file, memory, sample_rate, channels)) {
        return false;
    }
    this->play(memory, (float)sample_rate);

    return true;
}

void audio::play(const std::vector<float>& memory, float sample_rate, int channels) {
    if (sample_rate < 0) {
        sample_rate = this->sample_rate;
    }

    PaStream* stream;
    Pa_OpenDefaultStream(&stream,
        0,
        channels,
        paFloat32,
        sample_rate,
        this->frame_buffer,
        nullptr,
        nullptr);

    Pa_StartStream(stream);
    Pa_WriteStream(stream, memory.data(), memory.size());
    Pa_Sleep(5*1000);
    Pa_StopStream(stream);
    Pa_CloseStream(stream);

} // pcm

bool audio::record(const std::string& file, int channels) {
    std::vector<float> memory;
    this->impl_->record(memory);
    return save_pcm_to_mpeg(file, memory, this->sample_rate, channels);
}

void audio::record(std::vector<float>& memory) {
    this->impl_->record(memory);
} // pcm

bool audio::save_to_file(const std::string& file, const std::vector<float>& memory, float sample_rate, int channels) {
    // save_pcm_to_mpeg(file, memory, this->sample_rate, channels)
    if (sample_rate < 0) {
        sample_rate = this->sample_rate;
    }

    return save_pcm_to_mpeg(file, memory, sample_rate, channels);
}