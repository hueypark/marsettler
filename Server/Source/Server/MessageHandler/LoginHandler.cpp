#include "LoginHandler.h"

#include "Network/Connection.h"

#include <Message/LoginResponse_generated.h>
#include <Message/Login_generated.h>
#include <Message/Message.h>

void LoginHandler::Handle(const Connection* conn, const Message* message)
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

	// TODO(jaewan): 액터 생성 혹은 기존 액터 연결

	// TODO(jaewan): 결과 메시지 전달
}
