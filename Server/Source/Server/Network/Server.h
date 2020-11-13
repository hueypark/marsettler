#include "Connection.h"

#include <boost/asio.hpp>

// Server �� ������ ǥ���մϴ�.
class Server
{
public:
	// ������
	Server(boost::asio::io_context& io_context);

private:
	// ������ �޾Ƶ��̱� �����Ѵ�.
	void _StartAccept();

private:
	boost::asio::io_context& m_ioContext;
	boost::asio::ip::tcp::acceptor m_acceptor;
};
