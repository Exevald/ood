#include "adapters/ModernGraphicsRendererClassAdapter.h"
#include "adapters/ModernGraphicsRendererObjectAdapter.h"
#include "libs/Canvas.h"
#include "libs/CanvasDrawable.h"
#include "libs/GraphicsRenderer.h"

#include <gtest/gtest.h>
#include <sstream>
#include <string>

TEST(ModernGraphicsRendererTest, BeginDrawEndDraw)
{
	std::ostringstream out;
	modern_graphics_lib::ModernGraphicsRenderer renderer(out);

	renderer.BeginDraw();
	renderer.EndDraw();

	const std::string expectedOut = "<draw>\n</draw>\n";
	EXPECT_EQ(out.str(), expectedOut);
}

TEST(ModernGraphicsRendererTest, DrawLine_ThrowsIfNotInDrawingState)
{
	std::ostringstream out;
	const modern_graphics_lib::ModernGraphicsRenderer renderer(out);
	const modern_graphics_lib::Point p1{ 0, 0 };
	const modern_graphics_lib::Point p2{ 10, 10 };
	const modern_graphics_lib::RGBAColor color{ 1.0f, 0.0f, 0.0f, 1.0f };

	EXPECT_THROW(renderer.DrawLine(p1, p2, color), std::logic_error);
}

TEST(ModernGraphicsRendererTest, DrawLine_ValidUsage)
{
	std::ostringstream out;
	modern_graphics_lib::ModernGraphicsRenderer renderer(out);

	renderer.BeginDraw();
	renderer.DrawLine({ 1, 2 }, { 3, 4 }, { 0.5f, 0.6f, 0.7f, 1.0f });
	renderer.EndDraw();

	const std::string expectedOut = "<draw>\n"
									"<line fromX=1 fromY=2 toX=3 toY=4/>\n"
									"  <color r=0.5 g=0.6 b=0.7 a=1 />\n"
									"</line>\n"
									"</draw>\n";

	EXPECT_EQ(out.str(), expectedOut);
}

TEST(ModernGraphicsRendererObjectAdapterTest, SetColorMoveToLineTo)
{
	std::ostringstream out;
	modern_graphics_lib::ModernGraphicsRenderer renderer(out);
	ModernGraphicsRendererObjectAdapter adapter(renderer);

	renderer.BeginDraw();
	adapter.SetColor(0xFF0000);
	adapter.MoveTo(10, 20);
	adapter.LineTo(30, 40);
	renderer.EndDraw();

	const std::string expectedOut = "<draw>\n"
									"<line fromX=10 fromY=20 toX=30 toY=40/>\n"
									"  <color r=1 g=0 b=0 a=1 />\n"
									"</line>\n"
									"</draw>\n";

	EXPECT_EQ(out.str(), expectedOut);
}

TEST(ModernGraphicsRendererClassAdapterTest, FullDrawingCycle)
{
	std::ostringstream out;
	ModernGraphicsRendererClassAdapter adapter(out);

	adapter.BeginDraw();
	adapter.SetColor(0x00FF00);
	adapter.MoveTo(0, 0);
	adapter.LineTo(100, 100);
	adapter.EndDraw();

	const std::string expectedOut = "<draw>\n"
									"<line fromX=0 fromY=0 toX=100 toY=100/>\n"
									"  <color r=0 g=1 b=0 a=1 />\n"
									"</line>\n"
									"</draw>\n";

	EXPECT_EQ(out.str(), expectedOut);
}

TEST(TriangleTest, DrawsCorrectly)
{
	std::ostringstream out;
	modern_graphics_lib::ModernGraphicsRenderer renderer(out);
	ModernGraphicsRendererObjectAdapter adapter(renderer);

	renderer.BeginDraw();
	shape_drawing_lib::Triangle triangle(
		{ 0, 0 }, { 10, 0 }, { 5, 10 }, 0x0000FF);
	shape_drawing_lib::CanvasPainter painter(adapter);
	painter.Draw(triangle);
	renderer.EndDraw();

	std::string output = out.str();
	EXPECT_NE(output.find("fromX=0 fromY=0 toX=10 toY=0"), std::string::npos);
	EXPECT_NE(output.find("fromX=10 fromY=0 toX=5 toY=10"), std::string::npos);
	EXPECT_NE(output.find("fromX=5 fromY=10 toX=0 toY=0"), std::string::npos);
	EXPECT_NE(output.find("r=0 g=0 b=1"), std::string::npos);
}

TEST(RectangleTest, DrawsCorrectly)
{
	std::ostringstream out;
	modern_graphics_lib::ModernGraphicsRenderer renderer(out);
	ModernGraphicsRendererObjectAdapter adapter(renderer);

	renderer.BeginDraw();
	shape_drawing_lib::Rectangle rect({ 10, 20 }, 50, 30, 0xFFFF00);
	shape_drawing_lib::CanvasPainter painter(adapter);
	painter.Draw(rect);
	renderer.EndDraw();

	std::string output = out.str();

	EXPECT_NE(output.find("fromX=10 fromY=20 toX=60 toY=20"), std::string::npos);
	EXPECT_NE(output.find("fromX=60 fromY=20 toX=60 toY=50"), std::string::npos);
	EXPECT_NE(output.find("fromX=60 fromY=50 toX=10 toY=50"), std::string::npos);
	EXPECT_NE(output.find("fromX=10 fromY=50 toX=10 toY=20"), std::string::npos);
	EXPECT_NE(output.find("r=1 g=1 b=0"), std::string::npos);
}

TEST(CanvasTest, CallsPrintToCout)
{
	graphics_lib::Canvas canvas;
	EXPECT_NO_THROW({
		canvas.SetColor(0x123456);
		canvas.MoveTo(1, 2);
		canvas.LineTo(3, 4);
	});
}

TEST(RGBAColorTest, ConstructsCorrectly)
{
	const modern_graphics_lib::RGBAColor color(0.1f, 0.2f, 0.3f, 0.4f);
	EXPECT_FLOAT_EQ(color.r, 0.1f);
	EXPECT_FLOAT_EQ(color.g, 0.2f);
	EXPECT_FLOAT_EQ(color.b, 0.3f);
	EXPECT_FLOAT_EQ(color.a, 0.4f);
}