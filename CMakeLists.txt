cmake_minimum_required(VERSION 3.10)

project(scc)

set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -Wextra")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -Wextra")
set(CMAKE_CXX_STANDARD 17)

set(CMAKE_EXE_LINKER_FLAGS "-L/usr/lib64 -static")

# THIS MUST BE DEFINED TO EXT={ ".c", ".cc", ... }
# This default to {  }
# scc is intended to be easy to hack this helps with file extention checking
add_compile_definitions(EXT={"b"})

file(GLOB SOURCE_FILES 
    scc/*.c*
)

add_executable(scc ${SOURCE_FILES})
