#include "Connection.h"

#include "MessageHandler/MessageHandler.h"

#include <Message/Header_generated.h>
#include <Message/Message.h>
#include <Message/MessageID.h>
#include <boost/asio.hpp>

#include <iostream>

Connection::Connection(boost::asio::io_context& ioContext, const int32_t& headerSize)
	: m_socket(ioContext)
	, m_ioContext(ioContext)
	, m_messageOutHeaderBuilder(headerSize)
	, m_messageIns(0)
	, m_messageOutTemp(nullptr)
	, m_messageOutBuilders(0)
{
	m_messageInHeaderBuf.resize(headerSize);

	_ReadHeader();
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

	boost::asio::post(m_ioContext,
		[this]()
		{
			_Write();
		});
}

void Connection::Tick()
{
	m_messageIns.consume_all(
		[this](const Message* message)
		{
			MessageHandler::Handle(this, message);

			delete message;
		});
}

void Connection::Write(_MessageBuilder messageBuilder)
{
	auto func = new _MessageBuilder(messageBuilder);
	while (!m_messageOutBuilders.push(func))
	{
	}
}

void Connection::_ReadBody(const MessageID& id, const int32_t& size)
{
	m_messageInTemp = new Message(id, size);

	boost::asio::async_read(m_socket, boost::asio::buffer(m_messageInTemp->Data(), m_messageInTemp->Size()),
		[this](std::error_code ec, std::size_t length)
		{
			if (!ec)
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
	boost::asio::async_read(m_socket, boost::asio::buffer(m_messageInHeaderBuf.data(), m_messageInHeaderBuf.size()),
		[this](std::error_code ec, std::size_t length)
		{
			if (!ec)
			{
				std::cerr << ec.message() << std::endl;

				m_socket.close();

				return;
			}

			const fbs::Header* header = fbs::GetHeader(m_messageInHeaderBuf.data());
			if (!header)
			{
				std::cout << "Header is null." << std::endl;

				m_socket.close();

				return;
			}

			_ReadBody(MessageID(header->ID()), header->Size());
		});
}

void Connection::_Write()
{
	_MessageBuilder* builder = nullptr;
	while (true)
	{
		if (m_messageOutBuilders.pop(builder))
		{
			break;
		}

		std::this_thread::sleep_for(std::chrono::milliseconds(1));
	}

	m_messageOutTemp = (*builder)(m_messageOutBodyBuilder);

	delete builder;

	auto header = fbs::CreateHeader(m_messageOutHeaderBuilder, int32_t(m_messageOutTemp->ID()), m_messageOutTemp->Size());
	m_messageOutHeaderBuilder.Finish(header);

	boost::asio::async_write(m_socket,
		boost::array<boost::asio::const_buffer, 2>{
			boost::asio::buffer(m_messageOutHeaderBuilder.GetBufferPointer(), m_messageOutHeaderBuilder.GetSize()),
			boost::asio::buffer(m_messageOutTemp->Data(), m_messageOutTemp->Size())},
		[this](std::error_code ec, std::size_t length)
		{
			delete m_messageOutTemp;

			if (!ec)
			{
				std::cerr << ec.message() << std::endl;

				m_socket.close();

				return;
			}

			_Write();
		});
}
