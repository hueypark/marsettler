#pragma once

#include <boost/asio.hpp>
#include <unordered_map>

class Connection;

// Server 는 서버를 표현합니다.
class Server
{
public:
	// 생성자
	Server();

	// Start 는 서버를 시작합니다.
	void Start();

	// Stop 은 서버를 중단합니다.
	void Stop();

private:
	// _StartAccept 연결을 받기 시작합니다.
	void _StartAccept();

	// _Tick 주기적으로 서버를 처리합니다.
	void _Tick();

private:
	// _State 는 서버의 상태를 표현합니다.
	enum class _State
	{
		Running,
		StopRequested,
		Stopped,
	};

private:
	boost::asio::io_context m_ioContext;
	boost::asio::ip::tcp::acceptor m_acceptor;

	std::atomic<_State> m_state;

	int32_t m_headerSize;

	std::unordered_map<int64_t, std::shared_ptr<Connection> > m_connections;
};
