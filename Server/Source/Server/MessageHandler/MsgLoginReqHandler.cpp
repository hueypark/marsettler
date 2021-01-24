#include "MsgLoginReqHandler.h"

#include "Engine/Network/Connection.h"
#include "Engine/World/Actor.h"
#include "Engine/World/World.h"
#include "Message/Message.h"
#include "Message/MsgLoginReq_generated.h"
#include "Server/Context.h"
#include "bsoncxx/builder/stream/document.hpp"
#include "mongocxx/client.hpp"
#include "mongocxx/instance.hpp"

void MsgLoginReqHandler::Handle(Connection* conn, const Message* message)
{
	mongocxx::instance instance{};
	// mongocxx::uri uri("mongodb://localhost:27017");
	// mongocxx::client client{uri};

	const fbs::MsgLoginReq* loginReq = fbs::GetMsgLoginReq(message->Data());
	if (!loginReq)
	{
		// TODO(jaewan): 로그 기록

		return;
	}

	// TODO(jeawan): 인증 추가

	int64_t id = loginReq->ID();

	if (!id)
	{
		// TODO(jaewan): 물리장비가 여러개여도 고유한 ID 생성되게 수정
		static int64_t newID = 0;
		++newID;

		id = newID;
	}

	Connection::Ptr connPtr = conn->shared_from_this();
	Context::Instance.GameWorld->LoginActor(id, connPtr);
}
