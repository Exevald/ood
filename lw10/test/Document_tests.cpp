#include "domain/Document.h"
#include "domain/Primitive.h"

#include <gtest/gtest.h>

using namespace ::testing;

class DocumentTest : public Test
{
protected:
	static std::shared_ptr<Primitive> CreateRect(const std::string& id)
	{
		return std::make_shared<Primitive>(id, ShapeType::Rectangle, Frame{ 0, 0, 100, 100 });
	}
};

TEST_F(DocumentTest, AddsShapeToDocument)
{
	Document doc;
	const auto shape = CreateRect("rect1");
	doc.AddShape(shape);
	EXPECT_EQ(doc.GetShapes().size(), 1);
	EXPECT_EQ(doc.GetShape("rect1"), shape);
}

TEST_F(DocumentTest, RemovesShapeById)
{
	Document doc;
	const auto shape = CreateRect("rect1");
	doc.AddShape(shape);
	doc.RemoveShape("rect1");
	EXPECT_EQ(doc.GetShapes().size(), 0);
	EXPECT_EQ(doc.GetShape("rect1"), nullptr);
}

TEST_F(DocumentTest, RemovesAllShapes)
{
	Document doc;
	doc.AddShape(CreateRect("r1"));
	doc.AddShape(CreateRect("r2"));
	doc.Clear();
	EXPECT_TRUE(doc.GetShapes().empty());
}

TEST_F(DocumentTest, ReturnsNullIfNotFound)
{
	const Document doc;
	EXPECT_EQ(doc.GetShape("nonexistent"), nullptr);
}