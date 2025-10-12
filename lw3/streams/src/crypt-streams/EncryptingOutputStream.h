#pragma once

#include "../StreamDecorator.h"
#include "SubstitutionTable.h"

class EncryptingOutputStream final : public OutputStreamDecorator
{
public:
	EncryptingOutputStream(IOutputPtr&& stream, uint32_t key);

	void WriteByte(uint8_t data) override;
	void WriteBlock(const void* srcData, std::streamsize size) override;

private:
	SubstitutionTable::Table m_encTable{};
};