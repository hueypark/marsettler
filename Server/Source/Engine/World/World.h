#pragma once

#include <map>
#include <memory>

class Actor;
class Connection;
class NetworkActor;
class Worker;

// 월드
class World
{
public:
	// 생성자
	World();

	// 소멸자
	virtual ~World();

	// 로그인한다.
	void LoginActor(const int64_t id, const std::weak_ptr<Connection>& connection);

	// 매 틱마다 호출됩니다.
	void Tick();

private:
	// 액터를 반환한다.
	std::shared_ptr<Actor> _GetActor(const int64_t id);

	// 새 네트워크 액터를 만든다.
	std::shared_ptr<NetworkActor> _NewNetworkActor(const int64_t id);

private:
	// 액터들
	std::map<int64_t, std::shared_ptr<Actor>> m_actors;

	// 워커
	Worker* m_worker;
};
