#include "LoginResponseBuilder_generated.h"

#include "LoginResponse_generated.h"

LoginResponseBuilder::LoginResponseBuilder(const int64_t id) : m_id(id)
{
}

void LoginResponseBuilder::Build(flatbuffers::FlatBufferBuilder& builder) const
{
	auto loginResponse = fbs::CreateLoginResponse(builder, m_id);
	builder.Finish(loginResponse);
}

std::unique_ptr<MessageBuilder> LoginResponseBuilder::Clone() const
{
	return std::make_unique<LoginResponseBuilder>(m_id);
}

MessageID LoginResponseBuilder::ID() const
{
	return MessageID::LoginResponse;
}
