#include "ViewModel.h"
#include "../app/MacroCommand.h"
#include "../domain/Primitive.h"
#include "../domain/ShapeGroup.h"

#include <QUuid>

ViewModel::ViewModel(QObject* parent)
	: QObject(parent)
{
	m_document = std::make_shared<Document>();
	m_documentRepository = std::make_shared<JsonRepository>();
}

std::vector<ShapeViewModel> ViewModel::GetShapes() const
{
	std::vector<ShapeViewModel> result;
	for (const auto& s : m_document->GetShapes())
	{
		AppendShapesRecursive(s, result);
	}
	return result;
}

void ViewModel::AddRect()
{
	AddShape(ShapeType::Rectangle);
}

void ViewModel::AddTriangle()
{
	AddShape(ShapeType::Triangle);
}

void ViewModel::AddEllipse()
{
	AddShape(ShapeType::Ellipse);
}

void ViewModel::AddImage(const QString& path)
{
	std::string id = QUuid::createUuid().toString().toStdString();
	Frame frame{ 100, 100, 200, 150 };
	auto shape = std::make_shared<Primitive>(id, ShapeType::Image, frame, 0xFF0000FF, path.toStdString());
	const auto command = std::make_shared<AddShapeCommand>(m_document, shape);

	m_historyManager.Execute(command);
	emit UpdateSignal();
}

void ViewModel::GroupSelected()
{
	if (m_selectedShapesIds.size() < 2)
	{
		return;
	}
	std::string groupId = QUuid::createUuid().toString().toStdString();
	const auto group = std::make_shared<ShapeGroup>(groupId);

	for (const auto& id : m_selectedShapesIds)
	{
		if (const auto shape = m_document->GetShape(id))
		{
			group->Add(shape->Clone());
			m_document->RemoveShape(id);
		}
	}

	m_document->AddShape(group);
	m_selectedShapesIds.clear();
	m_selectedShapesIds.push_back(groupId);
	emit UpdateSignal();
}

void ViewModel::UngroupSelected()
{
	if (m_selectedShapesIds.empty())
	{
		return;
	}
	const auto macroCommand = std::make_shared<MacroCommand>();
	std::vector<std::string> newSelection;
	bool hasGroups = false;

	for (const auto& id : m_selectedShapesIds)
	{
		auto shape = m_document->GetShape(id);
		if (auto group = std::dynamic_pointer_cast<ShapeGroup>(shape))
		{
			auto cmd = std::make_shared<UngroupShapeCommand>(m_document, group);
			macroCommand->AddCommand(cmd);
			for (const auto& child : group->GetChildren())
			{
				newSelection.push_back(child->GetId());
			}
			hasGroups = true;
		}
	}
	if (hasGroups && !macroCommand->IsEmpty())
	{
		m_historyManager.Execute(macroCommand);
		m_selectedShapesIds = newSelection;
		emit UpdateSignal();
	}
}

void ViewModel::SelectAt(const QPointF pos, const bool ctrl)
{
	const Point point{ pos.x(), pos.y() };
	std::string hitId;

	auto shapes = m_document->GetShapes();
	for (auto it = shapes.rbegin(); it != shapes.rend(); ++it)
	{
		if ((*it)->GetFrame().Contains(point))
		{
			hitId = (*it)->GetId();
			break;
		}
	}
	if (!ctrl)
	{
		m_selectedShapesIds.clear();
	}
	if (!hitId.empty())
	{
		if (const auto it = std::find(m_selectedShapesIds.begin(), m_selectedShapesIds.end(), hitId);
			it == m_selectedShapesIds.end())
		{
			m_selectedShapesIds.push_back(hitId);
		}
		else if (ctrl)
		{
			m_selectedShapesIds.erase(it);
		}
	}
	emit UpdateSignal();
}

void ViewModel::StartTransform()
{
	m_initialFrames.clear();
	for (const auto& id : m_selectedShapesIds)
	{
		if (const auto shape = m_document->GetShape(id))
		{
			m_initialFrames[id] = shape->GetFrame();
		}
	}
}

void ViewModel::MoveSelected(const double dx, const double dy)
{
	for (const auto& id : m_selectedShapesIds)
	{
		if (const auto shape = m_document->GetShape(id))
		{
			shape->MoveFrame(dx, dy);
		}
	}
	emit UpdateSignal();
}

void ViewModel::ResizeSelected(const QRectF& newRect)
{
	for (const auto& id : m_selectedShapesIds)
	{
		if (const auto shape = m_document->GetShape(id))
		{
			Frame frame{ newRect.x(), newRect.y(), newRect.width(), newRect.height() };
			shape->SetFrame(frame);
		}
	}
	emit UpdateSignal();
}

void ViewModel::EndTransform()
{
	const auto macroCommand = std::make_shared<MacroCommand>();
	bool changed = false;
	for (const auto& [id, oldFrame] : m_initialFrames)
	{
		if (const auto shape = m_document->GetShape(id))
		{
			if (const auto currentFrame = shape->GetFrame();
				currentFrame.x != oldFrame.x
				|| currentFrame.y != oldFrame.y
				|| currentFrame.width != oldFrame.width
				|| currentFrame.height != oldFrame.height)
			{
				auto cmd = std::make_shared<TransformShapeCommand>(shape, oldFrame, currentFrame);
				macroCommand->AddCommand(cmd);
				changed = true;
			}
		}
	}
	if (changed && !macroCommand->IsEmpty())
	{
		m_historyManager.Execute(macroCommand);
	}

	m_initialFrames.clear();
}

void ViewModel::SetColorForSelected(const uint32_t color)
{
	for (const auto& id : m_selectedShapesIds)
	{
		if (const auto shape = m_document->GetShape(id))
		{
			shape->SetColor(color);
		}
	}
	emit UpdateSignal();
}

void ViewModel::DeleteSelected()
{
	if (m_selectedShapesIds.empty())
	{
		return;
	}
	const auto macroCommand = std::make_shared<MacroCommand>();
	for (const auto& id : m_selectedShapesIds)
	{
		if (const auto shape = m_document->GetShape(id))
		{
			auto cmd = std::make_shared<RemoveShapeCommand>(m_document, shape);
			macroCommand->AddCommand(cmd);
		}
	}
	if (!macroCommand->IsEmpty())
	{
		m_historyManager.Execute(macroCommand);
	}

	m_selectedShapesIds.clear();
	emit UpdateSignal();
}

void ViewModel::Copy()
{
	m_clipboard.clear();
	for (const auto& id : m_selectedShapesIds)
	{
		if (const auto shape = m_document->GetShape(id))
		{
			m_clipboard.push_back(shape->Clone());
		}
	}
}

void ViewModel::Paste()
{
	if (m_clipboard.empty())
	{
		return;
	}
	m_selectedShapesIds.clear();
	const auto macroCommand = std::make_shared<MacroCommand>();

	std::function<std::shared_ptr<IShape>(const std::shared_ptr<IShape>&)> regenerateHierarchy;
	regenerateHierarchy = [&](const std::shared_ptr<IShape>& source) -> std::shared_ptr<IShape> {
		std::string id = QUuid::createUuid().toString().toStdString();
		if (const auto primitive = std::dynamic_pointer_cast<Primitive>(source))
		{
			return std::make_shared<Primitive>(id, primitive->GetType(), primitive->GetFrame(), primitive->GetColor(), primitive->GetImagePath());
		}
		if (const auto group = std::dynamic_pointer_cast<ShapeGroup>(source))
		{
			auto newG = std::make_shared<ShapeGroup>(id);
			for (const auto& child : group->GetChildren())
			{
				newG->Add(regenerateHierarchy(child));
			}
			return newG;
		}
		return nullptr;
	};

	for (const auto& clipboardShape : m_clipboard)
	{
		auto tempShape = clipboardShape->Clone();
		std::shared_ptr<IShape> finalShape = nullptr;
		std::string rootId = QUuid::createUuid().toString().toStdString();

		if (const auto primitive = std::dynamic_pointer_cast<Primitive>(tempShape))
		{
			finalShape = std::make_shared<Primitive>(rootId, primitive->GetType(), primitive->GetFrame(), primitive->GetColor(), primitive->GetImagePath());
		}
		else if (const auto group = std::dynamic_pointer_cast<ShapeGroup>(tempShape))
		{
			const auto newGroup = std::make_shared<ShapeGroup>(rootId);
			for (const auto& child : group->GetChildren())
			{
				newGroup->Add(regenerateHierarchy(child));
			}
			finalShape = newGroup;
		}

		if (finalShape)
		{
			finalShape->MoveFrame(20, 20);
			auto cmd = std::make_shared<AddShapeCommand>(m_document, finalShape);
			macroCommand->AddCommand(cmd);
			m_selectedShapesIds.push_back(finalShape->GetId());
		}
	}
	if (!macroCommand->IsEmpty())
	{
		m_historyManager.Execute(macroCommand);
	}

	emit UpdateSignal();
}

void ViewModel::Save(const QString& path) const
{
	m_documentRepository->Save(*m_document, path.toStdString());
}

void ViewModel::Load(const QString& path)
{
	m_documentRepository->Load(*m_document, path.toStdString());
	m_selectedShapesIds.clear();
	emit UpdateSignal();
}

void ViewModel::Undo()
{
	m_historyManager.Undo();
	emit UpdateSignal();
}

void ViewModel::Redo()
{
	m_historyManager.Redo();
	emit UpdateSignal();
}

void ViewModel::AddShape(ShapeType type)
{
	std::string id = QUuid::createUuid().toString().toStdString();
	Frame frame{ 100, 100, 100, 100 };
	auto shape = std::make_shared<Primitive>(id, type, frame, 0xFF0000FF);

	const auto cmd = std::make_shared<AddShapeCommand>(m_document, shape);
	m_historyManager.Execute(cmd);
	emit UpdateSignal();
}

void ViewModel::AppendShapesRecursive(const std::shared_ptr<IShape>& shape, std::vector<ShapeViewModel>& result) const
{
	ShapeViewModel vm;
	auto [x, y, width, height] = shape->GetFrame();
	vm.id = QString::fromStdString(shape->GetId());
	vm.type = static_cast<int>(shape->GetType());
	vm.frame = QRectF(x, y, width, height);
	vm.color = shape->GetColor();
	vm.isSelected = std::find(m_selectedShapesIds.begin(), m_selectedShapesIds.end(), shape->GetId()) != m_selectedShapesIds.end();

	if (const auto primitive = std::dynamic_pointer_cast<Primitive>(shape))
	{
		vm.imagePath = QString::fromStdString(primitive->GetImagePath());
	}
	if (shape->GetType() != ShapeType::Group)
	{
		result.push_back(vm);
	}
	else
	{
		result.push_back(vm);
		for (const auto& child : shape->GetChildren())
		{
			AppendShapesRecursive(child, result);
		}
	}
}