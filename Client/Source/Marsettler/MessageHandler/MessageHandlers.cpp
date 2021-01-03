#include "MessageHandlers.h"

#include "Core/Log.h"
#include "Message/Message.h"
#include "MessageHandler/MsgLoginResHandler.h"

void MessageHandlers::Handle(const Message* message)
{
	switch (message->ID())
	{
	case MessageID::LoginRes:
		MsgLoginResHandler::Handle(message);
		break;
	default:
		LOG_PRINT("Handler is null. [MessageID: %d]", message->ID());
		break;
	}
}
