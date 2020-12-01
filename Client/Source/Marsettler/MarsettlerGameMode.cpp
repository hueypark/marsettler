// Copyright Epic Games, Inc. All Rights Reserved.

#include "MarsettlerGameMode.h"

#include "MarsettlerCharacter.h"
#include "MarsettlerPlayerController.h"
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
}
