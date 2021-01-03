#pragma once

#include "CoreMinimal.h"

#include <flatbuffers/flatbuffers.h>

class MessageBuilder;
class FRunnableThread;
class FSocket;

class NetworkComponent : public FRunnable
{
public:
	// 생성자
	NetworkComponent();

	// 소멸자
	virtual ~NetworkComponent();

	// 실행한다.
	virtual uint32 Run() override;

	// 중단한다.
	virtual void Stop();

	// 메시지를 씁니다.
	void WriteMessage(const MessageBuilder& builder);

private:
	// 서버와 연결을 끊습니다.
	void _CloseFromServer();

	// 서버에 연결합니다.
	bool _ConnectToServer();

	// 매 프레임 호출됩니다.
	void _Tick();

	// _WriteLogin 은 로그인 메시지를 보냅니다.
	void _WriteLogin();

private:
	FRunnableThread* m_thread;

	bool m_stop;

	FSocket* m_socket;

	std::vector<uint8> m_messageInHeaderBuf;

	std::vector<uint8> m_messageOutHeaderBuf;
	flatbuffers::FlatBufferBuilder m_messageOutBuilder;
};
