#include "Message.h"

Message::Message(const int32_t& id, const int32_t& size) : m_id(id)
{
	m_body.resize(size);
}

uint8_t* Message::Data()
{
	return m_body.data();
}

const int32_t Message::Size()
{
	return m_body.size();
}
