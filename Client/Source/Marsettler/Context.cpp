#include "Context.h"

#include "Connection.h"

Context Context::Instance;

Context::Context() : Connection(nullptr), MyActor(nullptr), World(nullptr)
{
}

Context::~Context()
{
}
