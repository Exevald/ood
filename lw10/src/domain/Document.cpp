#include "Document.h"

#include <algorithm>

void Document::AddShape(const std::shared_ptr<IShape>& shape)
{
	m_shapes.push_back(shape);
}

void Document::RemoveShape(const std::string& id)
{
	m_shapes.erase(std::remove_if(m_shapes.begin(), m_shapes.end(), [&](auto& s) { return s->GetId() == id; }), m_shapes.end());
}

void Document::Clear()
{
	m_shapes.clear();
}

std::shared_ptr<IShape> Document::GetShape(const std::string& id) const
{
	const auto it = std::find_if(m_shapes.begin(), m_shapes.end(),
		[&](auto& s) { return s->GetId() == id; });
	return (it != m_shapes.end()) ? *it : nullptr;
}

const std::vector<std::shared_ptr<IShape>>& Document::GetShapes() const
{
	return m_shapes;
}