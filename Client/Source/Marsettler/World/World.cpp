#include "World.h"

#include "World/Actor.h"

void World::AddActor(std::shared_ptr<Actor> actor)
{
	m_actors[actor->ID()] = actor;
}
