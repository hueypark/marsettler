#pragma once

class Connection;
class Message;
enum class MessageID;

class MessageHandler
{
public:
	// Handle 은 메시지를 처리합니다.
	static void Handle(Connection* conn, const Message* message);
};
