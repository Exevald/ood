#pragma once

#include "../StreamDecorator.h"
#include "Compressor.h"

class CompressingOutputStream final : public OutputStreamDecorator
{
public:
	explicit CompressingOutputStream(std::unique_ptr<IOutputDataStream> stream);

	void WriteByte(uint8_t data) override;
	void WriteBlock(const void* srcData, std::streamsize size) override;
	void Close() override;

private:
	std::vector<uint8_t> m_buffer;
};