package message

import (
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/hueypark/marsettler/server/pkg/message/fbs"
)

// MakeSignInRequest 는 계정생성 메시지를 생성합니다.
func MakeSignInRequest(b *flatbuffers.Builder, id int64) []byte {
	b.Reset()

	fbs.SignInRequestStart(b)
	fbs.SignInRequestAddId(b, id)
	offset := fbs.SignInRequestEnd(b)

	b.Finish(offset)

	return b.Bytes[b.Head():]
}
