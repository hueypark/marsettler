#pragma once

#include <map>

class Actor;

class World
{
public:
	// 액터를 추가한다.
	void AddActor(std::shared_ptr<Actor> actor);

private:
	// 액터들
	std::map<int64, std::shared_ptr<Actor>> m_actors;
};
