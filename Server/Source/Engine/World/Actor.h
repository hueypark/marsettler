#pragma once

#include "Engine/Math/Vector.h"

#include <cstdint>
#include <memory>

class MessageBuilder;

// 액터
class Actor
{
public:
	// 포인터
	using Ptr = std::shared_ptr<Actor>;

public:
	// 생성자
	Actor(const int64_t id, const Vector& position);

	// 소멸자
	virtual ~Actor();

	// ID를 반환한다.
	int64_t ID() const;

	// 위치를 반환한다.
	Vector Position() const;

private:
	// ID
	int64_t m_id;

	// 위치
	Vector m_position;
};
