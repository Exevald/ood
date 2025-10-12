#include "DecryptingInputStream.h"

DecryptingInputStream::DecryptingInputStream(std::unique_ptr<IInputDataStream> stream, const uint32_t key)
	: InputStreamDecorator(std::move(*this))
	, m_stream(std::move(stream))
	, m_decTable(SubstitutionTable::GenerateDecryptionTable(SubstitutionTable::GenerateEncryptionTable(key)))
{
}

bool DecryptingInputStream::IsEOF() const
{
	return m_stream->IsEOF();
}

uint8_t DecryptingInputStream::ReadByte()
{
	return m_decTable[m_stream->ReadByte()];
}

std::streamsize DecryptingInputStream::ReadBlock(void* dstBuffer, const std::streamsize size)
{
	const auto block = static_cast<uint8_t*>(dstBuffer);
	for (std::streamsize i = 0; i < size; ++i)
	{
		block[i] = ReadByte();
	}

	return size;
}