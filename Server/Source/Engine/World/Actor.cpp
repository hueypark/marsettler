#include "Actor.h"

Actor::Actor(const int64_t id, const Vector& position) : m_id(id), m_position(position)
{
}

int64_t Actor::ID() const
{
	return m_id;
}

Vector Actor::Position() const
{
	return m_position;
}

void Actor::Write(const MessageBuilder& builder)
{
	// TODO(jaewan): 메시지 실제로 보내게 수정
}
