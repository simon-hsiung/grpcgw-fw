package test

import (
	"main/protocol/pb/v1"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	"github.com/stretchr/testify/require"
)

func TestSampleRequest(t *testing.T) {
	cases := []struct {
		name  string
		msg   *pb.SampleRequest
		valid bool
	}{
		{
			name: "use guid",
			msg: &pb.SampleRequest{
				Guid: "12345678-1234-1234-1234-123456789012",
			},
			valid: true,
		},
		{
			name: "use id",
			msg: &pb.SampleRequest{
				Id: 1,
			},
			valid: true,
		},
		{
			name: "cannot use both",
			msg: &pb.SampleRequest{
				Guid: "12345678-1234-1234-1234-123456789012",
				Id:   1,
			},
			valid: false,
		},
		{
			name: "invalid guid",
			msg: &pb.SampleRequest{
				Guid: "12345678-1234-1234-1234-12345678901",
			},
			valid: false,
		},
		{
			name: "invalid id",
			msg: &pb.SampleRequest{
				Id: -1,
			},
			valid: false,
		},
	}

	v, err := protovalidate.New()
	require.NoError(t, err)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := v.Validate(c.msg)
			if c.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
