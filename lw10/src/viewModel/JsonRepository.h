#pragma once

#include "../domain/IRepository.h"

class JsonRepository final : public IRepository
{
public:
	void Save(const Document& doc, const std::string& path) override;
	void Load(Document& document, const std::string& path) override;
};