// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_MSGACTORPUSH_FBS_H_
#define FLATBUFFERS_GENERATED_MSGACTORPUSH_FBS_H_

#include "flatbuffers/flatbuffers.h"

#include "MsgActor_generated.h"
#include "MsgVector_generated.h"

namespace fbs {

struct MsgActorPush;
struct MsgActorPushBuilder;

struct MsgActorPush FLATBUFFERS_FINAL_CLASS : private flatbuffers::Table {
  typedef MsgActorPushBuilder Builder;
  enum FlatBuffersVTableOffset FLATBUFFERS_VTABLE_UNDERLYING_TYPE {
    VT_ACTOR = 4
  };
  const fbs::MsgActor *Actor() const {
    return GetStruct<const fbs::MsgActor *>(VT_ACTOR);
  }
  bool Verify(flatbuffers::Verifier &verifier) const {
    return VerifyTableStart(verifier) &&
           VerifyField<fbs::MsgActor>(verifier, VT_ACTOR) &&
           verifier.EndTable();
  }
};

struct MsgActorPushBuilder {
  typedef MsgActorPush Table;
  flatbuffers::FlatBufferBuilder &fbb_;
  flatbuffers::uoffset_t start_;
  void add_Actor(const fbs::MsgActor *Actor) {
    fbb_.AddStruct(MsgActorPush::VT_ACTOR, Actor);
  }
  explicit MsgActorPushBuilder(flatbuffers::FlatBufferBuilder &_fbb)
        : fbb_(_fbb) {
    start_ = fbb_.StartTable();
  }
  MsgActorPushBuilder &operator=(const MsgActorPushBuilder &);
  flatbuffers::Offset<MsgActorPush> Finish() {
    const auto end = fbb_.EndTable(start_);
    auto o = flatbuffers::Offset<MsgActorPush>(end);
    return o;
  }
};

inline flatbuffers::Offset<MsgActorPush> CreateMsgActorPush(
    flatbuffers::FlatBufferBuilder &_fbb,
    const fbs::MsgActor *Actor = 0) {
  MsgActorPushBuilder builder_(_fbb);
  builder_.add_Actor(Actor);
  return builder_.Finish();
}

inline const fbs::MsgActorPush *GetMsgActorPush(const void *buf) {
  return flatbuffers::GetRoot<fbs::MsgActorPush>(buf);
}

inline const fbs::MsgActorPush *GetSizePrefixedMsgActorPush(const void *buf) {
  return flatbuffers::GetSizePrefixedRoot<fbs::MsgActorPush>(buf);
}

inline bool VerifyMsgActorPushBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifyBuffer<fbs::MsgActorPush>(nullptr);
}

inline bool VerifySizePrefixedMsgActorPushBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifySizePrefixedBuffer<fbs::MsgActorPush>(nullptr);
}

inline void FinishMsgActorPushBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<fbs::MsgActorPush> root) {
  fbb.Finish(root);
}

inline void FinishSizePrefixedMsgActorPushBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<fbs::MsgActorPush> root) {
  fbb.FinishSizePrefixed(root);
}

}  // namespace fbs

#endif  // FLATBUFFERS_GENERATED_MSGACTORPUSH_FBS_H_
