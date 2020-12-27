#pragma once

#include "MessageID.h"

#include <cinttypes>
#include <vector>

// Message 는 메시지를 표현합니다.
class Message
{
public:
	// 생성자
	Message(const MessageID& id, const int32_t& size);

	// Data 는 데이터를 반환합니다.
	uint8_t* Data();

	// Data 는 데이터를 반환합니다.
	const uint8_t* Data() const;

	// ID 는 메시지 ID를 반환합니다.
	MessageID ID() const;

	// Size 는 데이터의 크기를 반환합니다.
	const int32 Size() const;

private:
	MessageID m_id;
	std::vector<uint8_t> m_body;
};