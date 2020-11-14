// automatically generated by the FlatBuffers compiler, do not modify


#ifndef FLATBUFFERS_GENERATED_LOGINRESPONSE_FBS_H_
#define FLATBUFFERS_GENERATED_LOGINRESPONSE_FBS_H_

#include "flatbuffers/flatbuffers.h"

namespace fbs {

struct LoginResponse;
struct LoginResponseBuilder;

struct LoginResponse FLATBUFFERS_FINAL_CLASS : private flatbuffers::Table {
  typedef LoginResponseBuilder Builder;
  enum FlatBuffersVTableOffset FLATBUFFERS_VTABLE_UNDERLYING_TYPE {
    VT_ID = 4
  };
  int64_t id() const {
    return GetField<int64_t>(VT_ID, 0);
  }
  bool Verify(flatbuffers::Verifier &verifier) const {
    return VerifyTableStart(verifier) &&
           VerifyField<int64_t>(verifier, VT_ID) &&
           verifier.EndTable();
  }
};

struct LoginResponseBuilder {
  typedef LoginResponse Table;
  flatbuffers::FlatBufferBuilder &fbb_;
  flatbuffers::uoffset_t start_;
  void add_id(int64_t id) {
    fbb_.AddElement<int64_t>(LoginResponse::VT_ID, id, 0);
  }
  explicit LoginResponseBuilder(flatbuffers::FlatBufferBuilder &_fbb)
        : fbb_(_fbb) {
    start_ = fbb_.StartTable();
  }
  LoginResponseBuilder &operator=(const LoginResponseBuilder &);
  flatbuffers::Offset<LoginResponse> Finish() {
    const auto end = fbb_.EndTable(start_);
    auto o = flatbuffers::Offset<LoginResponse>(end);
    return o;
  }
};

inline flatbuffers::Offset<LoginResponse> CreateLoginResponse(
    flatbuffers::FlatBufferBuilder &_fbb,
    int64_t id = 0) {
  LoginResponseBuilder builder_(_fbb);
  builder_.add_id(id);
  return builder_.Finish();
}

inline const fbs::LoginResponse *GetLoginResponse(const void *buf) {
  return flatbuffers::GetRoot<fbs::LoginResponse>(buf);
}

inline const fbs::LoginResponse *GetSizePrefixedLoginResponse(const void *buf) {
  return flatbuffers::GetSizePrefixedRoot<fbs::LoginResponse>(buf);
}

inline bool VerifyLoginResponseBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifyBuffer<fbs::LoginResponse>(nullptr);
}

inline bool VerifySizePrefixedLoginResponseBuffer(
    flatbuffers::Verifier &verifier) {
  return verifier.VerifySizePrefixedBuffer<fbs::LoginResponse>(nullptr);
}

inline void FinishLoginResponseBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<fbs::LoginResponse> root) {
  fbb.Finish(root);
}

inline void FinishSizePrefixedLoginResponseBuffer(
    flatbuffers::FlatBufferBuilder &fbb,
    flatbuffers::Offset<fbs::LoginResponse> root) {
  fbb.FinishSizePrefixed(root);
}

}  // namespace fbs

#endif  // FLATBUFFERS_GENERATED_LOGINRESPONSE_FBS_H_
