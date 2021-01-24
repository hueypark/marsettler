cmake_minimum_required(VERSION 3.19)

project(Message)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_FLAGS "/Zc:__cplusplus")
add_compile_options(/EHa)

include_directories(../../Package/flatbuffers-1.12.0/include)
include_directories(../../Source)

file(GLOB_RECURSE Src "*.cpp" "*.h")

add_library(Message ${Src})
