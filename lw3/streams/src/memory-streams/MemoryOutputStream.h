#pragma once

#include "../IDataStream.h"

#include <vector>

class MemoryOutputStream final : public IOutputDataStream
{
public:
	void WriteByte(uint8_t data) override;
	void WriteBlock(const void* srcData, std::streamsize size) override;
	void Close() override;
	[[nodiscard]] const std::vector<uint8_t>& GetData() const { return m_data; }

private:
	std::vector<uint8_t> m_data;
};