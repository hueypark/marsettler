#include <boost/asio.hpp>

// Connection �� Ŭ���̾�Ʈ���� ������ ǥ���մϴ�.
class Connection
{
public:
	Connection(boost::asio::io_context& io_context);

	// Socket �� ������ ��ȯ�մϴ�.
	boost::asio::ip::tcp::socket& Socket();

	// Start �� ������ �����մϴ�.
	void Start();

private:
	boost::asio::ip::tcp::socket m_socket;
};