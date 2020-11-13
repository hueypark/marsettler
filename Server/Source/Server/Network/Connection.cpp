#include "Connection.h"

Connection::Connection(boost::asio::io_context& io_context) : m_socket(io_context)
{
}

boost::asio::ip::tcp::socket& Connection::Socket()
{
	return m_socket;
}

void Connection::Start()
{
}
