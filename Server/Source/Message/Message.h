#include "MessageID.h"

#include <cinttypes>
#include <vector>

// Message �� �޽����� ǥ���մϴ�.
class Message
{
public:
	// ������
	Message(const MessageID& id, const int32_t& size);

	// Data �� ������ �����͸� ��ȯ�մϴ�.
	uint8_t* Data();

	// Size �� �������� ũ�⸦ ��ȯ�մϴ�.
	const int32_t Size();

private:
	MessageID m_id;
	std::vector<uint8_t> m_body;
};