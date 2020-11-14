#include "Server.h"

#include "Connection.h"

#include <Message/header_generated.h>

#include <iostream>

Server::Server()
	: m_acceptor(m_ioContext, boost::asio::ip::tcp::endpoint(boost::asio::ip::tcp::v4(), 8080)), m_state(_State::Running)
{
	flatbuffers::FlatBufferBuilder builder(1024);

	auto header = fbs::CreateHeader(builder, 0, 0);
	builder.Finish(header);

	m_headerSize = builder.GetSize();
}

void Server::Start()
{
	_StartAccept();

	std::thread serverThread([this]() {
		while (true)
		{
			_Tick();
		}
	});

	serverThread.detach();

	try
	{
		m_ioContext.run();
	}
	catch (std::exception& e)
	{
		std::cerr << e.what() << std::endl;
	}
}

void Server::Stop()
{
	m_state = _State::StopRequested;

	while (m_state != _State::Stopped)
	{
	}
}

void Server::_StartAccept()
{
	std::shared_ptr<Connection> connection = std::make_shared<Connection>(m_ioContext, m_headerSize);

	m_acceptor.async_accept(connection->Socket(), [this, connection](const boost::system::error_code& error) {
		if (!error)
		{
			connection->Start();
		}

		static int64_t connectionID = 0;
		++connectionID;

		m_connections.emplace(connectionID, connection);

		_StartAccept();
	});
}

void Server::_Tick()
{
	if (m_state == _State::StopRequested)
	{
		m_state = _State::Stopped;
		return;
	}

	for (auto& iter : m_connections)
	{
		std::shared_ptr<Connection>& connection = iter.second;

		connection->Tick();
	}
}
