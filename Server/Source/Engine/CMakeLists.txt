cmake_minimum_required(VERSION 3.19)

project(Engine)

include_directories(../../Package/boost_1_74_0)
include_directories(../../Package/flatbuffers-1.12.0/include)
include_directories(../../Source)
link_directories(../../Package/boost_1_74_0/stage/lib)

file(GLOB_RECURSE Src "*.cpp" "*.h")

add_library(Engine ${Src})

