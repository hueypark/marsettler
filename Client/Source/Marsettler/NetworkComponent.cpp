#include "NetworkComponent.h"

#include "Networking.h"

#include <Marsettler/Message/LoginBuilder_generated.h>

UNetworkComponent::UNetworkComponent() : m_messageInHeaderBuf(8), m_messageOutHeaderBuf(8)
{
	PrimaryComponentTick.bCanEverTick = true;
}

void UNetworkComponent::BeginPlay()
{
	Super::BeginPlay();

	m_socket = nullptr;

	if (_ConnectToServer())
	{
		_WriteLogin();
	}
}

void UNetworkComponent::EndPlay(const EEndPlayReason::Type endPlayReason)
{
	Super::EndPlay(endPlayReason);

	_CloseFromServer();
}

void UNetworkComponent::TickComponent(
	float DeltaTime, ELevelTick TickType, FActorComponentTickFunction* ThisTickFunction)
{
	Super::TickComponent(DeltaTime, TickType, ThisTickFunction);

	if (!m_socket)
	{
		return;
	}

	int32 bytesReadAll = 0;
	while (true)
	{
		if (bytesReadAll == m_messageInHeaderBuf.size())
		{
			break;
		}

		int32 bytesRead = 0;
		if (!m_socket->Recv(m_messageInHeaderBuf.data() + bytesReadAll,
				m_messageInHeaderBuf.size() - bytesReadAll, bytesRead))
		{
			_CloseFromServer();

			return;
		}

		bytesReadAll += bytesRead;
	}

	MessageID messageID;
	std::memcpy(&messageID, &m_messageInHeaderBuf[0], 4);

	int32 messageSize;
	std::memcpy(&messageSize, &m_messageInHeaderBuf[4], 4);
}

void UNetworkComponent::WriteMessage(const MessageBuilder& builder)
{
	builder.Build(m_messageOutBuilder);

	MessageID messageID = builder.ID();
	std::memcpy(&m_messageOutHeaderBuf[0], &messageID, 4);

	int32 messageSize = m_messageOutBuilder.GetSize();
	std::memcpy(&m_messageOutHeaderBuf[4], &messageSize, 4);

	int32 bytesSent;
	bool result =
		m_socket->Send(m_messageOutHeaderBuf.data(), m_messageOutHeaderBuf.size(), bytesSent);
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
	m_socket = ISocketSubsystem::Get(PLATFORM_SOCKETSUBSYSTEM)
				   ->CreateSocket(NAME_Stream, TEXT("Socket"), false);

	TSharedRef<FInternetAddr> addr =
		ISocketSubsystem::Get(PLATFORM_SOCKETSUBSYSTEM)->CreateInternetAddr();
	addr->SetIp(FIPv4Address(127, 0, 0, 1).Value);
	addr->SetPort(8080);

	return m_socket->Connect(*addr);
}

void UNetworkComponent::_WriteLogin()
{
	LoginBuilder loginBuilder(7);
	WriteMessage(loginBuilder);
}
