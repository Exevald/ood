#include "Drawer.h"
#include "Image.h"
#include <iostream>

void Print(const Image& img)
{
	Size s = img.GetSize();
	for (int y = 0; y < s.height; ++y)
	{
		for (int x = 0; x < s.width; ++x)
		{
			std::cout << img.GetPixel(x, y);
		}
		std::cout << '\n';
	}
}

int main()
{
	Image img(30, 11, '.');
	DrawCircle(img, { 15, 5 }, 4, '#');

	Print(img);
	return 0;
}