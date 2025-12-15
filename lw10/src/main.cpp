#include "client/client.h"
#include "server/server.h"

#include <iostream>
#include <string>

int main(const int argc, char** argv)
{
	if (argc < 2)
	{
		std::cerr << "Usage: " << argv[0] << " [server|client]" << std::endl;
		return 1;
	}
	if (const std::string mode = argv[1]; mode == "server")
	{
		return RunServer();
	}
	else
	{
		if (mode == "client")
		{
			return RunClient(argc, argv);
		}
		std::cerr << "Unknown mode: " << mode << ". Use 'server' or 'client'." << std::endl;
		return 1;
	}
}