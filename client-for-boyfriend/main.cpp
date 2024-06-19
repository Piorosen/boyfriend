#include <iostream>
// #include <audio/audio.h>

extern "C" void helloworld2(int data);

int main() {
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
    helloworld2(100);

    return 0;
}
