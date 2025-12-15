#pragma once
#include <QRectF>
#include <qstring.h>

struct ShapeViewModel
{
	QString id;
	int type{};
	QRectF frame;
	QString imagePath;
	bool isSelected{};
	uint32_t color{};
};
