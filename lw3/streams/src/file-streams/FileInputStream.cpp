#include "FileInputStream.h"

FileInputStream::FileInputStream(const std::string& filename)
	: m_file(filename, std::ios::binary)
{
	if (!m_file.is_open())
	{
		throw std::ios_base::failure("Cannot open file for reading");
	}
}

bool FileInputStream::IsEOF() const
{
	return m_file.eof();
}

uint8_t FileInputStream::ReadByte()
{
	if (IsEOF())
	{
		throw std::ios_base::failure("EOF");
	}
	uint8_t bytesRead;

	m_file.read(reinterpret_cast<char*>(&bytesRead), 1);
	if (m_file.gcount() != 1)
	{
		throw std::ios_base::failure("Read error");
	}

	return bytesRead;
}

std::streamsize FileInputStream::ReadBlock(void* dstBuffer, const std::streamsize size)
{
	m_file.read(static_cast<char*>(dstBuffer), size);
	if (m_file.bad())
	{
		throw std::ios_base::failure("Read error");
	}

	return m_file.gcount();
}