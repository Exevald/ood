#pragma once

#include "../app/Commands.h"
#include "../app/HistoryManager.h"
#include "../domain/Document.h"
#include "JsonRepository.h"
#include "ShapeViewModel.h"

#include <QObject>
#include <QRectF>
#include <map>
#include <memory>
#include <qstring.h>

class ViewModel final : public QObject
{
	Q_OBJECT

public:
	explicit ViewModel(QObject* parent = nullptr);

	[[nodiscard]] std::vector<ShapeViewModel> GetShapes() const;

	void AddRect();
	void AddTriangle();
	void AddEllipse();
	void AddImage(const QString& path);

	void GroupSelected();
	void UngroupSelected();
	void SelectAt(QPointF pos, bool ctrl);

	void StartTransform();
	void MoveSelected(double dx, double dy);
	void ResizeSelected(const QRectF& newRect);
	void EndTransform();

	void SetColorForSelected(uint32_t color);

	void DeleteSelected();
	void Copy();
	void Paste();

	void Save(const QString& path) const;
	void Load(const QString& path);

	void Undo();
	void Redo();

private:
	void AddShape(ShapeType type);
	void AppendShapesRecursive(const std::shared_ptr<IShape>& shape, std::vector<ShapeViewModel>& result) const;

	std::shared_ptr<Document> m_document;
	HistoryManager m_historyManager;
	std::shared_ptr<IRepository> m_documentRepository;
	std::vector<std::string> m_selectedShapesIds;
	std::vector<std::shared_ptr<IShape>> m_clipboard;
	std::map<std::string, Frame> m_initialFrames;

signals:
	void UpdateSignal();
};