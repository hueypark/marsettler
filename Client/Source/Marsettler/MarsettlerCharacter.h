﻿// Copyright Epic Games, Inc. All Rights Reserved.

#pragma once

#include "CoreMinimal.h"
#include "GameFramework/Character.h"

#include "MarsettlerCharacter.generated.h"

class NetworkComponent;

UCLASS(Blueprintable)
class AMarsettlerCharacter : public ACharacter
{
	GENERATED_BODY()

public:
	AMarsettlerCharacter();

	// 게임 시작시 호출됩니다.
	virtual void BeginPlay() override;

	// 게임 종료시 호출됩니다.
	virtual void EndPlay(const EEndPlayReason::Type endPlayReason) override;

	// Called every frame.
	virtual void Tick(float DeltaSeconds) override;

	/** Returns TopDownCameraComponent subobject **/
	FORCEINLINE class UCameraComponent* GetTopDownCameraComponent() const
	{
		return TopDownCameraComponent;
	}
	/** Returns CameraBoom subobject **/
	FORCEINLINE class USpringArmComponent* GetCameraBoom() const
	{
		return CameraBoom;
	}
	/** Returns CursorToWorld subobject **/
	FORCEINLINE class UDecalComponent* GetCursorToWorld()
	{
		return CursorToWorld;
	}

private:
	/** Top down camera */
	UPROPERTY(VisibleAnywhere, BlueprintReadOnly, Category = Camera, meta = (AllowPrivateAccess = "true"))
	class UCameraComponent* TopDownCameraComponent;

	/** Camera boom positioning the camera above the character */
	UPROPERTY(VisibleAnywhere, BlueprintReadOnly, Category = Camera, meta = (AllowPrivateAccess = "true"))
	class USpringArmComponent* CameraBoom;

	/** A decal that projects to the cursor location. */
	UPROPERTY(VisibleAnywhere, BlueprintReadOnly, Category = Camera, meta = (AllowPrivateAccess = "true"))
	class UDecalComponent* CursorToWorld;

	// 네트워크 컴포넌트
	NetworkComponent* m_networkComponent;
};
