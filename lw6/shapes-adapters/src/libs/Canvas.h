#pragma once

#include <iostream>

namespace graphics_lib
{

class ICanvas
{
public:
	virtual void SetColor(uint32_t rgbColor) = 0;
	virtual void MoveTo(int x, int y) = 0;
	virtual void LineTo(int x, int y) = 0;
	virtual ~ICanvas() = default;
};

class Canvas final : public ICanvas
{
public:
	void MoveTo(int x, int y) override;
	void LineTo(int x, int y) override;
	void SetColor(uint32_t rgbColor) override;
};
} // namespace graphics_lib