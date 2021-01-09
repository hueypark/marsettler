// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_MSGACTOR_FBS_H_
#define FLATBUFFERS_GENERATED_MSGACTOR_FBS_H_

#include "flatbuffers/flatbuffers.h"

#include "MsgVector_generated.h"

namespace fbs {

struct MsgActor;

FLATBUFFERS_MANUALLY_ALIGNED_STRUCT(8) MsgActor FLATBUFFERS_FINAL_CLASS {
 private:
  int64_t ID_;
  fbs::MsgVector Location_;

 public:
  MsgActor() {
    memset(static_cast<void *>(this), 0, sizeof(MsgActor));
  }
  MsgActor(int64_t _ID, const fbs::MsgVector &_Location)
      : ID_(flatbuffers::EndianScalar(_ID)),
        Location_(_Location) {
  }
  int64_t ID() const {
    return flatbuffers::EndianScalar(ID_);
  }
  const fbs::MsgVector &Location() const {
    return Location_;
  }
};
FLATBUFFERS_STRUCT_END(MsgActor, 16);

}  // namespace fbs

#endif  // FLATBUFFERS_GENERATED_MSGACTOR_FBS_H_
