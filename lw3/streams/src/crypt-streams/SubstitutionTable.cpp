#include "SubstitutionTable.h"

SubstitutionTable::Table SubstitutionTable::GenerateEncryptionTable(uint32_t key)
{
	Table table;
	for (uint16_t i = 0; i < 256; ++i)
	{
		table[i] = static_cast<uint8_t>(i);
	}
	std::mt19937 rng(key);
	std::shuffle(table.begin(), table.end(), rng);

	return table;
}
SubstitutionTable::Table SubstitutionTable::GenerateDecryptionTable(const Table& encTable)
{
	Table decTable;
	for (size_t i = 0; i < 256; ++i)
	{
		decTable[encTable[i]] = static_cast<uint8_t>(i);
	}

	return decTable;
}