#include "ShapeGroup.h"
#include "Geometry.h"

#include <utility>

ShapeGroup::ShapeGroup(std::string id)
	: m_id(std::move(id))
	, m_cachedFrame{ 0, 0, 0, 0 }
{
}

void ShapeGroup::SetId(const std::string& id)
{
	m_id = id;
}

std::string ShapeGroup::GetId() const
{
	return m_id;
}

ShapeType ShapeGroup::GetType() const
{
	return ShapeType::Group;
}

void ShapeGroup::Add(const std::shared_ptr<IShape>& component)
{
	m_children.push_back(component);
	RecalculateFrame();
}

void ShapeGroup::Remove(const std::string& id)
{
	m_children.erase(std::remove_if(m_children.begin(), m_children.end(),
						 [&](auto& c) { return c->GetId() == id; }),
		m_children.end());
	RecalculateFrame();
}

const std::vector<std::shared_ptr<IShape>>& ShapeGroup::GetChildren() const
{
	return m_children;
}

Frame ShapeGroup::GetFrame() const
{
	return m_cachedFrame;
}

void ShapeGroup::SetFrame(const Frame& newFrame)
{
	const double oldX = m_cachedFrame.x;
	const double oldY = m_cachedFrame.y;
	const double oldW = m_cachedFrame.width;
	const double oldH = m_cachedFrame.height;

	if (std::abs(oldW) < 1e-5 || std::abs(oldH) < 1e-5)
	{
		MoveFrame(newFrame.x - oldX, newFrame.y - oldY);
		return;
	}

	const double scaleX = newFrame.width / oldW;
	const double scaleY = newFrame.height / oldH;

	for (const auto& child : m_children)
	{
		auto [x, y, width, height] = child->GetFrame();

		const double relX = x - oldX;
		const double relY = y - oldY;

		const double nx = newFrame.x + relX * scaleX;
		const double ny = newFrame.y + relY * scaleY;
		const double nw = width * scaleX;
		const double nh = height * scaleY;

		child->SetFrame({ nx, ny, nw, nh });
	}

	RecalculateFrame();
}

void ShapeGroup::MoveFrame(const double dx, const double dy)
{
	for (const auto& child : m_children)
	{
		child->MoveFrame(dx, dy);
	}
	RecalculateFrame();
}

std::shared_ptr<IShape> ShapeGroup::Clone() const
{
	auto group = std::make_shared<ShapeGroup>(m_id);
	for (auto& child : m_children)
	{
		group->Add(child->Clone());
	}
	return group;
}

void ShapeGroup::SetColor(uint32_t color)
{
	for (const auto& child : m_children)
	{
		child->SetColor(color);
	}
}

uint32_t ShapeGroup::GetColor() const
{
	if (!m_children.empty())
	{
		return m_children.front()->GetColor();
	}
	return 0xFF000000;
}

void ShapeGroup::RecalculateFrame()
{
	if (m_children.empty())
	{
		m_cachedFrame = { 0, 0, 0, 0 };
		return;
	}
	m_cachedFrame = m_children[0]->GetFrame();
	for (size_t i = 1; i < m_children.size(); ++i)
	{
		m_cachedFrame = Frame::Unite(m_cachedFrame, m_children[i]->GetFrame());
	}
}