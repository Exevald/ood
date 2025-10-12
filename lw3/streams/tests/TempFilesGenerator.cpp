#include "TempFilesGenerator.h"

namespace
{
std::string GenerateRandomSuffix()
{
	static std::random_device rd;
	static std::mt19937 gen(rd());
	static std::uniform_int_distribution<> dis(0, 15);

	std::stringstream ss;
	ss << std::hex;
	for (int i = 0; i < 8; ++i)
	{
		ss << dis(gen);
	}

	return ss.str();
}
} // namespace

std::string TempFileGenerator::CreateTempFile(const std::vector<uint8_t>& data)
{
	std::string pattern = "test_XXXXXXXX.tmp";
	auto tmpDir = std::filesystem::temp_directory_path();

	for (int attempts = 0; attempts < 100; ++attempts)
	{
		std::string suffix = GenerateRandomSuffix();
		std::string filename = "test_" + suffix + ".tmp";
		std::filesystem::path path = tmpDir / filename;

		if (std::ofstream file(path, std::ios::binary); file.is_open())
		{
			file.write(reinterpret_cast<const char*>(data.data()), static_cast<long>(data.size()));
			file.close();
			return path.string();
		}
	}

	throw std::runtime_error("Could not create temporary file");
}