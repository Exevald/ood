#pragma once

#include "../StreamDecorator.h"
#include "SubstitutionTable.h"

class DecryptingInputStream final : public InputStreamDecorator
{
public:
	DecryptingInputStream(std::unique_ptr<IInputDataStream> stream, uint32_t key);

	[[nodiscard]] bool IsEOF() const override;
	uint8_t ReadByte() override;

	std::streamsize ReadBlock(void* dstBuffer, std::streamsize size) override;

private:
	std::unique_ptr<IInputDataStream> m_stream;
	SubstitutionTable::Table m_decTable;
};