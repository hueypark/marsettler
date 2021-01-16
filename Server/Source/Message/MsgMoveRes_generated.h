// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_MSGMOVERES_FBS_H_
#define FLATBUFFERS_GENERATED_MSGMOVERES_FBS_H_

#include "flatbuffers/flatbuffers.h"

namespace fbs {

struct MsgMoveRes;
struct MsgMoveResBuilder;

struct MsgMoveRes FLATBUFFERS_FINAL_CLASS : private flatbuffers::Table {
  typedef MsgMoveResBuilder Builder;
  bool Verify(flatbuffers::Verifier &verifier) const {
    return VerifyTableStart(verifier) &&
           verifier.EndTable();
  }
};

struct MsgMoveResBuilder {
  typedef MsgMoveRes Table;
  flatbuffers::FlatBufferBuilder &fbb_;
  flatbuffers::uoffset_t start_;
  explicit MsgMoveResBuilder(flatbuffers::FlatBufferBuilder &_fbb)
        : fbb_(_fbb) {
    start_ = fbb_.StartTable();
  }
  MsgMoveResBuilder &operator=(const MsgMoveResBuilder &);
  flatbuffers::Offset<MsgMoveRes> Finish() {
    const auto end = fbb_.EndTable(start_);
    auto o = flatbuffers::Offset<MsgMoveRes>(end);
    return o;
  }
};

inline flatbuffers::Offset<MsgMoveRes> CreateMsgMoveRes(
    flatbuffers::FlatBufferBuilder &_fbb) {
  MsgMoveResBuilder builder_(_fbb);
  return builder_.Finish();
}

inline const fbs::MsgMoveRes *GetMsgMoveRes(const void *buf) {
  return flatbuffers::GetRoot<fbs::MsgMoveRes>(buf);
}

inline const fbs::MsgMoveRes *GetSizePrefixedMsgMoveRes(const void *buf) {
  return flatbuffers::GetSizePrefixedRoot<fbs::MsgMoveRes>(buf);
}

inline bool VerifyMsgMoveResBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifyBuffer<fbs::MsgMoveRes>(nullptr);
}

inline bool VerifySizePrefixedMsgMoveResBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifySizePrefixedBuffer<fbs::MsgMoveRes>(nullptr);
}

inline void FinishMsgMoveResBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<fbs::MsgMoveRes> root) {
  fbb.Finish(root);
}

inline void FinishSizePrefixedMsgMoveResBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<fbs::MsgMoveRes> root) {
  fbb.FinishSizePrefixed(root);
}

}  // namespace fbs

#endif  // FLATBUFFERS_GENERATED_MSGMOVERES_FBS_H_