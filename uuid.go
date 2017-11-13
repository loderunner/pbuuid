// pbuuid
// A Go package to wrap UUIDs in byte arrays in protobuf.
// Copyright (c) 2017 Charles Francoise. All Rights Reserved
//
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package pbuuid

import (
	guuid "github.com/google/uuid"
	proto "github.com/loderunner/pbuuid/proto"
)

// UUID converts a `loderunner.pbuuid.UUID` proto to a `uuid.UUID`.
// It returns an error if the argument is invalid.
func UUID(u *proto.UUID) (guuid.UUID, error) {
	var g guuid.UUID
	err := g.UnmarshalBinary(u.GetBytes())
	return g, err
}

// UUIDProto converts the `uuid.UUID` to a `loderunner.pbuuid.UUID` proto.
// It returns an error if the resulting UUID is invalid.
func UUIDProto(g guuid.UUID) (*proto.UUID, error) {
	b, err := g.MarshalBinary()
	u := proto.UUID{
		Bytes: b,
	}
	return &u, err
}
