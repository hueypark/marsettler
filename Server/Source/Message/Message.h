#include <cinttypes>
#include <vector>

// Message �� �޽����� ǥ���մϴ�.
class Message
{
public:
	// ������
	Message(const int32_t& id, const int32_t& size);

	// Data �� ������ �����͸� ��ȯ�մϴ�.
	uint8_t* Data();

	// Size �� �������� ũ�⸦ ��ȯ�մϴ�.
	const int32_t Size();

private:
	int32_t m_id;
	std::vector<uint8_t> m_body;
};