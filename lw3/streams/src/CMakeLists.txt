add_executable(
        streams
        file-streams/FileInputStream.cpp
        file-streams/FileOutputStream.cpp
        memory-streams/MemoryInputStream.cpp
        memory-streams/MemoryOutputStream.cpp
        crypt-streams/SubstitutionTable.cpp
        crypt-streams/EncryptingOutputStream.cpp
        crypt-streams/DecryptingInputStream.cpp
        compress-streams/Compressor.cpp
        compress-streams/CompressingOutputStream.cpp
        compress-streams/DecompressingInputStream.cpp
        main.cpp
)

add_library(
        streams_lib STATIC
        file-streams/FileInputStream.cpp
        file-streams/FileOutputStream.cpp
        memory-streams/MemoryInputStream.cpp
        memory-streams/MemoryOutputStream.cpp
        crypt-streams/SubstitutionTable.cpp
        crypt-streams/EncryptingOutputStream.cpp
        crypt-streams/DecryptingInputStream.cpp
        compress-streams/Compressor.cpp
        compress-streams/CompressingOutputStream.cpp
        compress-streams/DecompressingInputStream.cpp
)

target_include_directories(streams_lib PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})