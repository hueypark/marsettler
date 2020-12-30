#include "Context.h"

#include "Engine/World/World.h"

Context Context::Instance;

Context::Context()
{
	GameWorld = std::make_shared<World>();
}
