#include "LoginResponseHandler.h"

#include "Context.h"
#include "Log/Log.h"
#include "Message/Message.h"
#include "Message/MessageID.h"
#include "Message/MsgLoginRes_generated.h"
#include "MessageHandler/MessageHandlers.h"
#include "World/Actor.h"
#include "World/World.h"

void LoginResponseHandler::Handle(const Message* message)
{
	const fbs::MsgLoginRes* loginRes = fbs::GetMsgLoginRes(message->Data());
	if (!loginRes)
	{
		return;
	}

	std::shared_ptr<World> world = std::make_shared<World>();
	std::shared_ptr<Actor> myActor = std::make_shared<Actor>(loginRes->Actor()->ID());

	world->AddActor(myActor);

	Context::Instance.World = world;
	Context::Instance.MyActor = myActor;
}
