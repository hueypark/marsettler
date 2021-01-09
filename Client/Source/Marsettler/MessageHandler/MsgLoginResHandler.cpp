#include "MsgLoginResHandler.h"

#include "Context.h"
#include "Core/Log.h"
#include "Kismet/GameplayStatics.h"
#include "MarsettlerPlayerController.h"
#include "Message/Message.h"
#include "Message/MessageID.h"
#include "Message/MsgLoginRes_generated.h"
#include "MessageHandler/MessageHandlers.h"
#include "World/Actor.h"
#include "World/World.h"

void MsgLoginResHandler::Handle(const Message* message)
{
	const fbs::MsgLoginRes* loginRes = fbs::GetMsgLoginRes(message->Data());
	if (!loginRes)
	{
		LOG_PRINT("loginRes is null.");

		return;
	}

	const fbs::MsgActor* actor = loginRes->Actor();
	if (!actor)
	{
		LOG_PRINT("actor is null.");

		return;
	}

	std::shared_ptr<World> world = std::make_shared<World>();
	std::shared_ptr<Actor> myActor =
		std::make_shared<Actor>(actor->ID(), FVector(actor->Location().X(), actor->Location().Y(), 0.0f));

	world->AddActor(myActor);

	Context::Instance.World = world;
	Context::Instance.MyActor = myActor;
}
