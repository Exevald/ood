#include "TempFilesGenerator.h"
#include "compress-streams/CompressingOutputStream.h"
#include "compress-streams/Compressor.h"
#include "compress-streams/DecompressingInputStream.h"
#include "crypt-streams/DecryptingInputStream.h"
#include "crypt-streams/EncryptingOutputStream.h"
#include "file-streams/FileInputStream.h"
#include "file-streams/FileOutputStream.h"
#include "memory-streams/MemoryInputStream.h"
#include "memory-streams/MemoryOutputStream.h"

#include <filesystem>
#include <fstream>
#include <gtest/gtest.h>
#include <ios>
#include <random>
#include <utility>

TEST(MemoryStream, TestReadWriteByte)
{
	MemoryOutputStream output;
	output.WriteByte(42);
	output.WriteByte(255);

	MemoryInputStream input(std::vector<uint8_t>{ 42, 255 });
	EXPECT_FALSE(input.IsEOF());
	EXPECT_EQ(input.ReadByte(), 42);
	EXPECT_EQ(input.ReadByte(), 255);
	EXPECT_TRUE(input.IsEOF());
}

TEST(MemoryStream, TestReadWriteBlock)
{
	const std::vector<uint8_t> data = { 1, 2, 3, 4, 5 };
	MemoryOutputStream output;
	output.WriteBlock(data.data(), static_cast<long>(data.size()));

	MemoryInputStream input(output.GetData());
	std::vector<uint8_t> buffer(5);
	const auto read = input.ReadBlock(buffer.data(), 5);
	EXPECT_EQ(read, 5);
	EXPECT_EQ(buffer, data);
	EXPECT_TRUE(input.IsEOF());
}

TEST(FileStream, TestReadWriteByte)
{
	auto tempFile = TempFileGenerator::CreateTempFile({});
	TempFileGuard guard(tempFile);

	FileOutputStream out(tempFile);
	out.WriteByte(100);
	out.WriteByte(200);
	out.Close();

	FileInputStream in(tempFile);
	EXPECT_EQ(in.ReadByte(), 100);
	EXPECT_EQ(in.ReadByte(), 200);

	EXPECT_THROW(in.ReadByte(), std::ios_base::failure);
}

TEST(FileStream, TestReadWriteBlock)
{
	std::vector<uint8_t> data = { 1, 2, 3, 4, 5, 6, 7, 8 };
	auto tempFile = TempFileGenerator::CreateTempFile({});
	TempFileGuard guard(tempFile);

	FileOutputStream out(tempFile);
	out.WriteBlock(data.data(), static_cast<long>(data.size()));
	out.Close();

	FileInputStream in(tempFile);
	std::vector<uint8_t> buffer(8);
	auto read = in.ReadBlock(buffer.data(), 8);
	EXPECT_EQ(read, 8);
	EXPECT_EQ(buffer, data);
}

TEST(Substitution, TestEncryptionDecryptionConsistency)
{
	constexpr uint32_t key = 12345;
	const auto encTable = SubstitutionTable::GenerateEncryptionTable(key);
	const auto decTable = SubstitutionTable::GenerateDecryptionTable(encTable);

	for (int b = 0; b < 256; ++b)
	{
		const uint8_t encrypted = encTable[b];
		uint8_t decrypted = decTable[encrypted];
		EXPECT_EQ(decrypted, static_cast<uint8_t>(b));
	}
}

TEST(StreamDecorator, TestEncryptDecryptRoundtrip)
{
	const std::vector<uint8_t> original = { 0, 1, 2, 255, 128, 64 };
	constexpr uint32_t key = 999;

	auto memOut = std::make_unique<MemoryOutputStream>();
	const MemoryOutputStream* rawMemOut = memOut.get();

	const auto encryptOut = std::make_unique<EncryptingOutputStream>(std::move(memOut), key);
	encryptOut->WriteBlock(original.data(), static_cast<long>(original.size()));
	encryptOut->Close();

	const auto& encryptedData = rawMemOut->GetData();

	auto memIn = std::make_unique<MemoryInputStream>(encryptedData);
	const auto decryptIn = std::make_unique<DecryptingInputStream>(std::move(memIn), key);

	std::vector<uint8_t> restored(original.size());
	decryptIn->ReadBlock(restored.data(), static_cast<long>(restored.size()));

	EXPECT_EQ(restored, original);
}

TEST(RLE, TestCompressDecompress)
{
	const std::vector<uint8_t> original = { 1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 4 };
	const auto compressed = Compressor::RLECompress(original.data(), original.size());
	const auto decompressed = Compressor::RLEDecompress(compressed.data(), compressed.size());
	EXPECT_EQ(decompressed, original);
}

TEST(RLE, TestSingleBytes)
{
	const std::vector<uint8_t> original = { 1, 2, 3, 4, 5 };
	const auto compressed = Compressor::RLECompress(original.data(), original.size());
	ASSERT_EQ(compressed.size(), 10);
	for (size_t i = 0; i < 5; ++i)
	{
		EXPECT_EQ(compressed[i * 2], 1);
		EXPECT_EQ(compressed[i * 2 + 1], static_cast<uint8_t>(i + 1));
	}
	const auto decompressed = Compressor::RLEDecompress(compressed.data(), compressed.size());
	EXPECT_EQ(decompressed, original);
}

TEST(StreamDecorator, TestCompressDecompressRoundtrip)
{
	const std::vector<uint8_t> data = { 5, 5, 5, 5, 6, 7, 7, 8 };

	auto memOut = std::make_unique<MemoryOutputStream>();
	const MemoryOutputStream* rawMemOut = memOut.get();

	const auto compOut = std::make_unique<CompressingOutputStream>(std::move(memOut));
	compOut->WriteBlock(data.data(), static_cast<std::streamsize>(data.size()));
	compOut->Close();

	const auto& compressed = rawMemOut->GetData();

	auto in = std::make_unique<MemoryInputStream>(compressed);
	const auto decompIn = std::make_unique<DecompressingInputStream>(std::move(in));

	std::vector<uint8_t> result(data.size());
	decompIn->ReadBlock(result.data(), static_cast<std::streamsize>(data.size()));
	EXPECT_EQ(result, data);
}