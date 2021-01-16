#pragma once

#include "Engine/Math/Vector.h"
#include "MessageBuilder.h"

// MsgActorPushBuilder 는 MsgActorPush 빌더입니다.
class MsgActorPushBuilder : public MessageBuilder
{
public:
	// 생성자
	MsgActorPushBuilder(const int64_t id, const Vector& location);

	// Build 는 FlatBuffers 메시지를 만듭니다.
	virtual void Build(flatbuffers::FlatBufferBuilder& builder) const override;

	// Clone 은 메시지의 사본을 만듭니다.
	virtual std::unique_ptr<MessageBuilder> Clone() const override;

	// MessageID 는 메시지 아이디를 반환합니다.
	virtual MessageID ID() const override;

private:
	// ID
	int64_t m_id;

	// 위치
	Vector m_location;
};
