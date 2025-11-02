#pragma once

#include "../libs/Canvas.h"
#include "../libs/GraphicsRenderer.h"

class ModernGraphicsRendererClassAdapter final
	: public graphics_lib::ICanvas
	, modern_graphics_lib::ModernGraphicsRenderer
{
public:
	explicit ModernGraphicsRendererClassAdapter(std::ostream& out);

	void BeginDraw() override;
	void EndDraw() override;

	void SetColor(uint32_t color) override;
	void MoveTo(int x, int y) override;
	void LineTo(int x, int y) override;

private:
	modern_graphics_lib::RGBAColor m_color;
	modern_graphics_lib::Point m_point;
};