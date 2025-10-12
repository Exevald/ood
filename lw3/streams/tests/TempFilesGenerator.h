#pragma once

#include <filesystem>
#include <fstream>
#include <random>
#include <string>

class TempFileGuard
{
public:
	explicit TempFileGuard(std::string path)
		: m_path(std::move(path))
	{
	}

	~TempFileGuard()
	{
		std::filesystem::remove(m_path);
	}

private:
	std::string m_path;
};

namespace TempFileGenerator
{
std::string CreateTempFile(const std::vector<uint8_t>& data);
} // namespace TempFileGenerator