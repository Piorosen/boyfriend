//
// Created by ChaCha on 6/19/24.
//
#include <iostream>
#include <args/args.h>
#include <cxxopts.hpp>

arguments parse_arguments(int argc, char* argv[]) {
    arguments args;

    try {
        cxxopts::Options options("MyProgram", "Description of my program");

        options.add_options()
          ("h,host", "Inference host", cxxopts::value<std::string>()->default_value(args.inference_host))
          ("p,port", "Inference port", cxxopts::value<int>()->default_value(std::to_string(args.inference_port)))
          ("t,threshold", "Voice threshold", cxxopts::value<float>()->default_value(std::to_string(args.voice_threshold)))
          ("d,duration", "Voice duration", cxxopts::value<float>()->default_value(std::to_string(args.voice_duration)))
          ("s,sample_rate", "Voice sample rate", cxxopts::value<float>()->default_value(std::to_string(args.voice_sample_rate)))
          ("b,buffer", "Voice frames buffer", cxxopts::value<int>()->default_value(std::to_string(args.voice_frames_buffer)))
          ("help", "Print help");

        auto result = options.parse(argc, argv);

        if (result.count("help")) {
            std::cout << options.help() << std::endl;
            exit(0);
        }

        args.inference_host = result["host"].as<std::string>();
        args.inference_port = result["port"].as<int>();
        args.voice_threshold = result["threshold"].as<float>();
        args.voice_duration = result["duration"].as<float>();
        args.voice_sample_rate = result["sample_rate"].as<float>();
        args.voice_frames_buffer = result["buffer"].as<int>();

    } catch (const cxxopts::exceptions::exception& e) {
        std::cerr << "Error parsing options: " << e.what() << std::endl;
        exit(1);
    }

    return args;
}