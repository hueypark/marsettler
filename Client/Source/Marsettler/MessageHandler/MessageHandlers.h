#pragma once

class Message;
enum class MessageID;

class MessageHandlers
{
public:
	/// 메시지를 처리합니다.
	static void Handle(const Message* message);
};
