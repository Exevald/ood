#pragma once

#include "IShape.h"

#include <algorithm>

class ShapeGroup final : public IShape
{
public:
	explicit ShapeGroup(std::string id);

	void SetId(const std::string& id) override;
	[[nodiscard]] std::string GetId() const override;
	[[nodiscard]] ShapeType GetType() const override;

	void Add(const std::shared_ptr<IShape>& component) override;
	void Remove(const std::string& id) override;

	[[nodiscard]] const std::vector<std::shared_ptr<IShape>>& GetChildren() const override;
	[[nodiscard]] Frame GetFrame() const override;

	void SetFrame(const Frame& newFrame) override;
	void MoveFrame(double dx, double dy) override;
	[[nodiscard]] std::shared_ptr<IShape> Clone() const override;

	void SetColor(uint32_t color) override;
	[[nodiscard]] uint32_t GetColor() const override;

private:
	void RecalculateFrame();

	std::string m_id;
	std::vector<std::shared_ptr<IShape>> m_children;
	Frame m_cachedFrame;
};