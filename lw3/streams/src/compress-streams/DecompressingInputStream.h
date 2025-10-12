#pragma once

#include "../StreamDecorator.h"

#include <vector>

class DecompressingInputStream final : public InputStreamDecorator
{
public:
	explicit DecompressingInputStream(std::unique_ptr<IInputDataStream> stream);

	[[nodiscard]] bool IsEOF() const override;
	uint8_t ReadByte() override;
	std::streamsize ReadBlock(void* dstBuffer, std::streamsize size) override;

private:
	std::vector<uint8_t> m_decompressed;
	size_t m_pos;
};