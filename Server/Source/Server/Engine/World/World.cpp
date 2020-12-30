#include "World.h"

#include "Engine/World/Actor.h"

std::shared_ptr<Actor> World::GetOrNewActor(const int64_t id)
{
	if (m_actors.find(id) == m_actors.end())
	{
		m_actors[id] = std::make_shared<Actor>(id);
	}

	return m_actors[id];
}
