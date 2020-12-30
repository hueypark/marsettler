#pragma once

#include <cstdint>

// 액터
class Actor
{
public:
	// 생성자
	Actor(const int64_t id);

private:
	// ID
	int64_t m_id;
};
