#include "EncryptingOutputStream.h"

EncryptingOutputStream::EncryptingOutputStream(IOutputPtr&& stream, const uint32_t key)
	: OutputStreamDecorator(std::move(stream))
	, m_encTable(SubstitutionTable::GenerateEncryptionTable(key))
{
}

void EncryptingOutputStream::WriteByte(const uint8_t data)
{
	m_stream->WriteByte(m_encTable[data]);
}

void EncryptingOutputStream::WriteBlock(const void* srcData, std::streamsize size)
{
	const auto buffer = static_cast<const uint8_t*>(srcData);
	for (std::streamsize i = 0; i < size; ++i)
	{
		WriteByte(buffer[i]);
	}
}