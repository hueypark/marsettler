#include "MsgMovePushHandler.h"

#include "Context.h"
#include "Core/Log.h"
#include "Message/Message.h"
#include "Message/MsgMovePush_generated.h"
#include "World/Actor.h"

void MsgMovePushHandler::Handle(const Message* message)
{
	const fbs::MsgMovePush* movePush = fbs::GetMsgMovePush(message->Data());
	if (!movePush)
	{
		LOG_PRINT("movePush is null.");

		return;
	}

	// TODO(jaewan): ID로 Actor 찾아서 갱신하게 수정
	Context::Instance.MyActor->SetLocation(FVector(movePush->Pos()->X(), movePush->Pos()->Y(), 0.0f));
}
