cmake_minimum_required(VERSION 3.19)

project(Engine)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_FLAGS "/Zc:__cplusplus")
add_compile_options(/EHa)

include_directories(../../Package/boost_1_74_0)
include_directories(../../Package/flatbuffers-1.12.0/include)
include_directories(../../Source)
link_directories(../../Package/boost_1_74_0/stage/lib)

file(GLOB_RECURSE Src "*.cpp" "*.h")

add_library(Engine ${Src})
