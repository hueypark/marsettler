#include "World.h"

#include "Engine/Math/Vector.h"
#include "Engine/Worker.h"
#include "Engine/World/NetworkActor.h"
#include "Message/MsgLoginResBuilder_generated.h"

World::World()
{
	m_worker = new Worker;
}

World::~World()
{
	delete m_worker;
}

void World::LoginActor(const int64_t id, const std::weak_ptr<Connection>& connection)
{
	m_worker->AddWork(
		[this, id, connection]()
		{
			NetworkActor::Ptr networkActor = nullptr;

			Actor::Ptr actor = _GetActor(id);
			if (actor)
			{
				networkActor = std::dynamic_pointer_cast<NetworkActor>(actor);
			}
			else
			{
				networkActor = _NewNetworkActor(id);
				actor = networkActor;
			}

			networkActor->SetConnection(connection);

			MsgLoginResBuilder loginRes(id, networkActor->Position());
			networkActor->Write(loginRes);
		});
}

void World::Tick()
{
	m_worker->Tick();
}

std::shared_ptr<Actor> World::_GetActor(const int64_t id)
{
	auto iter = m_actors.find(id);
	if (iter == m_actors.end())
	{
		return nullptr;
	}

	return iter->second;
}

std::shared_ptr<NetworkActor> World::_NewNetworkActor(const int64_t id)
{
	NetworkActor::Ptr actor = std::make_shared<NetworkActor>(id, Vector());

	m_actors[id] = actor;

	return actor;
}
