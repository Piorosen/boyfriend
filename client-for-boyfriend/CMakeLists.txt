cmake_minimum_required(VERSION 3.10)
project(main)

set(CMAKE_CXX_STANDARD 11)

include_directories(./include)
include_directories(./src/include)
link_directories(${CMAKE_SOURCE_DIR}/library)

find_package(portaudio CONFIG REQUIRED)
find_package(SndFile CONFIG REQUIRED)
find_package(spdlog CONFIG REQUIRED)

add_library(libaudio src/src/audio.cpp
                     src/src/internal/utils.cpp
                     src/src/internal/audo_impl.cpp
                     src/src/c_api.cpp
)

add_executable(main main.cpp)

target_link_libraries(libaudio PRIVATE portaudio_static)
target_link_libraries(libaudio PRIVATE SndFile::sndfile)
target_link_libraries(libaudio PRIVATE spdlog::spdlog)

target_link_libraries(main PUBLIC network)
target_link_libraries(main PUBLIC libaudio)
