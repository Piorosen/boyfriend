//
// Created by ChaCha on 6/19/24.
//

#include <vector>

#include <audio/audio.h>
#include <audio/c_api.h>

void setup(float threshold, float duration, float sample_rate, int frame_buffer) {
    audio::get_instance().setup(threshold, duration, sample_rate, frame_buffer);
}

void init() {
    audio::get_instance().init();
}

void start() {
    audio::get_instance().start();
}
void close() {
    audio::get_instance().close();
}
void terminate() {
    audio::get_instance().terminate();
}
//bool play(const std::string& file);
void play(const float* memory, int* const size, float sample_rate, int channels) {
    std::vector<float> mem;
    audio::get_instance().play(mem, sample_rate, channels);
    memory = mem.data();
    *size = mem.size();
}
//bool record(const std::string& file, int channels = 1);
void record(const float* memory, int* const size) {
    std::vector<float> mem;
    audio::get_instance().record(mem);
    memory = mem.data();
    *size = mem.size();
}// pcm

bool save_to_file(const char* file, const float* memory, int size, float sample_rate, int channels) {
    std::vector<float> mem;
    mem.reserve(size);
    mem.insert(mem.end(), memory, memory + size);
    return audio::get_instance().save_to_file(file, mem, sample_rate, channels);
}