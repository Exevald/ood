#include "compress-streams/CompressingOutputStream.h"
#include "compress-streams/DecompressingInputStream.h"
#include "crypt-streams/DecryptingInputStream.h"
#include "crypt-streams/EncryptingOutputStream.h"
#include "file-streams/FileInputStream.h"
#include "file-streams/FileOutputStream.h"

#include <fstream>
#include <iostream>
#include <optional>
#include <string>
#include <vector>

#include <memory>

class IOutputDataStream;
class IInputDataStream;

struct Args
{
	std::string inputFileName;
	std::string outputFileName;
	std::vector<std::string> options;
};

std::optional<Args> ParseArgs(const int argc, char* argv[])
{
	if (argc < 3)
	{
		return std::nullopt;
	}
	Args args;
	for (int i = 1; i < argc; i++)
	{
		if (i == argc - 2)
		{
			args.inputFileName = argv[i];
		}
		else if (i == argc - 1)
		{
			args.outputFileName = argv[i];
		}
		else
		{
			args.options.emplace_back(argv[i]);
		}
	}

	return args;
}

bool ProcessArgError(const std::optional<Args>& args)
{
	if (!args.has_value())
	{
		std::cout << "Invalid arguments count\n";
		std::cout << "Usage: ./transform [options] <input file> <output file>\n";
		return false;
	}
	return true;
}

std::unique_ptr<IInputDataStream> MakeInputStream(std::string const& filename, std::vector<std::string> const& options)
{
	std::unique_ptr<IInputDataStream> strm = std::make_unique<FileInputStream>(filename);
	for (auto it = options.rbegin(); it != options.rend(); ++it)
	{
		if (*it == "--decrypt")
		{
			strm = std::make_unique<DecryptingInputStream>(std::move(strm), std::stoi(*(it - 1)));
		}
		else if (*it == "--decompress")
		{
			strm = std::make_unique<DecompressingInputStream>(std::move(strm));
		}
	}

	return strm;
}

std::unique_ptr<IOutputDataStream> MakeOutputStream(std::string const& filename, std::vector<std::string> const& options)
{
	std::unique_ptr<IOutputDataStream> strm = std::make_unique<FileOutputStream>(filename);
	for (size_t i = 0; i < options.size(); i++)
	{
		if (options[i] == "--encrypt")
		{
			strm = std::make_unique<EncryptingOutputStream>(std::move(strm), stoi(options[i + 1]));
		}
		else if (options[i] == "--compress")
		{
			strm = std::make_unique<CompressingOutputStream>(std::move(strm));
		}
	}

	return strm;
}

int main(const int argc, char* argv[])
{
	const auto args = ParseArgs(argc, argv);
	if (!ProcessArgError(args))
	{
		return 1;
	}

	const auto input = MakeInputStream(args->inputFileName, args->options);
	const auto output = MakeOutputStream(args->outputFileName, args->options);

	output->Close();

	return 0;
}