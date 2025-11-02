#pragma once

#include <iomanip>

namespace modern_graphics_lib
{

class Point
{
public:
	Point(const int x, const int y)
		: x(x)
		, y(y)
	{
	}
	int x;
	int y;
};

class RGBAColor
{
public:
	RGBAColor(float const r, float const g, float const b, float const a)
		: r(r)
		, g(g)
		, b(b)
		, a(a)
	{
	}
	float r, g, b, a;
};

class ModernGraphicsRenderer
{
public:
	explicit ModernGraphicsRenderer(std::ostream& strm);
	virtual ~ModernGraphicsRenderer();

	virtual void BeginDraw();
	void DrawLine(const Point& start, const Point& end, const RGBAColor& color) const;
	virtual void EndDraw();

private:
	std::ostream& m_out;
	bool m_drawing = false;
};
} // namespace modern_graphics_lib