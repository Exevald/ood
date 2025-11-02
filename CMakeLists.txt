cmake_minimum_required(VERSION 3.27)
project(ood)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

find_package(GTest REQUIRED)

add_subdirectory(lw3)
add_subdirectory(lw6)