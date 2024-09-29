cmake_minimum_required(VERSION 3.10)
project(scc)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -Wextra")

# Add libscc directory for header files
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/libscc)

# Define source files
set(SRC_FILES SCC/main.c SCC/lexer.c SCC/generator.c)

# Add executable target
add_executable(scc ${SRC_FILES})
