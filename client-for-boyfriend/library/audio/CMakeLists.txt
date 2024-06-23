cmake_minimum_required(VERSION 3.10)
project(main)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED True)
set(BUILD_SHARED_LIBS OFF)
#set(CMAKE_EXE_LINKER_FLAGS "-static")

include_directories(./include)
include_directories(src/audio/include)


link_directories(${CMAKE_SOURCE_DIR}/library)

find_package(portaudio CONFIG REQUIRED)
find_package(SndFile CONFIG REQUIRED)
find_package(spdlog CONFIG REQUIRED)
find_package(cxxopts CONFIG REQUIRED)

add_library(audio SHARED src/audio/src/audio.cpp
                        src/audio/src/internal/utils.cpp
                        src/audio/src/internal/audo_impl.cpp
                        src/audio/src/c_api.cpp
)
add_library(args STATIC src/args/args.cpp)

add_executable(main main.cpp)

if(${CMAKE_SYSTEM_NAME} STREQUAL "Linux" OR ${CMAKE_SYSTEM_NAME} STREQUAL "Darwin")
    target_link_libraries(audio PRIVATE portaudio_static)
else()
    message( FATAL_ERROR "You can not do this at all, CMake will exit." )
    target_link_libraries(audio PRIVATE portaudio)
endif()

target_link_libraries(audio PRIVATE SndFile::sndfile)
target_link_libraries(audio PRIVATE spdlog::spdlog)
target_link_libraries(audio PRIVATE spdlog::spdlog)

target_link_libraries(args PRIVATE cxxopts::cxxopts)
target_link_libraries(main PUBLIC args)

target_link_libraries(main PUBLIC audio)
