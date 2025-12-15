#pragma once

#include "Geometry.h"

#include <memory>
#include <string>
#include <vector>

class IShape
{
public:
	virtual void SetId(const std::string& id) = 0;
	[[nodiscard]] virtual std::string GetId() const = 0;
	[[nodiscard]] virtual ShapeType GetType() const = 0;

	virtual void SetFrame(const Frame& rect) = 0;
	virtual void MoveFrame(double dx, double dy) = 0;

	[[nodiscard]] virtual std::shared_ptr<IShape> Clone() const = 0;

	virtual void Add(const std::shared_ptr<IShape>& component) {};
	virtual void Remove(const std::string& id) {};

	[[nodiscard]] virtual Frame GetFrame() const = 0;
	[[nodiscard]] virtual const std::vector<std::shared_ptr<IShape>>& GetChildren() const
	{
		static const std::vector<std::shared_ptr<IShape>> empty;
		return empty;
	}

	virtual void SetColor(uint32_t color) = 0;
	[[nodiscard]] virtual uint32_t GetColor() const = 0;

	virtual ~IShape() = default;
};