#include "DecompressingInputStream.h"
#include "Compressor.h"

#include <cstring>

DecompressingInputStream::DecompressingInputStream(std::unique_ptr<IInputDataStream> stream)
	: InputStreamDecorator(std::move(stream))
	, m_pos(0)
{
	std::vector<uint8_t> compressed;
	constexpr size_t BUF_SIZE = 4096;
	uint8_t buffer[BUF_SIZE];
	while (!m_stream->IsEOF())
	{
		const auto read = m_stream->ReadBlock(buffer, BUF_SIZE);
		compressed.insert(compressed.end(), buffer, buffer + read);
	}
	m_decompressed = Compressor::RLEDecompress(compressed.data(), compressed.size());
}

bool DecompressingInputStream::IsEOF() const
{
	return m_pos >= m_decompressed.size();
}

uint8_t DecompressingInputStream::ReadByte()
{
	if (IsEOF())
	{
		throw std::ios_base::failure("EOF");
	}

	return m_decompressed[m_pos++];
}

std::streamsize DecompressingInputStream::ReadBlock(void* dstBuffer, const std::streamsize size)
{
	const std::streamsize toRead = std::min(size, static_cast<std::streamsize>(m_decompressed.size() - m_pos));
	std::memcpy(dstBuffer, m_decompressed.data() + m_pos, toRead);
	m_pos += toRead;

	return toRead;
}