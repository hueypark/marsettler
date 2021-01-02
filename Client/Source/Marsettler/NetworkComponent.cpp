#include "NetworkComponent.h"

#include "Log/Log.h"
#include "Message/Message.h"
#include "Message/MsgLoginReqBuilder_generated.h"
#include "MessageHandler/MessageHandlers.h"
#include "Networking.h"

UNetworkComponent::UNetworkComponent() : m_stop(false), m_messageInHeaderBuf(8), m_messageOutHeaderBuf(8)
{
	m_socket = nullptr;

	if (_ConnectToServer())
	{
		_WriteLogin();
	}

	m_thread = FRunnableThread::Create(this, TEXT("NetworkComponent"), 0, TPri_BelowNormal);
}

UNetworkComponent::~UNetworkComponent()
{
	_CloseFromServer();
}

uint32 UNetworkComponent::Run()
{
	while (true)
	{
		FPlatformProcess::Sleep(0.1);

		if (m_stop)
		{
			return 0;
		}

		_Tick();
	}

	return 0;
}

void UNetworkComponent::Stop()
{
	m_stop = true;

	if (m_thread)
	{
		m_thread->WaitForCompletion();
		m_thread = nullptr;
	}
}

void UNetworkComponent::WriteMessage(const MessageBuilder& builder)
{
	builder.Build(m_messageOutBuilder);

	MessageID messageID = builder.ID();
	std::memcpy(&m_messageOutHeaderBuf[0], &messageID, 4);

	int32 messageSize = m_messageOutBuilder.GetSize();
	std::memcpy(&m_messageOutHeaderBuf[4], &messageSize, 4);

	int32 bytesSent;
	bool result = m_socket->Send(m_messageOutHeaderBuf.data(), m_messageOutHeaderBuf.size(), bytesSent);
	if (!result)
	{
		_CloseFromServer();

		return;
	}

	result = m_socket->Send(m_messageOutBuilder.GetBufferPointer(), messageSize, bytesSent);
	if (!result)
	{
		_CloseFromServer();

		return;
	}
}

void UNetworkComponent::_CloseFromServer()
{
	if (!m_socket)
	{
		return;
	}

	m_socket->Close();
	ISocketSubsystem::Get(PLATFORM_SOCKETSUBSYSTEM)->DestroySocket(m_socket);
	m_socket = nullptr;
}

bool UNetworkComponent::_ConnectToServer()
{
	m_socket = ISocketSubsystem::Get(PLATFORM_SOCKETSUBSYSTEM)->CreateSocket(NAME_Stream, TEXT("Socket"), false);

	TSharedRef<FInternetAddr> addr = ISocketSubsystem::Get(PLATFORM_SOCKETSUBSYSTEM)->CreateInternetAddr();
	addr->SetIp(FIPv4Address(127, 0, 0, 1).Value);
	addr->SetPort(8080);

	return m_socket->Connect(*addr);
}

void UNetworkComponent::_Tick()
{
	if (!m_socket)
	{
		return;
	}

	int32 bytesReadAll = 0;
	uint32 pendingDataSize = 0;
	while (true)
	{
		if (!m_socket->HasPendingData(pendingDataSize))
		{
			if (m_stop)
			{
				_CloseFromServer();

				return;
			}

			continue;
		}

		int32 bytesRead = 0;
		if (!m_socket->Recv(
				m_messageInHeaderBuf.data() + bytesReadAll, m_messageInHeaderBuf.size() - bytesReadAll, bytesRead))
		{
			_CloseFromServer();

			return;
		}

		bytesReadAll += bytesRead;

		if (bytesReadAll == m_messageInHeaderBuf.size())
		{
			break;
		}
	}

	MessageID messageID;
	std::memcpy(&messageID, &m_messageInHeaderBuf[0], 4);

	int32 messageSize;
	std::memcpy(&messageSize, &m_messageInHeaderBuf[4], 4);

	bytesReadAll = 0;
	Message message(messageID, messageSize);
	while (true)
	{
		if (!m_socket->HasPendingData(pendingDataSize))
		{
			if (m_stop)
			{
				_CloseFromServer();

				return;
			}

			continue;
		}

		int32 bytesRead = 0;
		if (!m_socket->Recv(message.Data() + bytesReadAll, message.Size() - bytesReadAll, bytesRead))
		{
			_CloseFromServer();

			return;
		}

		bytesReadAll += bytesRead;

		if (bytesReadAll == messageSize)
		{
			break;
		}
	}

	MessageHandlers::Handle(&message);
}

void UNetworkComponent::_WriteLogin()
{
	MsgLoginReqBuilder loginReq(0);
	WriteMessage(loginReq);
}
