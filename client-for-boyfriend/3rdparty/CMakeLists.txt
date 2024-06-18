cmake_minimum_required(VERSION 3.10)
project(test)

set(CMAKE_CXX_STANDARD 11)

find_package(libsoundio CONFIG REQUIRED)

add_executable(test main.cpp)

target_link_libraries(test PRIVATE libsoundio::libsoundio libsoundio::libsoundio_shared)
