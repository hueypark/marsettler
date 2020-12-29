#pragma once

class Actor
{
public:
	// 생성자
	Actor(const int64 id);

	// ID를 반환한다.
	int64 ID() const;

private:
	// ID
	int64 m_id;
};
