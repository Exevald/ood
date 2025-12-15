#pragma once

#include "IShape.h"

#include <memory>
#include <vector>

class Document
{
public:
	void AddShape(const std::shared_ptr<IShape>& shape);
	void RemoveShape(const std::string& id);
	void Clear();

	[[nodiscard]] std::shared_ptr<IShape> GetShape(const std::string& id) const;
	[[nodiscard]] const std::vector<std::shared_ptr<IShape>>& GetShapes() const;

private:
	std::vector<std::shared_ptr<IShape>> m_shapes;
};
