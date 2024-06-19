cmake_minimum_required(VERSION 3.10)
project(main)

set(CMAKE_CXX_STANDARD 11)

include_directories(./include)
include_directories(./src/include)

find_package(portaudio CONFIG REQUIRED)
find_package(SndFile CONFIG REQUIRED)
find_package(spdlog CONFIG REQUIRED)

#add_library(libaudio src/src/audio.cpp
#                     src/src/internal/utils.cpp)
add_executable(main main.cpp)

#target_link_libraries(libaudio PUBLIC portaudio_static)
#target_link_libraries(libaudio PUBLIC SndFile::sndfile)
#target_link_libraries(libaudio PUBLIC spdlog::spdlog)

target_link_libraries(main PUBLIC portaudio_static)
target_link_libraries(main PUBLIC SndFile::sndfile)
target_link_libraries(main PUBLIC spdlog::spdlog)
#target_link_libraries(main PUBLIC libaudio)
