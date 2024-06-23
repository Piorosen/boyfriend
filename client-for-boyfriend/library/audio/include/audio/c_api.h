//
// Created by ChaCha on 6/19/24.
//

#ifndef BINDING_H
#define BINDING_H

#ifdef __cplusplus
extern "C" {
#endif
void setup(float threshold = 0.5f, float duration = 0.5f, float sample_rate = 44100, int frame_buffer = 256);
void init();
void start();
void close();
void terminate();
//bool play(const std::string& file);
void play(const float* memory, int* const size, float sample_rate = -1, int channels = 1); // pcm
//bool record(const std::string& file, int channels = 1);
void record(const float* memory, int* const size); // pcm

bool save_to_file(const char* file, const float* memory, int size, float sample_rate = -1, int channels = 1);
#ifdef __cplusplus
}
#endif

#endif //BINDING_H
