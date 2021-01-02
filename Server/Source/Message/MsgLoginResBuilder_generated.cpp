#include "MsgLoginResBuilder_generated.h"

#include "MsgLoginRes_generated.h"
#include "MsgVector_generated.h"

MsgLoginResBuilder::MsgLoginResBuilder(const int64_t id, const Vector& position) : m_id(id), m_position(position)
{
}

void MsgLoginResBuilder::Build(flatbuffers::FlatBufferBuilder& builder) const
{
	auto position = fbs::MsgVector(m_position.X, m_position.Y);

	auto actor = fbs::CreateMsgActor(builder, m_id, &position);

	auto loginResponse = fbs::CreateMsgLoginRes(builder, actor);
	builder.Finish(loginResponse);
}

std::unique_ptr<MessageBuilder> MsgLoginResBuilder::Clone() const
{
	return std::make_unique<MsgLoginResBuilder>(m_id, m_position);
}

MessageID MsgLoginResBuilder::ID() const
{
	return MessageID::LoginRes;
}
