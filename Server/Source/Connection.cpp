#include "Connection.h"

inline Connection::Connection(boost::asio::io_context& io_context) : m_socket(io_context)
{
}

inline boost::asio::ip::tcp::socket& Connection::Socket()
{
	return m_socket;
}

inline void Connection::Start()
{
}
