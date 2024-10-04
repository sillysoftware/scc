cmake_minimum_required(VERSION 3.10)
project(scc)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -Wextra")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra")
set(CMAKE_CXX_STANDARD 17)

set(CMAKE_EXE_LINKER_FLAGS "-L/usr/lib64 -static")

set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_SOURCE_DIR}/scc)

file(GLOB SRC_FILES
  scc/*.c
  scc/*.cc
  scc/*.cpp
)

add_executable(scc ${SRC_FILES})
