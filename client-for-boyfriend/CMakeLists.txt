cmake_minimum_required(VERSION 3.10)
project(main)

set(CMAKE_CXX_STANDARD 11)

include_directories(./include)
include_directories(src/audio/include)
link_directories(${CMAKE_SOURCE_DIR}/library)

find_package(portaudio CONFIG REQUIRED)
find_package(SndFile CONFIG REQUIRED)
find_package(spdlog CONFIG REQUIRED)

add_library(audio src/audio/src/audio.cpp
                        src/audio/src/internal/utils.cpp
                        src/audio/src/internal/audo_impl.cpp
                        src/audio/src/c_api.cpp
                        src/args/args.cpp
)

add_executable(main main.cpp)

target_link_libraries(audio PRIVATE portaudio_static)
target_link_libraries(audio PRIVATE SndFile::sndfile)
target_link_libraries(audio PRIVATE spdlog::spdlog)

target_link_libraries(main PUBLIC network)
target_link_libraries(main PUBLIC audio)
