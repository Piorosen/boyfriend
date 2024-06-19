//
// Created by ChaCha on 6/19/24.
//

#include <internal/utils.h>
#include <sndfile.h>

bool load_pcm_from_mpeg(const std::string& filename, std::vector<float>& buffer, int32_t& rate, int32_t& channels) {
    SF_INFO sfInfo;
    SNDFILE* inFile = sf_open(filename.c_str(), SFM_READ, &sfInfo);
    if (inFile == nullptr) {
        // std::cerr << "Failed to open file for reading: " << sf_strerror(inFile) << std::endl;
        return false;
    }
    rate = sfInfo.samplerate;
    channels = sfInfo.channels;

    buffer.reserve(sfInfo.frames);
    sf_read_float(inFile, buffer.data(), sfInfo.frames);
    sf_close(inFile);
    return true;
}

bool save_pcm_to_mpeg(const std::string& filename, const std::vector<float>& buffer, int32_t rate, int32_t channels) {
    SF_INFO sfInfo;
    sfInfo.channels = channels;
    sfInfo.samplerate = rate;
    sfInfo.format = SF_FORMAT_WAV | SF_FORMAT_PCM_16;

    SNDFILE* outFile = sf_open(filename.c_str(), SFM_WRITE, &sfInfo);
    if (outFile == nullptr) {
        // std::cerr << "Failed to open file for writing: " << sf_strerror(outFile) << std::endl;
        return false;
    }

    sf_write_float(outFile, buffer.data(), buffer.size());
    sf_close(outFile);
    return true;
}
