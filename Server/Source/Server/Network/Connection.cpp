#include "Connection.h"

#include "MessageHandler/MessageHandler.h"

#include <Message/Message.h>
#include <Message/MessageID.h>
#include <boost/asio.hpp>

#include <iostream>

Connection::Connection(boost::asio::io_context& ioContext)
	: m_socket(ioContext)
	, m_ioContext(ioContext)
	, m_messageInHeaderBuf(8)
	, m_messageInTemp(nullptr)
	, m_messageIns(0)
	, m_messageOutFlag(false)
	, m_messageOutHeaderBuf(8)
{
}

Connection::~Connection()
{
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
	m_messageIns.consume_all(
		[this](const Message* message)
		{
			try
			{
				MessageHandler::Handle(this, message);
			}
			catch (const std::exception& e)
			{
				std::cout << e.what() << std::endl;

				m_socket.close();

				return;
			}
			catch (const std::string& e)
			{
				std::cout << e << std::endl;

				m_socket.close();

				return;
			}
			catch (...)
			{
				std::cout << "Something wrong." << std::endl;

				m_socket.close();

				return;
			}

			delete message;
		});

	if (m_messageOutFlag.load())
	{
		_Write();
	}
}

void Connection::Write(const MessageBuilder& builder)
{
	m_messageOutMux.lock();
	m_messageOutBuilders.emplace(builder.Clone());
	m_messageOutFlag.store(true);
	m_messageOutMux.unlock();
}

void Connection::_ReadBody(const MessageID& id, const int32_t& size)
{
	m_messageInTemp = new Message(id, size);

	boost::asio::async_read(
		m_socket, boost::asio::buffer(m_messageInTemp->Data(), m_messageInTemp->Size()),
		[this](std::error_code ec, std::size_t length)
		{
			if (ec)
			{
				std::cout << ec.message() << std::endl;

				m_socket.close();

				return;
			}

			while (!m_messageIns.push(m_messageInTemp))
			{
			}

			_ReadHeader();
		});
}

void Connection::_ReadHeader()
{
	boost::asio::async_read(
		m_socket, boost::asio::buffer(m_messageInHeaderBuf.data(), m_messageInHeaderBuf.size()),
		[this](std::error_code ec, std::size_t length)
		{
			if (ec)
			{
				std::cerr << ec.message() << std::endl;

				m_socket.close();

				return;
			}

			MessageID messageID;
			std::memcpy(&messageID, &m_messageInHeaderBuf[0], 4);

			int32_t messageSize;
			std::memcpy(&messageSize, &m_messageInHeaderBuf[4], 4);

			_ReadBody(messageID, messageSize);
		});
}

void Connection::_Write()
{
	std::unique_ptr<MessageBuilder> builder = nullptr;

	m_messageOutMux.lock();
	if (m_messageOutBuilders.empty())
	{
		m_messageOutFlag.store(false);

		m_messageOutMux.unlock();

		return;
	}
	else
	{
		builder = std::move(m_messageOutBuilders.front());
		m_messageOutBuilders.pop();
	}
	m_messageOutMux.unlock();

	builder->Build(m_messageOutBuilder);

	MessageID messageID = builder->ID();
	std::memcpy(&m_messageOutHeaderBuf[0], &messageID, 4);

	int32_t messageSize = m_messageOutBuilder.GetSize();
	std::memcpy(&m_messageOutHeaderBuf[4], &messageSize, 4);

	boost::asio::async_write(
		m_socket,
		boost::array<boost::asio::const_buffer, 2>{
			boost::asio::buffer(m_messageOutHeaderBuf.data(), m_messageOutHeaderBuf.size()),
			boost::asio::buffer(m_messageOutBuilder.GetBufferPointer(), messageSize)},
		[this](std::error_code ec, std::size_t length)
		{
			if (ec)
			{
				std::cerr << ec.message() << std::endl;

				m_socket.close();

				return;
			}

			_Write();
		});
}
