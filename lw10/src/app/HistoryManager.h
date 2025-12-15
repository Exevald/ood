#pragma once

#include "ICommand.h"

#include <memory>
#include <stack>

class HistoryManager
{
public:
	void Execute(const std::shared_ptr<ICommand>& command);
	void Undo();
	void Redo();

private:
	std::stack<std::shared_ptr<ICommand>> m_undoStack;
	std::stack<std::shared_ptr<ICommand>> m_redoStack;
};
