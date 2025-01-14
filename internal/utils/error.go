package utils

import (
	"github.com/TXOne-Stellar/stellar-lib/errutil/richerr"
	"github.com/TXOne-Stellar/stellar-pb/errcode/v1"
	"golang.org/x/exp/constraints"
)

const (
	MyServiceCode = 10
)

var (
	// general errors related to flow, state, etc.
	ErrInternal = BuildErr(errcode.Category_Internal, 1, "internal error")
	ErrState    = BuildErr(errcode.Category_Internal, 2, "state error")
	ErrCanceled = BuildErr(errcode.Category_Internal, 3, "canceled")

	// data errors related to formatting, conversion, etc.
	ErrArgMissing = BuildErr(errcode.Category_Arguments, 11, "argument missing")
	ErrArgInvalid = BuildErr(errcode.Category_Arguments, 12, "argument invalid")
	ErrMarshaling = BuildErr(errcode.Category_Internal, 13, "marshaling error")
	ErrCrypto     = BuildErr(errcode.Category_Internal, 14, "crypto error")

	// file related errors
	ErrFileIO      = BuildErr(errcode.Category_Files, 21, "file io error")
	ErrFileContent = BuildErr(errcode.Category_Internal, 22, "file content error")

	// network related errors
	ErrNetwork = BuildErr(errcode.Category_Network, 31, "network error")
	ErrTls     = BuildErr(errcode.Category_RPC, 32, "tls error")
	ErrHttp    = BuildErr(errcode.Category_RPC, 33, "http error")
	ErrGrpc    = BuildErr(errcode.Category_RPC, 34, "grpc error")

	// cache related errors
	ErrRedis = BuildErr(errcode.Category_Cache, 41, "redis error")
)

func BuildErr[
	Category, Code constraints.Integer,
](
	category Category,
	code Code,
	desc string, args ...any,
) func() *richerr.Error {
	return func() *richerr.Error {
		return richerr.New(
			richerr.BuildCode(MyServiceCode, category, code),
		).SetMessage(desc, args...)
	}
}
