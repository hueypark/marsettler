#include "LoginResponseHandler.h"

#include "Context.h"
#include "Log/Log.h"
#include "Message/LoginResponse_generated.h"
#include "Message/Message.h"
#include "Message/MessageID.h"
#include "MessageHandler/MessageHandlers.h"
#include "World/Actor.h"
#include "World/World.h"

void LoginResponseHandler::Handle(const Message* message)
{
	const fbs::LoginResponse* loginResponse = fbs::GetLoginResponse(message->Data());
	if (!loginResponse)
	{
		return;
	}

	if (!Context::Instance.MyActor)
	{
		return;
	}

	std::shared_ptr<World> world = std::make_shared<World>();
	std::shared_ptr<Actor> myActor = std::make_shared<Actor>(loginResponse->ID());

	world->AddActor(myActor);

	Context::Instance.World = world;
	Context::Instance.MyActor = myActor;
}
