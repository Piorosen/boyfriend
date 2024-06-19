cmake_minimum_required(VERSION 3.10)
project(test)

set(CMAKE_CXX_STANDARD 11)

include_directories(./include)

find_package(portaudio CONFIG REQUIRED)
find_package(SndFile CONFIG REQUIRED)
find_package(spdlog CONFIG REQUIRED)

#add_library(core)
add_executable(test record.cpp)

target_link_libraries(test PRIVATE portaudio_static)
target_link_libraries(test PRIVATE SndFile::sndfile)
target_link_libraries(test PRIVATE spdlog::spdlog)

