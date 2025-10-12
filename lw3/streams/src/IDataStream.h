#pragma once

#include <ios>
#include <memory>

class IOutputDataStream
{
public:
	virtual void WriteByte(uint8_t data) = 0;
	virtual void WriteBlock(const void* srcData, std::streamsize size) = 0;
	virtual void Close() = 0;

	virtual ~IOutputDataStream() = default;
};

using IOutputPtr = std::unique_ptr<IOutputDataStream>;

class IInputDataStream
{
public:
	[[nodiscard]] virtual bool IsEOF() const = 0;
	virtual uint8_t ReadByte() = 0;
	virtual std::streamsize ReadBlock(void* dstBuffer, std::streamsize size) = 0;

	virtual ~IInputDataStream() = default;
};

using IInputPtr = std::unique_ptr<IInputDataStream>;