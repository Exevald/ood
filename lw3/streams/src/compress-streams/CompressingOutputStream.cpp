#include "CompressingOutputStream.h"

CompressingOutputStream::CompressingOutputStream(std::unique_ptr<IOutputDataStream> stream)
	: OutputStreamDecorator(std::move(stream))
{
}

void CompressingOutputStream::WriteByte(const uint8_t data)
{
	m_buffer.push_back(data);
}

void CompressingOutputStream::WriteBlock(const void* srcData, std::streamsize size)
{
	const auto pos = static_cast<const uint8_t*>(srcData);
	m_buffer.insert(m_buffer.end(), pos, pos + size);
}

void CompressingOutputStream::Close()
{
	if (!m_buffer.empty())
	{
		const auto compressed = Compressor::RLECompress(m_buffer.data(), m_buffer.size());
		m_stream->WriteBlock(compressed.data(), static_cast<std::streamsize>(compressed.size()));
	}
	m_stream->Close();
}