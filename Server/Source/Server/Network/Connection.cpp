#include "Connection.h"

#include <Message/Header_generated.h>
#include <Message/Message.h>

#include <boost/asio.hpp>
#include <iostream>

Connection::Connection(boost::asio::io_context& io_context, const int32_t& headerSize)
	: m_socket(io_context), m_messageTemp(nullptr), m_messages(0)
{
	m_headerBuf.resize(headerSize);

	_ReadHeader();
}

boost::asio::ip::tcp::socket& Connection::Socket()
{
	return m_socket;
}

void Connection::Start()
{
	_ReadHeader();
}

void Connection::Tick()
{
	m_messages.consume_all([](const Message* message) {
		// TODO(jaewan): 메시지 처리

		delete message;
	});
}

void Connection::_ReadBody(const int32_t& id, const int32_t& size)
{
	m_messageTemp = new Message(id, size);

	boost::asio::async_read(m_socket, boost::asio::buffer(m_messageTemp->Data(), m_messageTemp->Size()),
		[this](std::error_code ec, std::size_t length) {
			if (!ec)
			{
				std::cout << ec.message() << std::endl;

				m_socket.close();

				return;
			}

			while (!m_messages.push(m_messageTemp))
			{
			}

			m_messageTemp = nullptr;
		});
}

void Connection::_ReadHeader()
{
	boost::asio::async_read(
		m_socket, boost::asio::buffer(m_headerBuf.data(), m_headerBuf.size()), [this](std::error_code ec, std::size_t length) {
			if (!ec)
			{
				std::cout << ec.message() << std::endl;

				m_socket.close();

				return;
			}

			const fbs::Header* header = fbs::GetHeader(m_headerBuf.data());
			if (!header)
			{
				std::cout << "Header is null." << std::endl;

				m_socket.close();

				return;
			}
		});
}
