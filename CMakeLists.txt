cmake_minimum_required(VERSION 3.10)
project(scc)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -Wextra")

# Define source files
file(GLOB SRC_FILES
  SCC/*.c
  SCC/*.cc
  SCC/*.cpp
)

# Add executable target
add_executable(scc ${SRC_FILES})
