#include "LoginHandler.h"

#include "Context.h"
#include "Engine/World/World.h"
#include "Message/LoginResponseBuilder_generated.h"
#include "Message/Login_generated.h"
#include "Message/Message.h"
#include "Network/Connection.h"

void LoginHandler::Handle(Connection* conn, const Message* message)
{
	const fbs::Login* login = fbs::GetLogin(message->Data());
	if (!login)
	{
		// TODO(jaewan): 로그 기록

		return;
	}

	// TODO(jeawan): 인증 추가

	int64_t id = login->ID();

	if (!id)
	{
		// TODO(jaewan): 물리장비가 여러개여도 고유한 ID 생성되게 수정
		static int64_t newID = 0;
		++newID;

		id = newID;
	}

	std::shared_ptr<Actor> actor = Context::Instance.GameWorld->GetOrNewActor(id);

	// TODO(jaewan): 결과 메시지 전달
	LoginResponseBuilder loginResponse(id);
	conn->Write(loginResponse);
}
