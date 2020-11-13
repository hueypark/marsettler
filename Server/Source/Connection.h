#include <boost/asio.hpp>

// Connection 은 클라이언트와의 연결을 표현합니다.
class Connection
{
public:
	Connection(boost::asio::io_context& io_context);

	// Socket 은 소켓을 반환합니다.
	boost::asio::ip::tcp::socket& Socket();

	// Start 는 연결을 시작합니다.
	void Start();

private:
	boost::asio::ip::tcp::socket m_socket;
};