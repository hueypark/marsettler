#pragma once

#include "MessageBuilder.h"

// LoginBuilder 는 로그인 메시지 빌더입니다.
class LoginBuilder : public MessageBuilder
{
public:
	// 생성자
	LoginBuilder(const int64_t id);

	// Build 는 FlatBuffers 메시지를 만듭니다.
	virtual void Build(flatbuffers::FlatBufferBuilder& builder) const override;

	// Clone 은 메시지의 사본을 만듭니다.
	virtual std::unique_ptr<MessageBuilder> Clone() const override;

	// MessageID 는 메시지 아이디를 반환합니다.
	virtual MessageID ID() const override;

private:
	// id 는 고유한 아이디입니다.
	int64_t m_id;
};
