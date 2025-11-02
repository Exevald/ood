include(GoogleTest)

add_executable(
        shapes-adapters-tests
        ShapesAdapters_test.cpp
)

target_link_libraries(shapes-adapters-tests PRIVATE GTest::GTest GTest::gtest_main shapes-adapters-lib)
gtest_discover_tests(shapes-adapters-tests)