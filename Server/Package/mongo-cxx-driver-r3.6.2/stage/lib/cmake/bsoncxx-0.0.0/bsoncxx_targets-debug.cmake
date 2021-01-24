#----------------------------------------------------------------
# Generated CMake target import file for configuration "Debug".
#----------------------------------------------------------------

# Commands may need to know the format version.
set(CMAKE_IMPORT_FILE_VERSION 1)

# Import target "mongo::bsoncxx_shared" for configuration "Debug"
set_property(TARGET mongo::bsoncxx_shared APPEND PROPERTY IMPORTED_CONFIGURATIONS DEBUG)
set_target_properties(mongo::bsoncxx_shared PROPERTIES
  IMPORTED_IMPLIB_DEBUG "${_IMPORT_PREFIX}/lib/bsoncxx.lib"
  IMPORTED_LINK_DEPENDENT_LIBRARIES_DEBUG "mongo::bson_shared"
  IMPORTED_LOCATION_DEBUG "${_IMPORT_PREFIX}/bin/bsoncxx.dll"
  )

list(APPEND _IMPORT_CHECK_TARGETS mongo::bsoncxx_shared )
list(APPEND _IMPORT_CHECK_FILES_FOR_mongo::bsoncxx_shared "${_IMPORT_PREFIX}/lib/bsoncxx.lib" "${_IMPORT_PREFIX}/bin/bsoncxx.dll" )

# Commands beyond this point should not need to know the version.
set(CMAKE_IMPORT_FILE_VERSION)
