#include "CanvasDrawable.h"

shape_drawing_lib::Triangle::Triangle(const Point& vertex1, const Point& vertex2, const Point& vertex3, const uint32_t color)
	: m_vertex1(vertex1)
	, m_vertex2(vertex2)
	, m_vertex3(vertex3)
	, m_color(color)
{
}
void shape_drawing_lib::Triangle::Draw(graphics_lib::ICanvas& canvas) const
{
	canvas.SetColor(m_color);
	canvas.MoveTo(m_vertex1.x, m_vertex1.y);
	canvas.LineTo(m_vertex2.x, m_vertex2.y);
	canvas.LineTo(m_vertex3.x, m_vertex3.y);
	canvas.LineTo(m_vertex1.x, m_vertex1.y);
}

shape_drawing_lib::Rectangle::Rectangle(const Point& leftTop, int const width, int const height, uint32_t const color)
	: m_leftTop(leftTop)
	, m_width(width)
	, m_height(height)
	, m_color(color)
{
}

void shape_drawing_lib::Rectangle::Draw(graphics_lib::ICanvas& canvas) const
{
	canvas.SetColor(m_color);
	canvas.MoveTo(m_leftTop.x, m_leftTop.y);
	canvas.LineTo(m_leftTop.x + m_width, m_leftTop.y);
	canvas.LineTo(m_leftTop.x + m_width, m_leftTop.y + m_height);
	canvas.LineTo(m_leftTop.x, m_leftTop.y + m_height);
	canvas.LineTo(m_leftTop.x, m_leftTop.y);
}

shape_drawing_lib::CanvasPainter::CanvasPainter(graphics_lib::ICanvas& canvas)
	: m_canvas(canvas)
{
}

void shape_drawing_lib::CanvasPainter::Draw(const ICanvasDrawable& drawable) const
{
	drawable.Draw(m_canvas);
}