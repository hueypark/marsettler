#pragma once

#include "Components/ActorComponent.h"
#include "CoreMinimal.h"

#include <flatbuffers/flatbuffers.h>

#include "NetworkComponent.generated.h"

class MessageBuilder;
class FSocket;

UCLASS(ClassGroup = (Custom), meta = (BlueprintSpawnableComponent))
class MARSETTLER_API UNetworkComponent : public UActorComponent
{
	GENERATED_BODY()

public:
	// 생성자
	UNetworkComponent();

protected:
	// 게임 시작시 호출됩니다.
	virtual void BeginPlay() override;

	// 게임 종료시 호출됩니다.
	virtual void EndPlay(const EEndPlayReason::Type endPlayReason) override;

public:
	// 매 프레임 호출됩니다.
	virtual void TickComponent(float DeltaTime,
		ELevelTick TickType,
		FActorComponentTickFunction* ThisTickFunction) override;

	// 메시지를 씁니다.
	void WriteMessage(const MessageBuilder& builder);

private:
	// 서버와 연결을 끊습니다.
	void _CloseFromServer();

	// 서버에 연결합니다.
	bool _ConnectToServer();

	// _WriteLogin 은 로그인 메시지를 보냅니다.
	void _WriteLogin();

private:
	FSocket* m_socket;

	std::vector<uint8> m_messageInHeaderBuf;

	std::vector<uint8> m_messageOutHeaderBuf;
	flatbuffers::FlatBufferBuilder m_messageOutBuilder;
};
