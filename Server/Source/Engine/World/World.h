#pragma once

#include <map>
#include <memory>

class Actor;
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
	void LoginActor(const int64_t id);

	// 매 틱마다 호출됩니다.
	void Tick();

private:
	// 액터가 있으면 바로 반환하고 없으면 생성해 반환한다.
	std::shared_ptr<Actor> _GetOrNewActor(const int64_t id);

private:
	// 액터들
	std::map<int64_t, std::shared_ptr<Actor>> m_actors;

	// 워커
	Worker* m_worker;
};
