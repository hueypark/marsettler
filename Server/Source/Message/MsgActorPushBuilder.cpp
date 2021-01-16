#include "MsgActorPushBuilder.h"

#include "MsgLoginRes_generated.h"
#include "MsgVector_generated.h"

MsgActorPushBuilder::MsgActorPushBuilder(const int64_t id, const Vector& location) : m_id(id), m_location(location)
{
}

void MsgActorPushBuilder::Build(flatbuffers::FlatBufferBuilder& builder) const
{
	fbs::MsgVector location(m_location.X, m_location.Y);

	fbs::MsgActor actor(m_id, location);

	auto loginResponse = fbs::CreateMsgLoginRes(builder, &actor);
	builder.Finish(loginResponse);
}

std::unique_ptr<MessageBuilder> MsgActorPushBuilder::Clone() const
{
	return std::make_unique<MsgActorPushBuilder>(m_id, m_location);
}

MessageID MsgActorPushBuilder::ID() const
{
	return MessageID::LoginRes;
}
