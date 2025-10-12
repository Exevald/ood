#include "MemoryInputStream.h"

#include <cstring>

MemoryInputStream::MemoryInputStream(std::vector<uint8_t> data)
	: m_data(std::move(data))
	, m_pos(0)
{
}

bool MemoryInputStream::IsEOF() const
{
	return m_pos >= m_data.size();
}

uint8_t MemoryInputStream::ReadByte()
{
	if (IsEOF())
	{
		throw std::ios_base::failure("EOF");
	}

	return m_data[m_pos++];
}

std::streamsize MemoryInputStream::ReadBlock(void* dstBuffer, const std::streamsize size)
{
	const std::streamsize toRead = std::min(size, static_cast<std::streamsize>(m_data.size() - m_pos));
	std::memcpy(dstBuffer, m_data.data() + m_pos, toRead);
	m_pos += toRead;

	return toRead;
}
