#pragma once

class Message;

class LoginResponseHandler
{
public:
	// 메시지를 처리합니다.
	static void Handle(const Message* message);
};
