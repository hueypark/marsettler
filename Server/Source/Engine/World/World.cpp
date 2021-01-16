#include "World.h"

#include "Engine/Math/Vector.h"
#include "Engine/Worker.h"
#include "Engine/World/Actor.h"
#include "Message/MsgLoginResBuilder_generated.h"

World::World()
{
	m_worker = new Worker;
}

World::~World()
{
	delete m_worker;
}

void World::LoginActor(const int64_t id)
{
	m_worker->AddWork(
		[this, id]()
		{
			Actor::Ptr actor = _GetOrNewActor(id);

			MsgLoginResBuilder loginRes(id, actor->Position());
			actor->Write(loginRes);
		});
}

void World::Tick()
{
	m_worker->Tick();
}

std::shared_ptr<Actor> World::_GetOrNewActor(const int64_t id)
{
	if (m_actors.find(id) == m_actors.end())
	{
		m_actors[id] = std::make_shared<Actor>(id, Vector());
	}

	return m_actors[id];
}
