#include "MessageHandler.h"

#include "Message/Message.h"
#include "MessageHandler/LoginHandler.h"

#include <iostream>

void MessageHandler::Handle(const Connection* conn, const Message* message)
{
	// 플랫버퍼 내부에서 발생하는 예외를 잡기위해 try 문 추가했습니다.
	try
	{
		switch (message->ID())
		{
		case MessageID::Login:
			{
				LoginHandler::Handle(conn, message);
			}
			break;
		default:
			{
				std::cerr << "Unhandled message. ID: " << int(message->ID()) << std::endl;
			}
			break;
		}
	}
	catch (std::exception& e)
	{
		std::cerr << e.what() << std::endl;
	}
}
