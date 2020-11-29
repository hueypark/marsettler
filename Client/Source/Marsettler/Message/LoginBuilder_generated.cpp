#include "LoginBuilder_generated.h"

#include "Login_generated.h"

LoginBuilder::LoginBuilder(const int64_t id) : m_id(id)
{
}

void LoginBuilder::Build(flatbuffers::FlatBufferBuilder& builder) const
{
	auto loginResponse = fbs::CreateLogin(builder, m_id);
	builder.Finish(loginResponse);
}

std::unique_ptr<MessageBuilder> LoginBuilder::Clone() const
{
	return std::make_unique<LoginBuilder>(m_id);
}

MessageID LoginBuilder::ID() const
{
	return MessageID::LoginResponse;
}
