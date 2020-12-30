#pragma once

#include <memory>

class World;

// 전역 맥락을 가진 컨텍스트
class Context
{
private:
	// 생성자(private 접근 지정자 외부에서 접근 못하게 설정)
	Context();

public:
	// 인스턴스
	static Context Instance;

	// 월드
	std::shared_ptr<World> GameWorld;
};
