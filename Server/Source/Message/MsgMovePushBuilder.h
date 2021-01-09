#pragma once

#include "Engine/Math/Vector.h"
#include "Message/MsgVector_generated.h"
#include "MessageBuilder.h"

// MsgMovePush 빌더
class MsgMovePushBuilder : public MessageBuilder
{
public:
	// 생성자
	MsgMovePushBuilder(const int64_t id, const fbs::MsgVector& location);

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
	const fbs::MsgVector m_location;
};
