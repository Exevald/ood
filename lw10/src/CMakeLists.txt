cmake_minimum_required(VERSION 3.15)
project(GrpcShapeEditor)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTOUIC ON)
set(CMAKE_AUTORCC ON)

find_package(PkgConfig REQUIRED)
pkg_check_modules(GRPC REQUIRED grpc++ grpc)
pkg_check_modules(PROTOBUF REQUIRED protobuf)
find_program(PROTOC_EXECUTABLE protoc REQUIRED)
find_program(GRPC_CPP_PLUGIN_EXECUTABLE grpc_cpp_plugin REQUIRED)

find_package(Qt5 COMPONENTS Core Widgets Gui REQUIRED)

set(PROTO_DIR "${CMAKE_CURRENT_SOURCE_DIR}/api")
set(PROTO_FILE "${PROTO_DIR}/webeditor.proto")
set(PROTO_SRC "${CMAKE_CURRENT_BINARY_DIR}/webeditor.pb.cc")
set(PROTO_HDR "${CMAKE_CURRENT_BINARY_DIR}/webeditor.pb.h")
set(GRPC_SRC "${CMAKE_CURRENT_BINARY_DIR}/webeditor.grpc.pb.cc")
set(GRPC_HDR "${CMAKE_CURRENT_BINARY_DIR}/webeditor.grpc.pb.h")

add_custom_command(
        OUTPUT "${PROTO_SRC}" "${PROTO_HDR}" "${GRPC_SRC}" "${GRPC_HDR}"
        COMMAND ${PROTOC_EXECUTABLE}
        ARGS --grpc_out "${CMAKE_CURRENT_BINARY_DIR}"
        --cpp_out "${CMAKE_CURRENT_BINARY_DIR}"
        -I "${PROTO_DIR}"
        --plugin=protoc-gen-grpc=${GRPC_CPP_PLUGIN_EXECUTABLE}
        "webeditor.proto"
        DEPENDS "${PROTO_FILE}"
)

add_library(web-editor-api ${PROTO_SRC} ${GRPC_SRC} ${PROTO_HDR} ${GRPC_HDR})
target_include_directories(web-editor-api PUBLIC
        ${GRPC_INCLUDE_DIRS}
        ${PROTOBUF_INCLUDE_DIRS}
        ${CMAKE_CURRENT_BINARY_DIR}
)
target_link_libraries(web-editor-api PUBLIC ${GRPC_LIBRARIES} ${PROTOBUF_LIBRARIES})

add_executable(web-editor
        main.cpp
        app/HistoryManager.cpp
        domain/Document.cpp
        domain/ShapeGroup.cpp
        domain/Primitive.cpp
        server/server.cpp
        client/client.cpp
        viewModel/ViewModel.cpp
        viewModel/JsonRepository.cpp
        view/CanvasWidget.cpp
        view/MainWindow.cpp
)

add_library(
        web-editor-model-lib STATIC
        domain/Document.cpp
        domain/ShapeGroup.cpp
        domain/Primitive.cpp
        app/HistoryManager.cpp
)

target_include_directories(web-editor-model-lib PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})

target_include_directories(web-editor PRIVATE src)
target_link_libraries(web-editor
        web-editor-api
        Qt5::Widgets Qt5::Core Qt5::Gui
        ${GRPC_LIBRARIES} ${PROTOBUF_LIBRARIES}
        pthread
)