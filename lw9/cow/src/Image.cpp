#include "Image.h"

Image::Image(int width, int height, char fill)
	: m_width(width)
	, m_height(height)
{
	if (width <= 0 || height <= 0)
	{
		throw std::invalid_argument("Image dimensions must be positive");
	}

	auto tilesX = (width + TILE_SIZE - 1) / TILE_SIZE;
	auto tilesY = (height + TILE_SIZE - 1) / TILE_SIZE;

	m_tiles.resize(tilesY);
	for (int y = 0; y < tilesY; ++y)
	{
		m_tiles[y].resize(tilesX);
		for (int x = 0; x < tilesX; ++x)
		{
			m_tiles[y][x] = CoW<Tile>(fill);
		}
	}
}

char Image::GetPixel(int x, int y) const
{
	if (x < 0 || x >= m_width || y < 0 || y >= m_height)
	{
		return ' ';
	}

	int tileX = static_cast<int>(x / TILE_SIZE);
	int tileY = static_cast<int>(y / TILE_SIZE);
	int relX = static_cast<int>(x % TILE_SIZE);
	int relY = static_cast<int>(y % TILE_SIZE);

	return m_tiles[tileY][tileX]->GetPixel(relX, relY);
}

Size Image::GetSize() const
{
	return { m_width, m_height };
}

void Image::SetPixel(int x, int y, char color)
{
	if (x < 0 || x >= m_width || y < 0 || y >= m_height)
	{
		return;
	}

	int tileX = static_cast<int>(x / TILE_SIZE);
	int tileY = static_cast<int>(y / TILE_SIZE);
	int relX = static_cast<int>(x % TILE_SIZE);
	int relY = static_cast<int>(y % TILE_SIZE);

	m_tiles[tileY][tileX].Write()->SetPixel(relX, relY, color);
}
