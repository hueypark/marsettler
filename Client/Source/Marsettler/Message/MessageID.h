#pragma once

#include <stdint.h>

// MessageID 는 고유한 메시지의 ID 입니다.
enum class MessageID : int32_t
{
	Login = 0,
	LoginResponse = 1,
};