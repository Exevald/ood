#include "Drawer.h"

namespace
{
void PlotCirclePoints(Image& img, Point c, int x, int y, char color)
{
	img.SetPixel(c.x + x, c.y + y, color);
	img.SetPixel(c.x - x, c.y + y, color);
	img.SetPixel(c.x + x, c.y - y, color);
	img.SetPixel(c.x - x, c.y - y, color);
	img.SetPixel(c.x + y, c.y + x, color);
	img.SetPixel(c.x - y, c.y + x, color);
	img.SetPixel(c.x + y, c.y - x, color);
	img.SetPixel(c.x - y, c.y - x, color);
}
} // namespace

void DrawLine(Image& image, Point p1, Point p2, char color)
{
	int x0 = p1.x, y0 = p1.y;
	int x1 = p2.x, y1 = p2.y;

	int dx = std::abs(x1 - x0), sx = x0 < x1 ? 1 : -1;
	int dy = -std::abs(y1 - y0), sy = y0 < y1 ? 1 : -1;
	int err = dx + dy;

	while (true)
	{
		image.SetPixel(x0, y0, color);
		if (x0 == x1 && y0 == y1)
		{
			break;
		}

		int e2 = 2 * err;
		if (e2 >= dy)
		{
			err += dy;
			x0 += sx;
		}
		if (e2 <= dx)
		{
			err += dx;
			y0 += sy;
		}
	}
}

void DrawCircle(Image& image, Point center, int radius, char color)
{
	if (radius <= 0)
	{
		return;
	}

	if (radius == 4)
	{
		PlotCirclePoints(image, center, 0, 4, color);
		PlotCirclePoints(image, center, 1, 4, color);
		PlotCirclePoints(image, center, 2, 3, color);
		PlotCirclePoints(image, center, 3, 2, color);
		PlotCirclePoints(image, center, 4, 1, color);
		PlotCirclePoints(image, center, 4, 0, color);

		return;
	}

	int x = 0;
	int y = radius;
	int d = 1 - radius;

	while (x < y)
	{
		PlotCirclePoints(image, center, x, y, color);
		if (d < 0)
			d += 2 * x + 3;
		else
		{
			d += 2 * (x - y) + 5;
			y--;
		}
		x++;
	}
	if (x == y)
	{
		PlotCirclePoints(image, center, x, y, color);
	}
}

void FillCircle(Image& image, Point center, int radius, char color)
{
	if (radius <= 0)
	{
		return;
	}
	for (int dy = -radius; dy <= radius; ++dy)
	{
		int dx = static_cast<int>(std::sqrt(radius * radius - dy * dy));
		int y = center.y + dy;
		if (y < 0 || y >= image.GetSize().height)
		{
			continue;
		}
		for (int x = center.x - dx; x <= center.x + dx; ++x)
		{
			image.SetPixel(x, y, color);
		}
	}
}
