#include "adapters/ModernGraphicsRendererObjectAdapter.h"
#include "libs/CanvasDrawable.h"

namespace app
{
void PaintPicture(const shape_drawing_lib::CanvasPainter& painter)
{
	using namespace shape_drawing_lib;
	const Triangle triangle({ 10, 15 }, { 100, 200 }, { 150, 250 });
	const Rectangle rectangle({ 30, 40 }, 18, 24);
	painter.Draw(triangle);
	painter.Draw(rectangle);
}

void PaintPictureOnCanvas()
{
	graphics_lib::Canvas simpleCanvas;
	const shape_drawing_lib::CanvasPainter painter(simpleCanvas);
	PaintPicture(painter);
}
void PaintPictureOnModernGraphicsRenderer()
{
	modern_graphics_lib::ModernGraphicsRenderer renderer(std::cout);
	ModernGraphicsRendererObjectAdapter modernCanvas(renderer);
	const shape_drawing_lib::CanvasPainter painter(modernCanvas);
	renderer.BeginDraw();
	PaintPicture(painter);
	renderer.EndDraw();
}
} // namespace app

int main()
{
	std::cout << "Should we use new API (y)?" << std::endl
			  << "> ";
	if (std::string userInput; getline(std::cin, userInput) && (userInput == "y" || userInput == "Y"))
	{
		app::PaintPictureOnModernGraphicsRenderer();
	}
	else
	{
		app::PaintPictureOnCanvas();
	}
	return 0;
}