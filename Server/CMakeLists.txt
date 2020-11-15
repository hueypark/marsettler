cmake_minimum_required(VERSION 3.19)

set_property(DIRECTORY PROPERTY VS_STARTUP_PROJECT Server)

add_subdirectory(Source/Message)
add_subdirectory(Source/Server)
