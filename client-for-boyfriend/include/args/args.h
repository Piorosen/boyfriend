//
// Created by ChaCha on 6/19/24.
//

#ifndef ARGS_H
#define ARGS_H
#include <string>

struct arguments {
  std::string inference_host = "127.0.0.1";
  int inference_port = 8080;

  float voice_threshold = 0.5f;
  float voice_duration = 1.0f;
  float voice_sample_rate = 44100.0f;
  int voice_frames_buffer = 512;
};

arguments parse_arguments(int argc, char* argv[]);

#endif //ARGS_H
