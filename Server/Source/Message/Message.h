#include "MessageID.h"

#include <cinttypes>
#include <vector>

// Message 는 메시지를 표현합니다.
class Message
{
public:
	// 생성자
	Message(const MessageID& id, const int32_t& size);

	// Data 는 데이터 포인터를 반환합니다.
	uint8_t* Data();

	// Size 는 데이터의 크기를 반환합니다.
	const int32_t Size();

private:
	MessageID m_id;
	std::vector<uint8_t> m_body;
};