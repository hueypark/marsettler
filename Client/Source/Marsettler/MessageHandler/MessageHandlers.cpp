#include "MessageHandlers.h"

#include "Core/Log.h"
#include "Message/Message.h"
#include "MessageHandler/LoginResponseHandler.h"

void MessageHandlers::Handle(const Message* message)
{
	switch (message->ID())
	{
	case MessageID::LoginRes:
		LoginResponseHandler::Handle(message);
		break;
	default:
		LOG_PRINT("Handler is null. [MessageID: %d]", message->ID());
		break;
	}
}
