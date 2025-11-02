#include "ModernGraphicsRendererObjectAdapter.h"

namespace
{
float GetColor(uint32_t const color, uint32_t const offset)
{
	constexpr uint32_t byteMask{ 0xFF };
	const uint32_t mask{ byteMask << offset };
	return static_cast<float>((color & mask) >> offset & byteMask) / 0xFF;
}

modern_graphics_lib::RGBAColor CreateRGBAColor(const uint32_t color)
{
	constexpr uint32_t blueOffset{ 0x00 };
	constexpr uint32_t greenOffset{ 0x08 };
	constexpr uint32_t redOffset{ 0x10 };

	const auto red = GetColor(color, redOffset);
	const auto green = GetColor(color, greenOffset);
	const auto blue = GetColor(color, blueOffset);

	return modern_graphics_lib::RGBAColor{ red, green, blue, 1.0f };
}
} // namespace

ModernGraphicsRendererObjectAdapter::ModernGraphicsRendererObjectAdapter(modern_graphics_lib::ModernGraphicsRenderer& renderer)
	: m_renderer(renderer)
	, m_color(CreateRGBAColor(0))
	, m_point({ 0, 0 })
{
}

void ModernGraphicsRendererObjectAdapter::SetColor(uint32_t const color)
{
	m_color = CreateRGBAColor(color);
}

void ModernGraphicsRendererObjectAdapter::MoveTo(int const x, int const y)
{
	m_point = { x, y };
}

void ModernGraphicsRendererObjectAdapter::LineTo(int const x, int const y)
{
	modern_graphics_lib::Point const to = { x, y };
	m_renderer.DrawLine(m_point, to, m_color);
	m_point = to;
}