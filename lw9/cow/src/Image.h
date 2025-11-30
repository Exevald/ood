#pragma once

#include "CoW.h"
#include "Tile.h"

#include <cmath>
#include <stdexcept>
#include <vector>

struct Point
{
	int x, y;
};

struct Size
{
	int width, height;
};

class Image
{
public:
	Image(int width, int height, char fill = ' ');

	void SetPixel(int x, int y, char color);

	[[nodiscard]] char GetPixel(int x, int y) const;
	[[nodiscard]] Size GetSize() const;

private:
	int m_width, m_height;
	std::vector<std::vector<CoW<Tile>>> m_tiles;
};