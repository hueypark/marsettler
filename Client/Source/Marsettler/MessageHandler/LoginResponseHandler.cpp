#include "LoginResponseHandler.h"

#include "Log/Log.h"
#include "Message/LoginResponse_generated.h"
#include "Message/Message.h"
#include "Message/MessageID.h"
#include "MessageHandler/MessageHandlers.h"

void LoginResponseHandler::Handle(const Message* message)
{
	const fbs::LoginResponse* loginResponse = fbs::GetLoginResponse(message->Data());
	if (!loginResponse)
	{
		return;
	}

	LOG_PRINT("[id: %lld]", loginResponse->ID());
}
