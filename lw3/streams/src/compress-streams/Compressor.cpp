#include "Compressor.h"

#include <ios>

std::vector<uint8_t> Compressor::RLECompress(const uint8_t* data, const size_t size)
{
	std::vector<uint8_t> result;
	for (size_t i = 0; i < size;)
	{
		uint8_t value = data[i];
		size_t count = 1;
		while (i + count < size && data[i + count] == value && count < 255)
		{
			++count;
		}

		result.push_back(static_cast<uint8_t>(count));
		result.push_back(value);
		i += count;
	}

	return result;
}

std::vector<uint8_t> Compressor::RLEDecompress(const uint8_t* data, const size_t size)
{
	std::vector<uint8_t> result;
	for (size_t i = 0; i < size; i += 2)
	{
		if (i + 1 >= size)
		{
			throw std::ios_base::failure("Invalid RLE data");
		}
		const uint8_t count = data[i];
		uint8_t value = data[i + 1];
		result.insert(result.end(), count, value);
	}

	return result;
}