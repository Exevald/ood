#pragma once

#include <atomic>
#include <cstddef>

constexpr std::size_t TILE_SIZE = 8;

class Tile
{
public:
	explicit Tile(char fill = ' ');
	Tile(const Tile& other);

	Tile& operator=(const Tile& other);

	void SetPixel(int x, int y, char color);
	[[nodiscard]] char GetPixel(int x, int y) const;
	static std::size_t GetInstanceCount();

private:
	char m_pixels[TILE_SIZE][TILE_SIZE]{};
	static std::atomic<std::size_t> m_instanceCount;
};