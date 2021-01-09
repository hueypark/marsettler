#include "Actor.h"

Actor::Actor(const int64 id, const FVector location) : m_id(id), m_location(location)
{
}

bool Actor::GetLocationUpdated() const
{
	return m_locationUpdated;
}

int64 Actor::ID() const
{
	return m_id;
}

FVector Actor::Location() const
{
	return m_location;
}

void Actor::SetLocation(const FVector location)
{
	m_location = location;

	m_locationUpdated = true;
}

void Actor::SetLocationUpdated(const bool updated)
{
	m_locationUpdated = updated;
}
