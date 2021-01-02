#pragma once

#include "Engine/Math/Vector.h"

#include <cstdint>

// 액터
class Actor
{
public:
	// 생성자
	Actor(const int64_t id, const Vector& position);

	// 위치를 반환한다.
	Vector Position() const;

private:
	// ID
	int64_t m_id;

	// 위치
	Vector m_position;
};
