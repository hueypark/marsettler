#include "Actor.h"

Actor::Actor(const int64_t id, const Vector& position) : m_id(id), m_position(position)
{
}

Actor::~Actor()
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
