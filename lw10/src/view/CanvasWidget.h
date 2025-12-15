#pragma once

#include "../viewModel/ViewModel.h"
#include <QMap>
#include <QWidget>

class CanvasWidget final : public QWidget
{
	Q_OBJECT

public:
	explicit CanvasWidget(ViewModel* vm, QWidget* parent = nullptr);

protected:
	void paintEvent(QPaintEvent* event) override;
	void mousePressEvent(QMouseEvent* event) override;
	void mouseMoveEvent(QMouseEvent* event) override;
	void mouseReleaseEvent(QMouseEvent* event) override;

private:
	void DrawShape(QPainter& painter, const ShapeViewModel& shape);
	static void DrawSelection(QPainter& painter, const ShapeViewModel& shape);
	static void DrawImagePlaceholder(QPainter& painter, const QRectF& frame, const QString& text);

	enum class ResizeHandle
	{
		None,
		TopLeft,
		TopRight,
		BottomLeft,
		BottomRight,
		Top,
		Bottom,
		Left,
		Right
	};

	static ResizeHandle GetHandleAt(const QRectF& rect, const QPointF& pos);
	void UpdateCursor(const QPointF& pos);

	ResizeHandle m_currentHandle = ResizeHandle::None;
	QRectF m_initialResizeRect;
	QPointF m_dragStartPos;
	ViewModel* m_vm;
	bool m_isDragging = false;
	QPointF m_lastMousePos;
	QMap<QString, QPixmap> m_imageCache;
};