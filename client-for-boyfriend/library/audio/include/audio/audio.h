//
// Created by ChaCha on 6/19/24.
//

#ifndef AUDIO_H
#define AUDIO_H

#include <string>
#include <vector>
#include <stdint.h>
#include <memory>
#include <mutex>

class audio {
private:
    audio(const audio&) = delete;
    audio(const audio&&) = delete;
    audio& operator=(const audio&) = delete;
    audio();

    class audio_impl* impl_;

    float threshold;
    float duration;
    float sample_rate;
    int frame_buffer;
public:
    ~audio();

    static audio& get_instance() {
        static std::unique_ptr<audio> instance(new audio);
        return *instance;
    }

    void setup(float threshold = 0.5f, float duration = 0.5f, float sample_rate = 44100, int frame_buffer = 512);

    void init();
    void start();
    void close();
    void terminate();

    bool play(const std::string& file);
    void play(const std::vector<float>& memory, float sample_rate = -1, int channels = 1); // pcm

    bool record(const std::string& file, int channels = 1);
    void record(std::vector<float>& memory); // pcm

    bool save_to_file(const std::string& file, const std::vector<float>& memory, float sample_rate = -1, int channels = 1);
};



#endif //AUDIO_H
