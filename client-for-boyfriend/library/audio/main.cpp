#include <iostream>
#include <audio/audio.h>
#include <args/args.h>

// https://github.com/triton-inference-server/client/blob/main/src/c%2B%2B/CMakeLists.txt

int main(int argc, char* argv[]) {
    arguments args = parse_arguments(argc, argv);

    // Use the arguments
    std::cout << "Inference Host: " << args.inference_host << std::endl;
    std::cout << "Inference Port: " << args.inference_port << std::endl;
    std::cout << "Voice Threshold: " << args.voice_threshold << std::endl;
    std::cout << "Voice Duration: " << args.voice_duration << std::endl;
    std::cout << "Voice Sample Rate: " << args.voice_sample_rate << std::endl;
    std::cout << "Voice Frames Buffer: " << args.voice_frames_buffer << std::endl;



    // std::vector<float> data;
    // audio::get_instance().setup(0.5, 5);
    // audio::get_instance().init();
    // audio::get_instance().start();
    // audio::get_instance().record(data);
    // audio::get_instance().close();
    //
    // audio::get_instance().play(data);
    // audio::get_instance().save_to_file("nya.wav", data);
    //
    // audio::get_instance().terminate();

    return 0;
}
