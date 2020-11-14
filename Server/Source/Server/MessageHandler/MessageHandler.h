#pragma once

class Message;

class MessageHandler
{
public:
	// Handle 은 메시지를 처리합니다.
	void Handle(const Message* message);
};
