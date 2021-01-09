#pragma once

#include <stdint.h>

// MessageID 는 고유한 메시지의 ID 입니다.
enum class MessageID : int32_t
{
	LoginReq = 0,
	LoginRes = 1,
	MovePush = 2,
	MoveReq = 3,
	MoveRes = 4,
};
