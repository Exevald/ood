#pragma once

#include <array>
#include <random>

namespace SubstitutionTable
{
using Table = std::array<uint8_t, 256>;

Table GenerateEncryptionTable(uint32_t key);
Table GenerateDecryptionTable(const Table& encTable);
} // namespace SubstitutionTable