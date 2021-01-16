#include "MessageHandler.h"

#include "Message/Message.h"
#include "Server/MessageHandler/MsgLoginReqHandler.h"
#include "Server/MessageHandler/MsgMoveReqHandler.h"

#include <iostream>

void MessageHandler::Handle(Connection* conn, const Message* message)
{
	// 플랫버퍼 내부에서 발생하는 예외를 잡기위해 try 문 추가했습니다.
	try
	{
		switch (message->ID())
		{
		case MessageID::LoginReq:
			{
				MsgLoginReqHandler::Handle(conn, message);
			}
			break;
		case MessageID::MoveReq:
			{
				MsgMoveReqHandler::Handle(conn, message);
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
