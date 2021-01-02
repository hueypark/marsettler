#include "MsgLoginReqBuilder_generated.h"

#include "MsgLoginReq_generated.h"

MsgLoginReqBuilder::MsgLoginReqBuilder(const int64_t id) : m_id(id)
{
}

void MsgLoginReqBuilder::Build(flatbuffers::FlatBufferBuilder& builder) const
{
	auto loginReq = fbs::CreateMsgLoginReq(builder, m_id);
	builder.Finish(loginReq);
}

std::unique_ptr<MessageBuilder> MsgLoginReqBuilder::Clone() const
{
	return std::make_unique<MsgLoginReqBuilder>(m_id);
}

MessageID MsgLoginReqBuilder::ID() const
{
	return MessageID::LoginRes;
}
