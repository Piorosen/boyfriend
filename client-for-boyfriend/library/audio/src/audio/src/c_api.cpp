//
// Created by ChaCha on 6/19/24.
//

#include <vector>

#include <audio/audio.h>
#include <audio/c_api.h>

std::vector<float> globally_vector;

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
void play(float sample_rate, int channels) {
    audio::get_instance().play(globally_vector, sample_rate, channels);
}
//bool record(const std::string& file, int channels = 1);
void record() {
    globally_vector.clear();
    audio::get_instance().record(globally_vector);
}// pcm

unsigned char save_to_file(const char* file, float sample_rate, int channels) {
    return audio::get_instance().save_to_file(file, globally_vector, sample_rate, channels);
}

void demo() {
    std::vector<float> data;
    audio::get_instance().setup(0.5, 5);
    audio::get_instance().init();
    audio::get_instance().start();
    audio::get_instance().record(data);
    audio::get_instance().close();

    audio::get_instance().play(data);
    audio::get_instance().save_to_file("nya.wav", data);

    audio::get_instance().terminate();
}
