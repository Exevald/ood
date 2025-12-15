#include "Primitive.h"

#include <utility>

Primitive::Primitive(std::string id, ShapeType type, const Frame& frame, uint32_t color, std::string imgPath)
	: m_id(std::move(id))
	, m_type(type)
	, m_frame(frame)
	, m_imagePath(std::move(imgPath))
	, m_color(color)
{
}

void Primitive::SetId(const std::string& id)
{
	m_id = id;
}

std::string Primitive::GetId() const
{
	return m_id;
}

ShapeType Primitive::GetType() const
{
	return m_type;
}

std::string Primitive::GetImagePath() const
{
	return m_imagePath;
}

Frame Primitive::GetFrame() const
{
	return m_frame;
}

void Primitive::SetFrame(const Frame& frame)
{
	m_frame = frame;
}

void Primitive::MoveFrame(const double dx, const double dy)
{
	m_frame.x += dx;
	m_frame.y += dy;
}

std::shared_ptr<IShape> Primitive::Clone() const
{
	return std::make_shared<Primitive>(*this);
}

void Primitive::SetColor(const uint32_t color)
{
	m_color = color;
}

uint32_t Primitive::GetColor() const
{
	return m_color;
}