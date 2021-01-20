#include "NetworkActor.h"

#include "Engine/Network/Connection.h"

NetworkActor::NetworkActor(const int64_t id, const Vector& position) : Super(id, position)
{
}

void NetworkActor::SetConnection(const std::weak_ptr<Connection>& connection)
{
	m_connection = connection;
}

void NetworkActor::Write(const MessageBuilder& builder)
{
	Connection::Ptr connection = m_connection.lock();
	if (!connection)
	{
		// TODO(jaewan): 로그 기록

		return;
	}

	connection->Write(builder);
}
