cmake_minimum_required(VERSION 3.19)

project(Server)

include_directories(../../Package/boost_1_74_0)
include_directories(../../Package/flatbuffers-1.12.0/include)
include_directories(../../Source)
link_directories(../../Package/boost_1_74_0/stage/lib)

file(GLOB_RECURSE ServerSrc "*.cpp" "*.h")

add_executable(Server ${ServerSrc})
target_link_libraries (Server Engine Message)
