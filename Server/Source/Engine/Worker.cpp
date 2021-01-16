#include "Worker.h"

Worker::Worker() : m_works(0)
{
}

void Worker::AddWork(const Work& work)
{
	Work* newWork = new Work(work);

	while (!m_works.push(newWork))
	{
	}
}

void Worker::Tick()
{
	m_works.consume_all(
		[this](const Work* work)
		{
			(*work)();

			delete work;
		});
}
