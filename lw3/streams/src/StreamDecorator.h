#pragma once

#include "IDataStream.h"

class InputStreamDecorator : public IInputDataStream
{
public:
	[[nodiscard]] bool IsEOF() const override
	{
		return m_stream->IsEOF();
	}

	uint8_t ReadByte() override
	{
		return m_stream->ReadByte();
	}

	std::streamsize ReadBlock(void* dstBuffer, std::streamsize const size) override
	{
		return m_stream->ReadBlock(dstBuffer, size);
	}

protected:
	explicit InputStreamDecorator(IInputPtr&& stream)
		: m_stream(std::move(stream))
	{
	}

public:
	IInputPtr m_stream;
};

class OutputStreamDecorator : public IOutputDataStream
{
public:
	void WriteByte(uint8_t const data) override
	{
		m_stream->WriteByte(data);
	}

	void WriteBlock(const void* srcData, std::streamsize const size) override
	{
		m_stream->WriteBlock(srcData, size);
	}

	void Close() override
	{
		m_stream->Close();
	}

protected:
	explicit OutputStreamDecorator(IOutputPtr&& stream)
		: m_stream(std::move(stream))
	{
	}

public:
	IOutputPtr m_stream;
};