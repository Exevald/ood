#pragma once

#include "../IDataStream.h"

#include <fstream>

class FileOutputStream final : public IOutputDataStream
{
public:
	explicit FileOutputStream(const std::string& filename);

	void WriteByte(uint8_t data) override;
	void WriteBlock(const void* srcData, std::streamsize size) override;
	void Close() override;

private:
	std::ofstream m_file;
};