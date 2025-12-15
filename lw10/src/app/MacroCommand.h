#pragma once

#include "ICommand.h"

#include <memory>
#include <vector>

class MacroCommand final : public ICommand
{
public:
	void AddCommand(std::shared_ptr<ICommand> cmd)
	{
		m_commands.push_back(std::move(cmd));
	}

	void Execute() override
	{
		for (const auto& cmd : m_commands)
		{
			cmd->Execute();
		}
	}
	void Undo() override
	{
		for (auto it = m_commands.rbegin(); it != m_commands.rend(); ++it)
		{
			(*it)->Undo();
		}
	}
	[[nodiscard]] bool IsEmpty() const
	{
		return m_commands.empty();
	}

private:
	std::vector<std::shared_ptr<ICommand>> m_commands;
};