// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_MOVEPUSH_FBS_H_
#define FLATBUFFERS_GENERATED_MOVEPUSH_FBS_H_

#include "flatbuffers/flatbuffers.h"

#include "Vector_generated.h"

namespace fbs {

struct MovePush;
struct MovePushBuilder;

struct MovePush FLATBUFFERS_FINAL_CLASS : private flatbuffers::Table {
  typedef MovePushBuilder Builder;
  bool Verify(flatbuffers::Verifier &verifier) const {
    return VerifyTableStart(verifier) &&
           verifier.EndTable();
  }
};

struct MovePushBuilder {
  typedef MovePush Table;
  flatbuffers::FlatBufferBuilder &fbb_;
  flatbuffers::uoffset_t start_;
  explicit MovePushBuilder(flatbuffers::FlatBufferBuilder &_fbb)
        : fbb_(_fbb) {
    start_ = fbb_.StartTable();
  }
  MovePushBuilder &operator=(const MovePushBuilder &);
  flatbuffers::Offset<MovePush> Finish() {
    const auto end = fbb_.EndTable(start_);
    auto o = flatbuffers::Offset<MovePush>(end);
    return o;
  }
};

inline flatbuffers::Offset<MovePush> CreateMovePush(
    flatbuffers::FlatBufferBuilder &_fbb) {
  MovePushBuilder builder_(_fbb);
  return builder_.Finish();
}

inline const fbs::MovePush *GetMovePush(const void *buf) {
  return flatbuffers::GetRoot<fbs::MovePush>(buf);
}

inline const fbs::MovePush *GetSizePrefixedMovePush(const void *buf) {
  return flatbuffers::GetSizePrefixedRoot<fbs::MovePush>(buf);
}

inline bool VerifyMovePushBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifyBuffer<fbs::MovePush>(nullptr);
}

inline bool VerifySizePrefixedMovePushBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifySizePrefixedBuffer<fbs::MovePush>(nullptr);
}

inline void FinishMovePushBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<fbs::MovePush> root) {
  fbb.Finish(root);
}

inline void FinishSizePrefixedMovePushBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<fbs::MovePush> root) {
  fbb.FinishSizePrefixed(root);
}

}  // namespace fbs

#endif  // FLATBUFFERS_GENERATED_MOVEPUSH_FBS_H_