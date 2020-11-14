#include <boost/asio.hpp>
#include <unordered_map>

class Connection;

// Server 는 서버를 표현합니다.
class Server
{
public:
	// 생성자
	Server();

	// Start 는 서버를 실행시킵니다.
	void Start();

	// Stop 은 서버를 중단시킵니다.
	void Stop();

private:
	// _StartAccept 는 연결을 받아들이기 시작합니다.
	void _StartAccept();

	// _Tick 은 주기적으로 실행되며 서버를 처리합니다.
	void _Tick();

private:
	// _State 는 서버의 상태를 나타냅니다.
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
