#include "domain/Geometry.h"
#include "domain/Primitive.h"

#include <gtest/gtest.h>

TEST(PrimitiveTest, ConstructorInitializesMembers)
{
	constexpr Frame frame{ 1, 2, 3, 4 };
	constexpr Primitive primitive("id1", ShapeType::Ellipse, frame, 0xFF00FF00, "path.png");
	EXPECT_EQ(primitive.GetId(), "id1");
	EXPECT_EQ(primitive.GetType(), ShapeType::Ellipse);
	EXPECT_EQ(primitive.GetFrame().x, 1);
	EXPECT_EQ(primitive.GetFrame().y, 2);
	EXPECT_EQ(primitive.GetFrame().width, 3);
	EXPECT_EQ(primitive.GetFrame().height, 4);
	EXPECT_EQ(primitive.GetColor(), 0xFF00FF00U);
	EXPECT_EQ(primitive.GetImagePath(), "path.png");
}

TEST(PrimitiveTest, SetGetId)
{
	Primitive primitive("old", ShapeType::Rectangle, Frame{});
	primitive.SetId("new");
	EXPECT_EQ(primitive.GetId(), "new");
}

TEST(PrimitiveTest, SetGetFrame)
{
	Primitive primitive("id", ShapeType::Rectangle, Frame{ 0, 0, 10, 10 });
	primitive.SetFrame({ 5, 5, 20, 20 });
	EXPECT_EQ(primitive.GetFrame().x, 5);
	EXPECT_EQ(primitive.GetFrame().width, 20);
}

TEST(PrimitiveTest, MoveFrame)
{
	Primitive primitive("id", ShapeType::Rectangle, Frame{ 10, 10, 5, 5 });
	primitive.MoveFrame(3, -2);
	EXPECT_EQ(primitive.GetFrame().x, 13);
	EXPECT_EQ(primitive.GetFrame().y, 8);
}

TEST(PrimitiveTest, CloneCreatesDeepCopy)
{
	const Primitive original("orig", ShapeType::Triangle, Frame{ 1, 1, 2, 2 }, 0x12345678, "img.png");
	const auto copy = original.Clone();
	ASSERT_NE(copy, nullptr);
	EXPECT_EQ(copy->GetId(), "orig");
	EXPECT_EQ(copy->GetType(), ShapeType::Triangle);
	EXPECT_EQ(copy->GetFrame().width, 2);
	EXPECT_EQ(copy->GetColor(), 0x12345678U);
	EXPECT_EQ(dynamic_cast<Primitive*>(copy.get())->GetImagePath(), "img.png");
}