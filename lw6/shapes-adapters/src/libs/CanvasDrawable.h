#pragma once

#include "Canvas.h"

namespace shape_drawing_lib
{
struct Point
{
	int x;
	int y;
};

class ICanvasDrawable
{
public:
	virtual void Draw(graphics_lib::ICanvas& canvas) const = 0;
	virtual ~ICanvasDrawable() = default;
};

class Triangle final : public ICanvasDrawable
{
public:
	Triangle(const Point& vertex1, const Point& vertex2, const Point& vertex3, uint32_t color = 0x000000);

	void Draw(graphics_lib::ICanvas& canvas) const override;

private:
	Point m_vertex1, m_vertex2, m_vertex3;
	uint32_t m_color;
};

class Rectangle final : public ICanvasDrawable
{
public:
	Rectangle(const Point& leftTop, int width, int height, uint32_t color = 0x000000);

	void Draw(graphics_lib::ICanvas& canvas) const override;

private:
	Point m_leftTop;
	int m_width, m_height;
	uint32_t m_color;
};

class CanvasPainter
{
public:
	explicit CanvasPainter(graphics_lib::ICanvas& canvas);

	void Draw(const ICanvasDrawable& drawable) const;

private:
	graphics_lib::ICanvas& m_canvas;
};
} // namespace shape_drawing_lib