#pragma once

#include <cstdint>
#include <vector>

namespace Compressor
{
std::vector<uint8_t> RLECompress(const uint8_t* data, size_t size);
std::vector<uint8_t> RLEDecompress(const uint8_t* data, size_t size);
} // namespace Compressor