cmake_minimum_required(VERSION 3.10)

project(scc)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -Wextra")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra")
set(CMAKE_CXX_STANDARD 17)

set(CMAKE_EXE_LINKER_FLAGS "-L/usr/lib64 -static")

include_directories(scc/basic)

set(C_SOURCE_FILES
    scc/basic/basic.c
    scc/error.c
)

set(CXX_SOURCE_FILES
    scc/main.cc
)

add_executable(scc ${C_SOURCE_FILES} ${CXX_SOURCE_FILES})
