#include "GraphicsRenderer.h"

modern_graphics_lib::ModernGraphicsRenderer::ModernGraphicsRenderer(std::ostream& strm)
	: m_out(strm)
{
}

modern_graphics_lib::ModernGraphicsRenderer::~ModernGraphicsRenderer()
{
	if (m_drawing)
	{
		ModernGraphicsRenderer::EndDraw();
	}
}

void modern_graphics_lib::ModernGraphicsRenderer::BeginDraw()
{
	if (m_drawing)
	{
		throw std::logic_error("Drawing has already begun");
	}
	m_out << "<draw>" << std::endl;
	m_drawing = true;
}

void modern_graphics_lib::ModernGraphicsRenderer::DrawLine(const Point& start, const Point& end, const RGBAColor& color) const
{
	if (!m_drawing)
	{
		throw std::logic_error("DrawLine is allowed between BeginDraw()/EndDraw() only");
	}
	m_out << "<line fromX=" << start.x << " fromY=" << start.y
		  << " toX=" << end.x << " toY=" << end.y << "/>" << std::endl;
	m_out << std::setprecision(2) << "  <color r=" << color.r << " g=" << color.g << " b=" << color.b << " a=" << color.a << " />" << std::endl;
	m_out << "</line>" << std::endl;
}

void modern_graphics_lib::ModernGraphicsRenderer::EndDraw()
{
	if (!m_drawing)
	{
		throw std::logic_error("Drawing has not been started");
	}
	m_out << "</draw>" << std::endl;
	m_drawing = false;
}