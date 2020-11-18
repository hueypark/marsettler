#pragma once

#include "MessageID.h"

#include <flatbuffers/flatbuffers.h>

// MessageBuilder 는 메시지 빌더 인터페이스입니다.
class MessageBuilder
{
public:
	// Build 는 FlatBuffers 메시지를 만듭니다.
	virtual void Build(flatbuffers::FlatBufferBuilder& builder) const = 0;

	// Clone 은 메시지의 사본을 만듭니다.
	virtual std::unique_ptr<MessageBuilder> Clone() const = 0;

	// MessageID 는 메시지 아이디를 반환합니다.
	virtual MessageID ID() const = 0;
};