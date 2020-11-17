#pragma once

#include <boost/asio.hpp>
#include <boost/lockfree/queue.hpp>
#include <flatbuffers/flatbuffers.h>

class Message;
class MessageHandler;
enum class MessageID;

// Connection 은 클라이언트와의 연결을 표현합니다.
class Connection
{
public:
	// 생성자
	Connection(boost::asio::io_context& ioContext, const int32_t& headerSize);

	// 소멸자
	virtual ~Connection();

	// Socket 은 소켓을 반환합니다.
	boost::asio::ip::tcp::socket& Socket();

	// Start 는 연결을 시작합니다.
	void Start();

	// Tick 은 주기적으로 실행되며 연결을 처리합니다.
	void Tick();

	// Write 는 메시지를 씁니다.
	void Write(std::function<Message*(flatbuffers::FlatBufferBuilder& builder)> newMessage);

private:
	// _ReadBody 는 바디를 읽습니다.
	void _ReadBody(const MessageID& id, const int32_t& size);

	// _ReadHeader 는 헤더를 읽습니다.
	void _ReadHeader();

	// _Write 는 메시지를 씁니다.
	void _Write();

private:
	// 메시지 빌더
	using _MessageBuilder = std::function<Message*(flatbuffers::FlatBufferBuilder& builder)>;

private:
	boost::asio::ip::tcp::socket m_socket;
	boost::asio::io_context& m_ioContext;

	std::vector<uint8_t> m_messageInHeaderBuf;
	Message* m_messageInTemp;
	boost::lockfree::queue<const Message*> m_messageIns;

	flatbuffers::FlatBufferBuilder m_messageOutHeaderBuilder;
	flatbuffers::FlatBufferBuilder m_messageOutBodyBuilder;
	Message* m_messageOutTemp;
	boost::lockfree::queue<_MessageBuilder*> m_messageOutBuilders;
};