#pragma once

#include "../libs/Canvas.h"
#include "../libs/GraphicsRenderer.h"

class ModernGraphicsRendererObjectAdapter final : public graphics_lib::ICanvas
{
public:
	explicit ModernGraphicsRendererObjectAdapter(modern_graphics_lib::ModernGraphicsRenderer& renderer);

	void SetColor(uint32_t color) override;
	void MoveTo(int x, int y) override;\
	void LineTo(int x, int y) override;

private:
	modern_graphics_lib::ModernGraphicsRenderer& m_renderer;
	modern_graphics_lib::RGBAColor m_color;
	modern_graphics_lib::Point m_point;
};