#pragma once

struct Point
{
	double x, y;
};

struct Size
{
	double width, height;
};

struct Frame
{
	double x, y, width, height;

	[[nodiscard]] bool Contains(const Point& p) const
	{
		return p.x >= x && p.x <= x + width && p.y >= y && p.y <= y + height;
	}

	static Frame Unite(const Frame& a, const Frame& b)
	{
		const double minX = (a.x < b.x) ? a.x : b.x;
		const double minY = (a.y < b.y) ? a.y : b.y;
		const double maxX = (a.x + a.width > b.x + b.width) ? a.x + a.width : b.x + b.width;
		const double maxY = (a.y + a.height > b.y + b.height) ? a.y + a.height : b.y + b.height;
		return {
			.x = minX,
			.y = minY,
			.width = maxX - minX,
			.height = maxY - minY
		};
	}
};

enum class ShapeType
{
	Rectangle,
	Triangle,
	Ellipse,
	Image,
	Group
};
