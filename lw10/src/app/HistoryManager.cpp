#include "HistoryManager.h"

void HistoryManager::Execute(const std::shared_ptr<ICommand>& command)
{
	command->Execute();
	m_undoStack.push(command);
	while (!m_redoStack.empty())
	{
		m_redoStack.pop();
	}
}

void HistoryManager::Undo()
{
	if (m_undoStack.empty())
	{
		return;
	}
	const auto command = m_undoStack.top();
	m_undoStack.pop();
	command->Undo();
	m_redoStack.push(command);
}

void HistoryManager::Redo()
{
	if (m_redoStack.empty())
	{
		return;
	}

	const auto command = m_redoStack.top();
	m_redoStack.pop();
	command->Execute();
	m_undoStack.push(command);
}