cmake_minimum_required(VERSION 3.19)

project(Message)

include_directories(../../Package/flatbuffers-1.12.0/include)

file(GLOB_RECURSE Src "*.cpp" "*.h")

add_library(Message ${Src})
