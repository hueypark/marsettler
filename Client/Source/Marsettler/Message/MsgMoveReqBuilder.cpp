#include "MsgMoveReqBuilder.h"

#include "MsgMoveReq_generated.h"

MsgMoveReqBuilder::MsgMoveReqBuilder(const int64_t id, const fbs::MsgVector& location) : m_id(id), m_location(location)
{
}

void MsgMoveReqBuilder::Build(flatbuffers::FlatBufferBuilder& builder) const
{
	auto moveReq = fbs::CreateMsgMoveReq(builder, m_id, &m_location);
	builder.Finish(moveReq);
}

std::unique_ptr<MessageBuilder> MsgMoveReqBuilder::Clone() const
{
	return std::make_unique<MsgMoveReqBuilder>(m_id, m_location);
}

MessageID MsgMoveReqBuilder::ID() const
{
	return MessageID::MoveReq;
}
