// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_MSGVECTOR_FBS_H_
#define FLATBUFFERS_GENERATED_MSGVECTOR_FBS_H_

#include "flatbuffers/flatbuffers.h"

namespace fbs {

struct MsgVector;

FLATBUFFERS_MANUALLY_ALIGNED_STRUCT(4) MsgVector FLATBUFFERS_FINAL_CLASS {
 private:
  float X_;
  float Y_;

 public:
  MsgVector() {
    memset(static_cast<void *>(this), 0, sizeof(MsgVector));
  }
  MsgVector(float _X, float _Y)
      : X_(flatbuffers::EndianScalar(_X)),
        Y_(flatbuffers::EndianScalar(_Y)) {
  }
  float X() const {
    return flatbuffers::EndianScalar(X_);
  }
  float Y() const {
    return flatbuffers::EndianScalar(Y_);
  }
};
FLATBUFFERS_STRUCT_END(MsgVector, 8);

}  // namespace fbs

#endif  // FLATBUFFERS_GENERATED_MSGVECTOR_FBS_H_