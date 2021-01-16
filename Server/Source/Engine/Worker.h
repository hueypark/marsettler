#pragma once

#include "boost/lockfree/queue.hpp"

#include <functional>

// Worker
//
// 작업을 받아 처리합니다.
class Worker
{
private:
	using Work = std::function<void()>;

public:
	/// 작업을 추가합니다.
	void AddWork(const Work& work);

	// 매 틱마다 호출됩니다.
	void Tick();

private:
	boost::lockfree::queue<Work*> m_works;
};
