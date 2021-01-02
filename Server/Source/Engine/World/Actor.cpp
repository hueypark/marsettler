#include "Actor.h"

Actor::Actor(const int64_t id, const Vector& position) : m_id(id), m_position(position)
{
}

Vector Actor::Position() const
{
	return m_position;
}
