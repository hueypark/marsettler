// Copyright Epic Games, Inc. All Rights Reserved.

#pragma once

#include "CoreMinimal.h"
#include "GameFramework/GameModeBase.h"

#include "MarsettlerGameMode.generated.h"

UCLASS(minimalapi)
class AMarsettlerGameMode : public AGameModeBase
{
	GENERATED_BODY()

public:
	AMarsettlerGameMode();

	virtual void StartPlay() override;

private:
	// _ConnectToServer 는 서버에 연결합니다.
	bool _ConnectToServer();

	// _WriteLogin 은 로그인 메시지를 보냅니다.
	void _WriteLogin();
};
