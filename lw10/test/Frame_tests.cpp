#include "domain/Geometry.h"

#include <gtest/gtest.h>

TEST(FrameTest, PointInsideReturnsTrue)
{
	constexpr Frame frame{ 10, 10, 50, 50 };
	constexpr Point point{ 20, 30 };
	EXPECT_TRUE(frame.Contains(point));
}

TEST(FrameTest, PointOutsideReturnsFalse)
{
	constexpr Frame frame{ 10, 10, 50, 50 };
	constexpr Point point{ 100, 100 };
	EXPECT_FALSE(frame.Contains(point));
}

TEST(FrameTest, CombineTwoFramesCorrectly)
{
	constexpr Frame frame1{ 0, 0, 10, 10 };
	constexpr Frame frame2{ 5, 5, 10, 10 };
	const auto [x, y, width, height] = Frame::Unite(frame1, frame2);
	EXPECT_EQ(x, 0);
	EXPECT_EQ(y, 0);
	EXPECT_EQ(width, 15);
	EXPECT_EQ(height, 15);
}

TEST(FrameTest, UniteWithNegativeCoordinates)
{
	constexpr Frame frame1{ -10, -10, 10, 10 };
	constexpr Frame frame2{ 0, 0, 10, 10 };
	const auto [x, y, width, height] = Frame::Unite(frame1, frame2);
	EXPECT_EQ(x, -10);
	EXPECT_EQ(y, -10);
	EXPECT_EQ(width, 20);
	EXPECT_EQ(height, 20);
}