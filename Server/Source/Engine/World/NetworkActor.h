#pragma once

#include "Engine/World/Actor.h"

class Connection;

// 네트워크 액터
class NetworkActor : public Actor
{
public:
	using Ptr = std::shared_ptr<NetworkActor>;

private:
	using Super = Actor;

public:
	// 생성자
	NetworkActor(const int64_t id, const Vector& position);

	// 연결을 설정합니다.
	void SetConnection(const std::weak_ptr<Connection>& connection);

	// Write 는 메시지를 씁니다.
	void Write(const MessageBuilder& builder);

private:
	// 연결
	std::weak_ptr<Connection> m_connection;
};
