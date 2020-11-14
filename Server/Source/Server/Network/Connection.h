#pragma once

#include <boost/asio.hpp>
#include <boost/lockfree/queue.hpp>

class Message;
class MessageHandler;
enum class MessageID;

// Connection 은 클라이언트와의 연결을 표현합니다.
class Connection
{
public:
	// 생성자
	Connection(boost::asio::io_context& io_context, const int32_t& headerSize);

	// 소멸자
	virtual ~Connection();

	// Socket 은 소켓을 반환합니다.
	boost::asio::ip::tcp::socket& Socket();

	// Start 는 연결을 시작합니다.
	void Start();

	// Tick 은 주기적으로 실행되며 연결을 처리합니다.
	void Tick();

private:
	// _ReadBody 는 바디를 읽습니다.
	void _ReadBody(const MessageID& id, const int32_t& size);

	// _ReadHeader 는 헤더를 읽습니다.
	void _ReadHeader();

private:
	boost::asio::ip::tcp::socket m_socket;

	std::vector<uint8_t> m_headerBuf;

	Message* m_messageTemp;
	boost::lockfree::queue<const Message*> m_messages;

	MessageHandler* m_messageHandler;
};