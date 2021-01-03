#pragma once

#include "Marsettler.h"

#define LOG_PRINT(s, ...) UE_LOG(LogPrint, Log, TEXT(s), ##__VA_ARGS__)

#define LOG_FATAL(s, ...) UE_LOG(LogFatal, Fatal, TEXT(s), ##__VA_ARGS__)
