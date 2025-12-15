#pragma once
#include "IShape.h"

class Primitive final : public IShape
{
public:
	Primitive(std::string id, ShapeType type, const Frame& frame,
		  uint32_t color = 0xFF0000FF, std::string imgPath = "");

	void SetId(const std::string& id) override;
	[[nodiscard]] std::string GetId() const override;
	[[nodiscard]] ShapeType GetType() const override;
	[[nodiscard]] std::string GetImagePath() const;
	[[nodiscard]] Frame GetFrame() const override;

	void SetFrame(const Frame& frame) override;
	void MoveFrame(double dx, double dy) override;
	[[nodiscard]] std::shared_ptr<IShape> Clone() const override;
	void SetColor(uint32_t color) override;
	[[nodiscard]] uint32_t GetColor() const override;

private:
	std::string m_id;
	ShapeType m_type;
	Frame m_frame;
	std::string m_imagePath;
	uint32_t m_color;
};
