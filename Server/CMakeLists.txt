cmake_minimum_required(VERSION 3.19)

project(Server)

set_property(DIRECTORY PROPERTY VS_STARTUP_PROJECT Server)

include_directories(Package/boost_1_74_0)
link_directories(Package/boost_1_74_0/stage/lib)

file(GLOB ServerSrc "Source/*.cpp", "Source/*.h")

add_executable(Server ${ServerSrc})
