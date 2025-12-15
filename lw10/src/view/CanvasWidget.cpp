#include "CanvasWidget.h"

#include <QMouseEvent>
#include <QPainter>
#include <QPainterPath>

constexpr double HANDLE_SIZE = 8.0;

CanvasWidget::CanvasWidget(ViewModel* vm, QWidget* parent)
	: QWidget(parent)
	, m_vm(vm)
{
	setFocusPolicy(Qt::StrongFocus);
	setMouseTracking(true);
	connect(m_vm, &ViewModel::UpdateSignal, this, qOverload<>(&QWidget::update));
}

void CanvasWidget::paintEvent(QPaintEvent* event)
{
	Q_UNUSED(event);
	QPainter painter(this);
	painter.setRenderHint(QPainter::Antialiasing);

	const auto shapes = m_vm->GetShapes();
	for (const auto& shape : shapes)
	{
		if (shape.type != 4)
		{
			DrawShape(painter, shape);
		}
	}
	for (const auto& shape : shapes)
	{
		if (shape.isSelected)
		{
			DrawSelection(painter, shape);
		}
	}
}

void CanvasWidget::DrawShape(QPainter& painter, const ShapeViewModel& shape)
{
	const QColor color = QColor::fromRgba(shape.color);
	painter.setBrush(color);
	painter.setPen(QPen(Qt::black, 1));

	const QRectF frame = shape.frame;
	switch (static_cast<ShapeType>(shape.type))
	{
	case ShapeType::Rectangle: {
		painter.drawRect(frame);
		break;
	}
	case ShapeType::Triangle: {
		QPolygonF triangle;
		triangle << QPointF(frame.left(), frame.bottom());
		triangle << QPointF(frame.right(), frame.bottom());
		triangle << QPointF(frame.center().x(), frame.top());
		painter.drawPolygon(triangle);
		break;
	}
	case ShapeType::Ellipse: {
		painter.drawEllipse(frame);
		break;
	}
	case ShapeType::Image: {
		const QString path = shape.imagePath;
		if (path.isEmpty())
		{
			DrawImagePlaceholder(painter, frame, "No Path");
			return;
		}
		if (!m_imageCache.contains(path))
		{
			const QPixmap pm(path);
			if (pm.isNull())
			{
				DrawImagePlaceholder(painter, frame, "Load Error");
				m_imageCache.insert(path, QPixmap());
				return;
			}
			m_imageCache.insert(path, pm);
		}
		if (const QPixmap& pix = m_imageCache[path]; pix.isNull())
		{
			DrawImagePlaceholder(painter, frame, "Error");
		}
		else
		{
			painter.drawPixmap(frame.toRect(), pix);
			painter.setBrush(Qt::NoBrush);
			painter.drawRect(frame);
		}
		break;
	}
	default:
		break;
	}
}

void CanvasWidget::DrawSelection(QPainter& painter, const ShapeViewModel& shape)
{
	const QPen pen(Qt::magenta, 2, Qt::DashLine);
	painter.setPen(pen);
	painter.setBrush(Qt::NoBrush);

	const QRectF selFrame = shape.frame.adjusted(-2, -2, 2, 2);
	painter.drawRect(selFrame);
	painter.setPen(Qt::black);
	painter.setBrush(Qt::white);
	const QRectF frame = shape.frame;

	auto drawHandle = [&](const QPointF center) {
		painter.drawRect(QRectF(center.x() - HANDLE_SIZE / 2, center.y() - HANDLE_SIZE / 2, HANDLE_SIZE, HANDLE_SIZE));
	};

	drawHandle(frame.topLeft());
	drawHandle(frame.topRight());
	drawHandle(frame.bottomLeft());
	drawHandle(frame.bottomRight());
}

void CanvasWidget::DrawImagePlaceholder(QPainter& painter, const QRectF& frame, const QString& text)
{
	painter.setBrush(Qt::lightGray);
	painter.setPen(Qt::black);
	painter.drawRect(frame);
	painter.drawText(frame, Qt::AlignCenter, text);
}

CanvasWidget::ResizeHandle CanvasWidget::GetHandleAt(const QRectF& rect, const QPointF& pos)
{
	auto check = [&](const double x, double y) {
		return QRectF(x - HANDLE_SIZE / 2, y - HANDLE_SIZE / 2, HANDLE_SIZE, HANDLE_SIZE).contains(pos);
	};

	if (check(rect.left(), rect.top()))
		return ResizeHandle::TopLeft;
	if (check(rect.right(), rect.top()))
		return ResizeHandle::TopRight;
	if (check(rect.left(), rect.bottom()))
		return ResizeHandle::BottomLeft;
	if (check(rect.right(), rect.bottom()))
		return ResizeHandle::BottomRight;

	return ResizeHandle::None;
}

void CanvasWidget::UpdateCursor(const QPointF& pos)
{
	if (m_isDragging)
	{
		return;
	}
	const auto shapes = m_vm->GetShapes();
	for (const auto& s : shapes)
	{
		if (s.isSelected)
		{
			if (const ResizeHandle handle = GetHandleAt(s.frame, pos); handle != ResizeHandle::None)
			{
				switch (handle)
				{
				case ResizeHandle::TopLeft:
				case ResizeHandle::BottomRight: {
					setCursor(Qt::SizeFDiagCursor);
					break;
				}
				case ResizeHandle::TopRight:
				case ResizeHandle::BottomLeft: {
					setCursor(Qt::SizeBDiagCursor);
					break;
				}
				default: {
				}
				}
				return;
			}
		}
	}
	bool overShape = false;
	for (auto it = shapes.rbegin(); it != shapes.rend(); ++it)
	{
		if (it->frame.contains(pos))
		{
			setCursor(Qt::SizeAllCursor);
			overShape = true;
			break;
		}
	}
	if (!overShape)
	{
		setCursor(Qt::ArrowCursor);
	}
}

void CanvasWidget::mousePressEvent(QMouseEvent* event)
{
	if (event->button() == Qt::LeftButton)
	{
		bool handleClicked = false;
		const auto shapes = m_vm->GetShapes();
		for (const auto& s : shapes)
		{
			if (s.isSelected)
			{
				m_currentHandle = GetHandleAt(s.frame, event->localPos());
				if (m_currentHandle != ResizeHandle::None)
				{
					m_isDragging = true;
					m_dragStartPos = event->localPos();
					m_lastMousePos = event->localPos();
					m_initialResizeRect = s.frame;
					handleClicked = true;
					break;
				}
			}
		}
		if (!handleClicked)
		{
			const bool ctrl = event->modifiers() & Qt::ControlModifier;
			m_vm->SelectAt(event->localPos(), ctrl);

			m_isDragging = true;
			m_dragStartPos = event->localPos();
			m_lastMousePos = event->localPos();
			m_currentHandle = ResizeHandle::None;
		}
		m_vm->StartTransform();
	}
}

void CanvasWidget::mouseMoveEvent(QMouseEvent* event)
{
	UpdateCursor(event->localPos());
	if (!m_isDragging)
	{
		return;
	}
	const QPointF delta = event->localPos() - m_dragStartPos;

	if (m_currentHandle != ResizeHandle::None)
	{
		QRectF newRect = m_initialResizeRect;
		switch (m_currentHandle)
		{
		case ResizeHandle::BottomRight: {
			newRect.setWidth(m_initialResizeRect.width() + delta.x());
			newRect.setHeight(m_initialResizeRect.height() + delta.y());
			break;
		}
		case ResizeHandle::TopLeft: {
			newRect.setTop(m_initialResizeRect.top() + delta.y());
			newRect.setLeft(m_initialResizeRect.left() + delta.x());
			break;
		}
		case ResizeHandle::BottomLeft: {
			newRect.setLeft(m_initialResizeRect.left() + delta.x());
			newRect.setBottom(m_initialResizeRect.bottom() + delta.y());
			break;
		}
		case ResizeHandle::TopRight: {
			newRect.setTop(m_initialResizeRect.top() + delta.y());
			newRect.setRight(m_initialResizeRect.right() + delta.x());
			break;
		}
		default:
			break;
		}
		m_vm->ResizeSelected(newRect.normalized());
	}
	else
	{
		const QPointF moveDelta = event->localPos() - m_lastMousePos;
		m_vm->MoveSelected(moveDelta.x(), moveDelta.y());
	}
	m_lastMousePos = event->localPos();
}

void CanvasWidget::mouseReleaseEvent(QMouseEvent* event)
{
	if (event->button() == Qt::LeftButton)
	{
		if (m_isDragging)
		{
			m_isDragging = false;
			m_vm->EndTransform();
		}
	}
}