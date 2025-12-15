#include "JsonRepository.h"
#include "../domain/Primitive.h"

#include <QFile>
#include <QJsonArray>
#include <QJsonDocument>
#include <QJsonObject>

void JsonRepository::Save(const Document& doc, const std::string& path)
{
	// поправить кейс с удалённой картинкой
	QJsonArray arr;
	for (const auto& shape : doc.GetShapes())
	{
		QJsonObject obj;
		obj["id"] = QString::fromStdString(shape->GetId());
		obj["type"] = static_cast<int>(shape->GetType());
		auto [x, y, width, height] = shape->GetFrame();
		obj["x"] = x;
		obj["y"] = y;
		obj["w"] = width;
		obj["h"] = height;

		if (const auto primitive = std::dynamic_pointer_cast<Primitive>(shape))
		{
			if (!primitive->GetImagePath().empty())
			{
				obj["image"] = QString::fromStdString(primitive->GetImagePath());
			}
		}
		arr.append(obj);
	}

	QJsonObject root;
	root["shapes"] = arr;
	if (QFile file(QString::fromStdString(path)); file.open(QIODevice::WriteOnly))
	{
		file.write(QJsonDocument(root).toJson());
	}
}

void JsonRepository::Load(Document& document, const std::string& path)
{
	document.Clear();
	QFile f(QString::fromStdString(path));
	if (!f.open(QIODevice::ReadOnly))
	{
		return;
	}

	auto root = QJsonDocument::fromJson(f.readAll()).object();
	auto arr = root["shapes"].toArray();

	// пофиксить загрузку невалидного json-а
	// починить кейс penis.json
	for (const auto& val : arr)
	{
		auto obj = val.toObject();
		auto type = static_cast<ShapeType>(obj["type"].toInt());
		Frame frame{ obj["x"].toDouble(), obj["y"].toDouble(), obj["w"].toDouble(), obj["h"].toDouble() };
		std::string id = obj["id"].toString().toStdString();
		std::string img = obj["image"].toString().toStdString();

		auto shape = std::make_shared<Primitive>(id, type, frame, 0xFF0000FF, img);
		document.AddShape(shape);
	}
}