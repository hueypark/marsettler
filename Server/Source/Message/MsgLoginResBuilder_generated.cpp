#include "MsgLoginResBuilder_generated.h"

#include "MsgLoginRes_generated.h"
#include "MsgVector_generated.h"

MsgLoginResBuilder::MsgLoginResBuilder(const int64_t id, const Vector& location) : m_id(id), m_location(location)
{
}

void MsgLoginResBuilder::Build(flatbuffers::FlatBufferBuilder& builder) const
{
	fbs::MsgVector location(m_location.X, m_location.Y);

	fbs::MsgActor actor(m_id, location);

	auto loginResponse = fbs::CreateMsgLoginRes(builder, &actor);
	builder.Finish(loginResponse);
}

std::unique_ptr<MessageBuilder> MsgLoginResBuilder::Clone() const
{
	return std::make_unique<MsgLoginResBuilder>(m_id, m_location);
}

MessageID MsgLoginResBuilder::ID() const
{
	return MessageID::LoginRes;
}
