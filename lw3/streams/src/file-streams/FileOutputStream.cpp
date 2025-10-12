#include "FileOutputStream.h"

FileOutputStream::FileOutputStream(const std::string& filename)
	: m_file(filename, std::ios::binary)
{
	if (!m_file.is_open())
	{
		throw std::ios_base::failure("Cannot open file for writing");
	}
}

void FileOutputStream::WriteByte(const uint8_t data)
{
	m_file.write(reinterpret_cast<const char*>(&data), 1);
	if (!m_file.good())
	{
		throw std::ios_base::failure("Write error");
	}
}

void FileOutputStream::WriteBlock(const void* srcData, const std::streamsize size)
{
	m_file.write(static_cast<const char*>(srcData), size);
	if (!m_file.good())
	{
		throw std::ios_base::failure("Write error");
	}
}

void FileOutputStream::Close()
{
	m_file.close();
}
