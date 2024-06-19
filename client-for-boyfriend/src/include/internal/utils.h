//
// Created by ChaCha on 6/19/24.
//

#ifndef UTILS_H
#define UTILS_H

#include <string>
#include <stdint.h>
#include <vector>

bool load_pcm_from_mpeg(const std::string& filename,
    std::vector<float>& audioBuffer,
    int32_t& rate,
    int32_t& channels);
bool save_pcm_to_mpeg(const std::string& filename,
    const std::vector<float>& audioBuffer,
    int32_t rate,
    int32_t channels);

#endif //UTILS_H
