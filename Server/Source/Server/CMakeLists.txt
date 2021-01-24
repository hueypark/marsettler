cmake_minimum_required(VERSION 3.19)

project(Server)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_FLAGS "/Zc:__cplusplus")
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/../Binary)
add_compile_options(/EHa)

include_directories(../../Package/boost_1_74_0)
include_directories(../../Package/flatbuffers-1.12.0/include)
include_directories(../../Package/mongo-cxx-driver-r3.6.2/stage/include/bsoncxx/v_noabi)
include_directories(../../Package/mongo-cxx-driver-r3.6.2/stage/include/mongocxx/v_noabi)
include_directories(../../Source)
link_directories(../../Package/boost_1_74_0/stage/lib)
link_directories(../../Package/mongo-c-driver-1.17.3/stage/lib)
link_directories(../../Package/mongo-cxx-driver-r3.6.2/stage/lib)

file(GLOB_RECURSE ServerSrc "*.cpp" "*.h")

add_executable(Server ${ServerSrc})
target_link_libraries (Server Engine Message bson-1.0.lib mongoc-1.0 bsoncxx mongocxx)
