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

	virtual void EndPlay(const EEndPlayReason::Type EndPlayReason) override;
};
