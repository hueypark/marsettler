#include "MsgMoveReqHandler.h"

#include "Message/Message.h"
#include "Message/MsgMovePushBuilder.h"
#include "Message/MsgMovePush_generated.h"
#include "Message/MsgMoveReq_generated.h"
#include "Network/Connection.h"

void MsgMoveReqHandler::Handle(Connection* conn, const Message* message)
{
	const fbs::MsgMoveReq* moveReq = fbs::GetMsgMoveReq(message->Data());
	if (!moveReq)
	{
		// TODO(jeawan): 로그 기록

		return;
	}

	MsgMovePushBuilder movePush(moveReq->ID(), *moveReq->Pos());
	conn->Write(movePush);
}
