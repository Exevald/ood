#pragma once

#include "Document.h"

#include <string>

class IRepository {
public:
	virtual void Save(const Document& doc, const std::string& path) = 0;
	virtual void Load(Document& doc, const std::string& path) = 0;

	virtual ~IRepository() = default;
};
