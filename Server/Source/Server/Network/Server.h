#include <boost/asio.hpp>
#include <unordered_map>

class Connection;

// Server �� ������ ǥ���մϴ�.
class Server
{
public:
	// ������
	Server();

	// Start �� ������ �����ŵ�ϴ�.
	void Start();

	// Stop �� ������ �ߴܽ�ŵ�ϴ�.
	void Stop();

private:
	// _StartAccept �� ������ �޾Ƶ��̱� �����մϴ�.
	void _StartAccept();

	// _Tick �� �ֱ������� ����Ǹ� ������ ó���մϴ�.
	void _Tick();

private:
	// _State �� ������ ���¸� ��Ÿ���ϴ�.
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
