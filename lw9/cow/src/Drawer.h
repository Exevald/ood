#pragma once

#include "Image.h"

#include <cmath>

void DrawLine(Image& image, Point p1, Point p2, char color);
void DrawCircle(Image& image, Point center, int radius, char color);
void FillCircle(Image& image, Point center, int radius, char color);