cmake_minimum_required(VERSION 3.19)

project(Server)

include_directories(Package/boost_1_74_0)
link_directories(Package/boost_1_74_0/stage/lib)

add_executable(Server Main.cpp)
