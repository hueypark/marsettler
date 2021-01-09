#include "MsgMovePushBuilder.h"

#include "MsgMovePush_generated.h"
#include "MsgVector_generated.h"

MsgMovePushBuilder::MsgMovePushBuilder(const int64_t id, const fbs::MsgVector& location)
	: m_id(id), m_location(location)
{
}

void MsgMovePushBuilder::Build(flatbuffers::FlatBufferBuilder& builder) const
{
	const auto movePush = fbs::CreateMsgMovePush(builder, m_id, &m_location);
	builder.Finish(movePush);
}

std::unique_ptr<MessageBuilder> MsgMovePushBuilder::Clone() const
{
	return std::make_unique<MsgMovePushBuilder>(m_id, m_location);
}

MessageID MsgMovePushBuilder::ID() const
{
	return MessageID::MovePush;
}
