include(GoogleTest)

add_executable(
        web-editor-model-tests
        Document_tests.cpp
        Frame_tests.cpp
)

target_link_libraries(web-editor-model-tests PRIVATE GTest::GTest GTest::gtest_main web-editor-model-lib)
gtest_discover_tests(web-editor-model-tests)