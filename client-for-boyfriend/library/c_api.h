//
// Created by ChaCha on 6/19/24.
//

#ifndef BINDING_H
#define BINDING_H

#ifdef __cplusplus
extern "C" {
#endif

// void setup(float threshold = 0.5f, float duration = 0.5f, float sample_rate = 44100, int frame_buffer = 256);
void setup(float threshold, float duration, float sample_rate, int frame_buffer);
void init();
void start();
void close();
void terminate();
//bool play(const std::string& file);
void play(float sample_rate, int channels); // pcm
//bool record(const std::string& file, int channels = 1);
void record(); // pcm
void demo();
unsigned char save_to_file(const char* file, float sample_rate, int channels);
#ifdef __cplusplus
}
#endif

#endif //BINDING_H
