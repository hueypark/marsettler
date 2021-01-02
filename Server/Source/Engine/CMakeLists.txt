cmake_minimum_required(VERSION 3.19)

project(Engine)

include_directories(../../Source)

file(GLOB_RECURSE Src "*.cpp" "*.h")

add_library(Engine ${Src})

