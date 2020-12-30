#pragma once

#include <map>
#include <memory>

class Actor;

// 월드
class World
{
public:
	// 액터가 있으면 바로 반환하고 없으면 생성해 반환한다.
	std::shared_ptr<Actor> GetOrNewActor(const int64_t id);

private:
	// 액터들
	std::map<int64_t, std::shared_ptr<Actor>> m_actors;
};