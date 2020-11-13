#include "Server.h"

Server::Server(boost::asio::io_context& io_context)
	: m_ioContext(io_context), m_acceptor(io_context, boost::asio::ip::tcp::endpoint(boost::asio::ip::tcp::v4(), 8080))
{
	_StartAccept();
}

void Server::_StartAccept()
{
	std::shared_ptr<Connection> connection = std::make_shared<Connection>(m_ioContext);

	m_acceptor.async_accept(connection->Socket(), [this, connection](const boost::system::error_code& error) {
		if (!error)
		{
			connection->Start();
		}

		_StartAccept();
	});
}
