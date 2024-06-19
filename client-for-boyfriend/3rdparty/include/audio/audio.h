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
    static std::unique_ptr<audio> instance;
    audio(const audio&) = delete;
    audio(const audio&&) = delete;
    audio& operator=(const audio&) = delete;
    static std::once_flag initInstanceFlag;
    audio() {}

public:


    static audio& getInstance() {
        std::call_once(initInstanceFlag, &audio::initSingleton);
        return *instance;
    }

    static void initSingleton() {
        instance.reset(new audio);
    }


    void setup(float threshold = 0.5f, float duration = 0.5f, float sampleRate = 44100);

    void play(std::string const& file);
    void play(std::vector<uint8_t> const& memory);


};



#endif //AUDIO_H
