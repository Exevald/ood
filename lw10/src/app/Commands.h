#pragma once

#include "../domain/Document.h"
#include "../domain/IShape.h"
#include "../domain/ShapeGroup.h"
#include "ICommand.h"

#include <memory>
#include <utility>

class AddShapeCommand final : public ICommand
{
public:
	AddShapeCommand(const std::shared_ptr<Document>& document, const std::shared_ptr<IShape>& shape)
		: m_document(document)
		, m_shape(shape)
	{
	}

	void Execute() override
	{
		m_document->AddShape(m_shape);
	}
	void Undo() override
	{
		m_document->RemoveShape(m_shape->GetId());
	}

private:
	std::shared_ptr<Document> m_document;
	std::shared_ptr<IShape> m_shape;
};

class TransformShapeCommand final : public ICommand
{
public:
	TransformShapeCommand(const std::shared_ptr<IShape>& shape, const Frame& oldRect, const Frame& newRect)
		: m_targetShape(shape)
		, m_oldRect(oldRect)
		, m_newRect(newRect)
	{
	}

	void Execute() override
	{
		m_targetShape->SetFrame(m_newRect);
	}
	void Undo() override
	{
		m_targetShape->SetFrame(m_oldRect);
	}

private:
	std::shared_ptr<IShape> m_targetShape;
	Frame m_oldRect;
	Frame m_newRect;
};

class RemoveShapeCommand final : public ICommand
{
public:
	RemoveShapeCommand(std::shared_ptr<Document> document, const std::shared_ptr<IShape>& shape)
		: m_doc(std::move(document))
		, m_shape(shape)
		, m_id(shape->GetId())
	{
	}

	void Execute() override
	{
		m_doc->RemoveShape(m_id);
	}
	void Undo() override
	{
		m_doc->AddShape(m_shape);
	}

private:
	std::shared_ptr<Document> m_doc;
	std::shared_ptr<IShape> m_shape;
	std::string m_id;
};

class UngroupShapeCommand final : public ICommand
{
public:
	UngroupShapeCommand(std::shared_ptr<Document> document, std::shared_ptr<ShapeGroup> group)
		: m_document(std::move(document))
		, m_group(std::move(group))
		, m_groupId(m_group->GetId())
	{
		m_children = m_group->GetChildren();
	}

	void Execute() override
	{
		m_document->RemoveShape(m_groupId);
		for (const auto& child : m_children)
		{
			m_document->AddShape(child);
		}
	}
	void Undo() override
	{
		for (const auto& child : m_children)
		{
			m_document->RemoveShape(child->GetId());
		}
		m_document->AddShape(m_group);
	}

private:
	std::shared_ptr<Document> m_document;
	std::shared_ptr<ShapeGroup> m_group;
	std::string m_groupId;
	std::vector<std::shared_ptr<IShape>> m_children;
};