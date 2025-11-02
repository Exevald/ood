#pragma once

#include "ModernGraphicsRendererClassAdapter.h"

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

ModernGraphicsRendererClassAdapter::ModernGraphicsRendererClassAdapter(std::ostream& out)
	: ModernGraphicsRenderer(out)
	, m_color(CreateRGBAColor(0))
	, m_point({ 0, 0 })
{
}

void ModernGraphicsRendererClassAdapter::BeginDraw()
{
	ModernGraphicsRenderer::BeginDraw();
}

void ModernGraphicsRendererClassAdapter::EndDraw()
{
	ModernGraphicsRenderer::EndDraw();
}

void ModernGraphicsRendererClassAdapter::SetColor(uint32_t const color)
{
	m_color = CreateRGBAColor(color);
}

void ModernGraphicsRendererClassAdapter::MoveTo(int const x, int const y)
{
	m_point = { x, y };
}

void ModernGraphicsRendererClassAdapter::LineTo(int const x, int const y)
{
	modern_graphics_lib::Point const to = { x, y };
	DrawLine(m_point, to, m_color);
	m_point = to;
}