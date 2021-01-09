#pragma once

#include "CoreMinimal.h"
#include "GameFramework/PlayerController.h"
#include "Message/MsgActor_generated.h"

#include "MarsettlerPlayerController.generated.h"

UCLASS()
class AMarsettlerPlayerController : public APlayerController
{
	GENERATED_BODY()

public:
	AMarsettlerPlayerController();

protected:
	virtual void BeginPlay() override;

	virtual void EndPlay(const EEndPlayReason::Type EndPlayReason) override;

	// Begin PlayerController interface
	virtual void PlayerTick(float DeltaTime) override;
	virtual void SetupInputComponent() override;
	// End PlayerController interface

	/** Navigate player to the current mouse cursor location. */
	void MoveToMouseCursor();

	/** Navigate player to the current touch location. */
	void MoveToTouchLocation(const ETouchIndex::Type FingerIndex, const FVector Location);

	/** Navigate player to the given world location. */
	void SetNewMoveDestination(const FVector destLocation);

private:
	// 이동 메시지를 전송한다.
	void _SendMoveMessage(const FVector destLocation);
};
