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
	, m_messageInTemp(nullptr)
	, m_messageIns(0)
	, m_messageOutFlag(true)
	, m_messageOutHeaderBuilder(headerSize)
{
	m_messageInHeaderBuf.resize(headerSize);
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
			MessageHandler::Handle(this, message);

			delete message;
		});

	if (!m_messageOutFlag.load())
	{
		_Write();
	}
}

void Connection::Write(const MessageBuilder& builder)
{
	m_messageOutMux.lock();
	m_messageOutBuilders.emplace(builder.Clone());
	m_messageOutMux.unlock();
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

	builder->Build(m_messageOutBodyBuilder);

	auto header = fbs::CreateHeader(m_messageOutHeaderBuilder, int32_t(builder->ID()), m_messageOutBodyBuilder.GetSize());
	m_messageOutHeaderBuilder.Finish(header);

	boost::asio::async_write(m_socket,
		boost::array<boost::asio::const_buffer, 2>{
			boost::asio::buffer(m_messageOutHeaderBuilder.GetBufferPointer(), m_messageOutHeaderBuilder.GetSize()),
			boost::asio::buffer(m_messageOutBodyBuilder.GetBufferPointer(), m_messageOutBodyBuilder.GetSize())},
		[this](std::error_code ec, std::size_t length)
		{
			if (!ec)
			{
				std::cerr << ec.message() << std::endl;

				m_socket.close();

				return;
			}

			_Write();
		});
}
