#include "Actor.h"

Actor::Actor(const int64 id) : m_id(id)
{
}

int64 Actor::ID() const
{
	return m_id;
}
