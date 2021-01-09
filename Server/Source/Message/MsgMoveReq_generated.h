// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_MSGMOVEREQ_FBS_H_
#define FLATBUFFERS_GENERATED_MSGMOVEREQ_FBS_H_

#include "flatbuffers/flatbuffers.h"

#include "MsgVector_generated.h"

namespace fbs {

struct MsgMoveReq;
struct MsgMoveReqBuilder;

struct MsgMoveReq FLATBUFFERS_FINAL_CLASS : private flatbuffers::Table {
  typedef MsgMoveReqBuilder Builder;
  enum FlatBuffersVTableOffset FLATBUFFERS_VTABLE_UNDERLYING_TYPE {
    VT_ID = 4,
    VT_POS = 6
  };
  int64_t ID() const {
    return GetField<int64_t>(VT_ID, 0);
  }
  const fbs::MsgVector *Pos() const {
    return GetStruct<const fbs::MsgVector *>(VT_POS);
  }
  bool Verify(flatbuffers::Verifier &verifier) const {
    return VerifyTableStart(verifier) &&
           VerifyField<int64_t>(verifier, VT_ID) &&
           VerifyField<fbs::MsgVector>(verifier, VT_POS) &&
           verifier.EndTable();
  }
};

struct MsgMoveReqBuilder {
  typedef MsgMoveReq Table;
  flatbuffers::FlatBufferBuilder &fbb_;
  flatbuffers::uoffset_t start_;
  void add_ID(int64_t ID) {
    fbb_.AddElement<int64_t>(MsgMoveReq::VT_ID, ID, 0);
  }
  void add_Pos(const fbs::MsgVector *Pos) {
    fbb_.AddStruct(MsgMoveReq::VT_POS, Pos);
  }
  explicit MsgMoveReqBuilder(flatbuffers::FlatBufferBuilder &_fbb)
        : fbb_(_fbb) {
    start_ = fbb_.StartTable();
  }
  MsgMoveReqBuilder &operator=(const MsgMoveReqBuilder &);
  flatbuffers::Offset<MsgMoveReq> Finish() {
    const auto end = fbb_.EndTable(start_);
    auto o = flatbuffers::Offset<MsgMoveReq>(end);
    return o;
  }
};

inline flatbuffers::Offset<MsgMoveReq> CreateMsgMoveReq(
    flatbuffers::FlatBufferBuilder &_fbb,
    int64_t ID = 0,
    const fbs::MsgVector *Pos = 0) {
  MsgMoveReqBuilder builder_(_fbb);
  builder_.add_ID(ID);
  builder_.add_Pos(Pos);
  return builder_.Finish();
}

inline const fbs::MsgMoveReq *GetMsgMoveReq(const void *buf) {
  return flatbuffers::GetRoot<fbs::MsgMoveReq>(buf);
}

inline const fbs::MsgMoveReq *GetSizePrefixedMsgMoveReq(const void *buf) {
  return flatbuffers::GetSizePrefixedRoot<fbs::MsgMoveReq>(buf);
}

inline bool VerifyMsgMoveReqBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifyBuffer<fbs::MsgMoveReq>(nullptr);
}

inline bool VerifySizePrefixedMsgMoveReqBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifySizePrefixedBuffer<fbs::MsgMoveReq>(nullptr);
}

inline void FinishMsgMoveReqBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<fbs::MsgMoveReq> root) {
  fbb.Finish(root);
}

inline void FinishSizePrefixedMsgMoveReqBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<fbs::MsgMoveReq> root) {
  fbb.FinishSizePrefixed(root);
}

}  // namespace fbs

#endif  // FLATBUFFERS_GENERATED_MSGMOVEREQ_FBS_H_
