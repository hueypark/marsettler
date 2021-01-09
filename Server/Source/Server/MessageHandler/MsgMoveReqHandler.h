#pragma once

class Connection;
class Message;

class MsgMoveReqHandler
{
public:
	// Handle 은 메시지를 처리합니다.
	static void Handle(Connection* conn, const Message* message);
};
