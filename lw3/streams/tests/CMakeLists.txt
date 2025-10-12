include(GoogleTest)

add_executable(
        streams_test
        TempFilesGenerator.cpp
        Streams_test.cpp
)

target_link_libraries(streams_test PRIVATE GTest::GTest GTest::gtest_main streams_lib)
gtest_discover_tests(streams_test)