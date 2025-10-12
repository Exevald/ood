#include "MemoryOutputStream.h"

void MemoryOutputStream::WriteByte(uint8_t data)
{
	m_data.push_back(data);
}

void MemoryOutputStream::WriteBlock(const void* srcData, std::streamsize size)
{
	const auto* pos = static_cast<const uint8_t*>(srcData);
	m_data.insert(m_data.end(), pos, pos + size);
}

void MemoryOutputStream::Close()
{
}