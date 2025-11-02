#include "Canvas.h"

#include <iomanip>

namespace
{
std::string GetHexStrFromUint32(uint32_t const& uint32)
{
	std::stringstream stream;
	stream << std::setfill('0') << std::setw(6) << std::hex << uint32;
	return stream.str();
}
} // namespace

void graphics_lib::Canvas::MoveTo(const int x, const int y)
{
	std::cout << "MoveTo (" << x << ", " << y << ")" << std::endl;
}

void graphics_lib::Canvas::LineTo(const int x, const int y)
{
	std::cout << "LineTo (" << x << ", " << y << ")" << std::endl;
}

void graphics_lib::Canvas::SetColor(const uint32_t rgbColor)
{
	std::cout << "SetColor (" << GetHexStrFromUint32(rgbColor) << ")" << std::endl;
}