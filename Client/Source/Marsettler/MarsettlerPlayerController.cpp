#include "MarsettlerPlayerController.h"

#include "Blueprint/AIBlueprintHelperLibrary.h"
#include "Connection.h"
#include "Context.h"
#include "Core/Log.h"
#include "Engine/World.h"
#include "MarsettlerCharacter.h"
#include "Message/MsgMoveReqBuilder.h"
#include "Runtime/Engine/Classes/Components/DecalComponent.h"
#include "World/Actor.h"

AMarsettlerPlayerController::AMarsettlerPlayerController()
{
	bShowMouseCursor = true;
	DefaultMouseCursor = EMouseCursor::Crosshairs;
}

void AMarsettlerPlayerController::BeginPlay()
{
	Super::BeginPlay();
}

void AMarsettlerPlayerController::EndPlay(const EEndPlayReason::Type EndPlayReason)
{
}

void AMarsettlerPlayerController::PlayerTick(float DeltaTime)
{
	Super::PlayerTick(DeltaTime);

	if (Context::Instance.MyActor && Context::Instance.MyActor->GetLocationUpdated())
	{
		Context::Instance.MyActor->SetLocationUpdated(false);

		UAIBlueprintHelperLibrary::SimpleMoveToLocation(this, Context::Instance.MyActor->Location());
	}
}

void AMarsettlerPlayerController::SetupInputComponent()
{
	// set up gameplay key bindings
	Super::SetupInputComponent();

	InputComponent->BindAction("SetDestination", IE_Released, this, &AMarsettlerPlayerController::MoveToMouseCursor);
	InputComponent->BindTouch(EInputEvent::IE_Released, this, &AMarsettlerPlayerController::MoveToTouchLocation);
}

void AMarsettlerPlayerController::MoveToMouseCursor()
{
	// Trace to see what is under the mouse cursor
	FHitResult Hit;
	GetHitResultUnderCursor(ECC_Visibility, false, Hit);

	if (Hit.bBlockingHit)
	{
		// We hit something, move there
		SetNewMoveDestination(Hit.ImpactPoint);
	}
}

void AMarsettlerPlayerController::MoveToTouchLocation(const ETouchIndex::Type FingerIndex, const FVector Location)
{
	FVector2D ScreenSpaceLocation(Location);

	// Trace to see what is under the touch location
	FHitResult HitResult;
	GetHitResultAtScreenPosition(ScreenSpaceLocation, CurrentClickTraceChannel, true, HitResult);
	if (HitResult.bBlockingHit)
	{
		// We hit something, move there
		SetNewMoveDestination(HitResult.ImpactPoint);
	}
}

void AMarsettlerPlayerController::SetNewMoveDestination(const FVector destLocation)
{
	APawn* const MyPawn = GetPawn();
	if (!MyPawn)
	{
		LOG_PRINT("Pawn is null");

		return;
	}

	float const distance = FVector::Dist(destLocation, MyPawn->GetActorLocation());

	// We need to issue move command only if far enough in order for walk animation to play correctly
	if (distance > 120.0f)
	{
		_SendMoveMessage(destLocation);
	}
}

void AMarsettlerPlayerController::_SendMoveMessage(const FVector destLocation)
{
	if (Context::Instance.MyActor)
	{
		MsgMoveReqBuilder msgMoveReq(Context::Instance.MyActor->ID(), fbs::MsgVector(destLocation.X, destLocation.Y));

		Context::Instance.Connection->WriteMessage(msgMoveReq);
	}
}
