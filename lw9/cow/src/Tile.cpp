#include "Tile.h"

std::atomic<std::size_t> Tile::m_instanceCount{ 0 };

Tile::Tile(char fill)
{
	++m_instanceCount;
	for (auto& row : m_pixels)
	{
		for (char& c : row)
		{
			c = fill;
		}
	}
}

Tile::Tile(const Tile& other)
{
	++m_instanceCount;
	for (std::size_t i = 0; i < TILE_SIZE; ++i)
	{
		for (std::size_t j = 0; j < TILE_SIZE; ++j)
		{
			m_pixels[i][j] = other.m_pixels[i][j];
		}
	}
}

Tile& Tile::operator=(const Tile& other)
{
	if (this != &other)
	{
		for (std::size_t i = 0; i < TILE_SIZE; ++i)
		{
			for (std::size_t j = 0; j < TILE_SIZE; ++j)
			{
				m_pixels[i][j] = other.m_pixels[i][j];
			}
		}
	}
	return *this;
}

void Tile::SetPixel(int x, int y, char color)
{
	if (x >= 0 && x < static_cast<int>(TILE_SIZE) && y >= 0 && y < static_cast<int>(TILE_SIZE))
	{
		m_pixels[y][x] = color;
	}
}

char Tile::GetPixel(int x, int y) const
{
	if (x >= 0 && x < static_cast<int>(TILE_SIZE) && y >= 0 && y < static_cast<int>(TILE_SIZE))
	{
		return m_pixels[y][x];
	}
	return ' ';
}

std::size_t Tile::GetInstanceCount()
{
	return m_instanceCount.load();
}
