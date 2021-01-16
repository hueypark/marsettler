#include "Message.h"

Message::Message(const MessageID& id, const int32_t& size) : m_id(id)
{
	m_body.resize(size);
}

uint8_t* Message::Data()
{
	return m_body.data();
}

const uint8_t* Message::Data() const
{
	return m_body.data();
}

MessageID Message::ID() const
{
	return m_id;
}

const int32_t Message::Size() const
{
	return (int32_t)(m_body.size());
}
