#include "Connection.h"

#include <boost/asio.hpp>

// Server 는 서버를 표현합니다.
class Server
{
public:
	// 생성자
	Server(boost::asio::io_context& io_context);

private:
	// 연결을 받아들이기 시작한다.
	void _StartAccept();

private:
	boost::asio::io_context& m_ioContext;
	boost::asio::ip::tcp::acceptor m_acceptor;
};
