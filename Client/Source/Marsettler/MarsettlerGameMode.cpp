// Copyright Epic Games, Inc. All Rights Reserved.

#include "MarsettlerGameMode.h"

#include "MarsettlerCharacter.h"
#include "MarsettlerPlayerController.h"
#include "Networking.h"
#include "UObject/ConstructorHelpers.h"

AMarsettlerGameMode::AMarsettlerGameMode()
{
	// use our custom PlayerController class
	PlayerControllerClass = AMarsettlerPlayerController::StaticClass();

	// set default pawn class to our Blueprinted character
	static ConstructorHelpers::FClassFinder<APawn> PlayerPawnBPClass(TEXT("/Game/TopDownCPP/Blueprints/TopDownCharacter"));
	if (PlayerPawnBPClass.Class != NULL)
	{
		DefaultPawnClass = PlayerPawnBPClass.Class;
	}
}

void AMarsettlerGameMode::StartPlay()
{
	Super::StartPlay();

	if (_ConnectToServer())
	{
		_WriteLogin();
	}
}

bool AMarsettlerGameMode::_ConnectToServer()
{
	FSocket* socket = ISocketSubsystem::Get(PLATFORM_SOCKETSUBSYSTEM)->CreateSocket(NAME_Stream, TEXT("Socket"), false);

	TSharedRef<FInternetAddr> addr = ISocketSubsystem::Get(PLATFORM_SOCKETSUBSYSTEM)->CreateInternetAddr();
	addr->SetIp(FIPv4Address(127, 0, 0, 1).Value);
	addr->SetPort(8080);

	return socket->Connect(*addr);
}

void AMarsettlerGameMode::_WriteLogin()
{
	// TODO(jaewan): 실제로 메시지 보내게 작업
}
