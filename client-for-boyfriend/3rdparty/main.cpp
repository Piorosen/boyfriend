#include <iostream>
#include <fstream>
#include <thread>
#include <soundio/soundio.h>

const int sampleRate = 44100;
const int bufferLen = 1024;
const int threshold = 500;
const int recordSeconds = 5;

bool detectSound(const int16_t *data, int length) {
    for (int i = 0; i < length; i++) {
        if (std::abs(data[i]) > threshold) {
            return true;
        }
    }
    return false;
}

void writeToFile(const int16_t *data, int length, std::ofstream &file) {
    file.write(reinterpret_cast<const char*>(data), length * sizeof(int16_t));
}

void recordCallback(struct SoundIoInStream *instream, int frame_count_min, int frame_count_max) {
    static std::ofstream file("output.raw", std::ios::binary | std::ios::app);
    static int16_t buffer[bufferLen];
    static int framesRecorded = 0;

    struct SoundIoChannelArea *areas;
    int frame_count = bufferLen;
    int err;

    if ((err = soundio_instream_begin_read(instream, &areas, &frame_count))) {
        std::cerr << "Error beginning read: " << soundio_strerror(err) << std::endl;
        return;
    }

    if (!frame_count)
        return;

    for (int frame = 0; frame < frame_count; frame += 1) {
        for (int channel = 0; channel < instream->layout.channel_count; channel += 1) {
            buffer[frame] = *((int16_t*)(areas[channel].ptr + areas[channel].step * frame));
        }
    }

    if ((err = soundio_instream_end_read(instream))) {
        std::cerr << "Error ending read: " << soundio_strerror(err) << std::endl;
        return;
    }

    if (detectSound(buffer, frame_count)) {
        writeToFile(buffer, frame_count, file);
        framesRecorded += frame_count;
        std::cout << framesRecorded << " " << sampleRate * recordSeconds << std::endl;
        if (framesRecorded >= sampleRate * recordSeconds) {
            std::cout << "Recording complete." << std::endl;
            soundio_instream_pause(instream, true);
        }
    }
}

int main() {
    struct SoundIo *soundio = soundio_create();
    if (!soundio) {
        std::cerr << "Out of memory" << std::endl;
        return 1;
    }

    int err = soundio_connect(soundio);
    if (err) {
        std::cerr << "Error connecting: " << soundio_strerror(err) << std::endl;
        return 1;
    }

    soundio_flush_events(soundio);

    int default_input_device_index = soundio_default_input_device_index(soundio);
    if (default_input_device_index < 0) {
        std::cerr << "No input device found" << std::endl;
        return 1;
    }

    struct SoundIoDevice *device = soundio_get_input_device(soundio, default_input_device_index);
    if (!device) {
        std::cerr << "Out of memory" << std::endl;
        return 1;
    }

    std::cout << "Input device: " << device->name << std::endl;
    struct SoundIoInStream *instream = soundio_instream_create(device);
    if (!instream) {
        std::cerr << "Out of memory" << std::endl;
        return 1;
    }

//    instream->format = SoundIoFormatS16LE;
    instream->format = device->current_format;
//    instream->sample_rate = device->sample_rates[device->sample_rate_current].max;
    instream->read_callback = recordCallback;

    if ((err = soundio_instream_open(instream))) {
        std::cerr << "Unable to open input stream: " << soundio_strerror(err) << std::endl;
        return 1;
    }

    if ((err = soundio_instream_start(instream))) {
        std::cerr << "Unable to start input stream: " << soundio_strerror(err) << std::endl;
        return 1;
    }

    std::cout << "Listening for sound..." << std::endl;

    while (true) {
        soundio_flush_events(soundio);
        std::this_thread::sleep_for(std::chrono::milliseconds(100));
    }

    soundio_instream_destroy(instream);
    soundio_device_unref(device);
    soundio_destroy(soundio);

    return 0;
}
