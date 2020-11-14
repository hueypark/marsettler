#include <boost/asio.hpp>
#include <boost/lockfree/queue.hpp>

class Message;

// Connection �� Ŭ���̾�Ʈ���� ������ ǥ���մϴ�.
class Connection
{
public:
	Connection(boost::asio::io_context& io_context, const int32_t& headerSize);

	// Socket �� ������ ��ȯ�մϴ�.
	boost::asio::ip::tcp::socket& Socket();

	// Start �� ������ �����մϴ�.
	void Start();

	// Tick �� �ֱ������� ����Ǹ� ������ ó���մϴ�.
	void Tick();

private:
	// _ReadBody �� �ٵ� �н��ϴ�.
	void _ReadBody(const int32_t& id, const int32_t& size);

	// _ReadHeader �� ����� �н��ϴ�.
	void _ReadHeader();

private:
	boost::asio::ip::tcp::socket m_socket;

	std::vector<uint8_t> m_headerBuf;

	Message* m_messageTemp;
	boost::lockfree::queue<const Message*> m_messages;
};