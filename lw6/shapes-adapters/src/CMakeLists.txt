add_executable(
        shapes-adapters
        ./libs/Canvas.cpp
        ./libs/CanvasDrawable.cpp
        ./libs/GraphicsRenderer.cpp
        ./adapters/ModernGraphicsRendererClassAdapter.cpp
        ./adapters/ModernGraphicsRendererObjectAdapter.cpp
        main.cpp
)

add_library(
        shapes-adapters-lib
        ./libs/Canvas.cpp
        ./libs/CanvasDrawable.cpp
        ./libs/GraphicsRenderer.cpp
        ./adapters/ModernGraphicsRendererClassAdapter.cpp
        ./adapters/ModernGraphicsRendererObjectAdapter.cpp
        main.cpp
)

target_include_directories(shapes-adapters-lib PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})