#include "domain/Primitive.h"
#include "domain/ShapeGroup.h"

#include <gtest/gtest.h>

using namespace ::testing;

class ShapeGroupTest : public Test
{
protected:
	static std::shared_ptr<Primitive> MakeRect(const std::string& id, const double x, const double y, const double w, const double h)
	{
		return std::make_shared<Primitive>(id, ShapeType::Rectangle, Frame{ x, y, w, h });
	}
};

TEST_F(ShapeGroupTest, ConstructorInitializesEmptyGroup)
{
	const ShapeGroup group("g1");
	EXPECT_EQ(group.GetId(), "g1");
	EXPECT_EQ(group.GetType(), ShapeType::Group);
	EXPECT_TRUE(group.GetChildren().empty());
	EXPECT_EQ(group.GetFrame().width, 0);
}

TEST_F(ShapeGroupTest, AddChildUpdatesFrame)
{
	ShapeGroup group("g1");
	const auto r1 = MakeRect("r1", 0, 0, 10, 10);
	const auto r2 = MakeRect("r2", 20, 20, 5, 5);
	group.Add(r1);
	EXPECT_EQ(group.GetFrame().width, 10);
	group.Add(r2);
	EXPECT_EQ(group.GetFrame().x, 0);
	EXPECT_EQ(group.GetFrame().y, 0);
	EXPECT_EQ(group.GetFrame().width, 25);
	EXPECT_EQ(group.GetFrame().height, 25);
}

TEST_F(ShapeGroupTest, RemoveChildUpdatesFrame)
{
	ShapeGroup group("g1");
	const auto r1 = MakeRect("r1", 0, 0, 10, 10);
	const auto r2 = MakeRect("r2", 20, 20, 5, 5);
	group.Add(r1);
	group.Add(r2);
	group.Remove("r2");
	EXPECT_EQ(group.GetChildren().size(), 1);
	EXPECT_EQ(group.GetFrame().width, 10);
}

TEST_F(ShapeGroupTest, MoveFrameMovesAllChildren)
{
	ShapeGroup group("g1");
	const auto r = MakeRect("r1", 5, 5, 10, 10);
	group.Add(r);
	group.MoveFrame(3, 4);
	EXPECT_EQ(r->GetFrame().x, 8);
	EXPECT_EQ(r->GetFrame().y, 9);
	EXPECT_EQ(group.GetFrame().x, 8);
	EXPECT_EQ(group.GetFrame().y, 9);
}

TEST_F(ShapeGroupTest, SetFrameScalesChildrenProportionally)
{
	ShapeGroup group("g1");
	const auto r1 = MakeRect("r1", 0, 0, 10, 10);
	const auto r2 = MakeRect("r2", 20, 0, 10, 10);
	group.Add(r1);
	group.Add(r2);

	constexpr Frame newFrame{ 0, 0, 60, 20 };
	group.SetFrame(newFrame);

	EXPECT_NEAR(r1->GetFrame().x, 0, 1e-5);
	EXPECT_NEAR(r1->GetFrame().width, 20, 1e-5);
	EXPECT_NEAR(r2->GetFrame().x, 40, 1e-5);
	EXPECT_NEAR(r2->GetFrame().width, 20, 1e-5);
	EXPECT_NEAR(group.GetFrame().width, 60, 1e-5);
}

TEST_F(ShapeGroupTest, CloneCreatesDeepCopyOfChildren)
{
	ShapeGroup original("group");
	original.Add(MakeRect("r1", 0, 0, 10, 10));
	const auto cloned = original.Clone();
	ASSERT_NE(cloned, nullptr);
	auto& children = cloned->GetChildren();
	ASSERT_EQ(children.size(), 1);
	EXPECT_EQ(children[0]->GetId(), "r1");
	EXPECT_NE(children[0], original.GetChildren()[0]);
}

TEST_F(ShapeGroupTest, SetColorAppliesToAllChildren)
{
	ShapeGroup group("g1");
	const auto r1 = MakeRect("r1", 0, 0, 10, 10);
	const auto r2 = MakeRect("r2", 0, 0, 10, 10);
	group.Add(r1);
	group.Add(r2);
	group.SetColor(0x11223344);
	EXPECT_EQ(r1->GetColor(), 0x11223344U);
	EXPECT_EQ(r2->GetColor(), 0x11223344U);
}

TEST_F(ShapeGroupTest, GetColorReturnsFirstChildColor)
{
	ShapeGroup group("g1");
	const auto r = MakeRect("r1", 0, 0, 10, 10);
	r->SetColor(0xAABBCCDD);
	group.Add(r);
	EXPECT_EQ(group.GetColor(), 0xAABBCCDDU);
}

TEST_F(ShapeGroupTest, GetColorEmptyGroup_ReturnsBlack)
{
	const ShapeGroup group("empty");
	EXPECT_EQ(group.GetColor(), 0xFF000000U);
}