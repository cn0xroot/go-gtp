// Copyright 2019 go-gtp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package messages_test

import (
	"testing"

	v2 "github.com/wmnsk/go-gtp/v2"
	"github.com/wmnsk/go-gtp/v2/ies"
	"github.com/wmnsk/go-gtp/v2/messages"
	"github.com/wmnsk/go-gtp/v2/testutils"
)

func TestCreateSessionResponse(t *testing.T) {
	cases := []testutils.TestCase{
		{
			Description: "Normal/FromSGWtoMME",
			Structured: messages.NewCreateSessionResponse(
				testutils.TestBearerInfo.TEID, testutils.TestBearerInfo.Seq,
				ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
				ies.NewFullyQualifiedTEID(v2.IFTypeS11S4SGWGTPC, 0xffffffff, "1.1.1.3", ""),
				ies.NewFullyQualifiedTEID(v2.IFTypeS5S8PGWGTPC, 0xffffffff, "1.1.1.2", "").WithInstance(1),
				ies.NewPDNAddressAllocation("2.2.2.2"),
				ies.NewAPNRestriction(v2.APNRestrictionPublic1),
				ies.NewBearerContext(
					ies.NewCause(v2.CauseRequestAccepted, 0, 0, 0, nil),
					ies.NewEPSBearerID(0x05),
					ies.NewFullyQualifiedTEID(v2.IFTypeS1USGWGTPU, 0xffffffff, "1.1.1.3", ""),
					ies.NewFullyQualifiedTEID(v2.IFTypeS5S8PGWGTPU, 0xffffffff, "1.1.1.2", "").WithInstance(1),
				),
				ies.NewFullyQualifiedCSID("1.1.1.2", 1),
				ies.NewFullyQualifiedCSID("1.1.1.3", 1).WithInstance(1),
				ies.NewChargingID(1),
			),
			Marshald: []byte{
				// Header
				0x48, 0x21, 0x00, 0x7d, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x01, 0x00,
				// Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				// F-TEID
				0x57, 0x00, 0x09, 0x00, 0x8b, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x03,
				// F-TEID
				0x57, 0x00, 0x09, 0x01, 0x87, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x02,
				// PAA
				0x4f, 0x00, 0x05, 0x00, 0x01, 0x02, 0x02, 0x02, 0x02,
				// APNRestriction
				0x7f, 0x00, 0x01, 0x00, 0x01,
				// BearerContext
				0x5d, 0x00, 0x25, 0x00,
				//   Cause
				0x02, 0x00, 0x02, 0x00, 0x10, 0x00,
				//   EBI
				0x49, 0x00, 0x01, 0x00, 0x05,
				//   F-TEID
				0x57, 0x00, 0x09, 0x00, 0x81, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x03,
				//   F-TEID
				0x57, 0x00, 0x09, 0x01, 0x85, 0xff, 0xff, 0xff, 0xff, 0x01, 0x01, 0x01, 0x02,
				// FQ-CSID
				0x84, 0x00, 0x07, 0x00, 0x01, 0x01, 0x01, 0x01, 0x02, 0x00, 0x01,
				0x84, 0x00, 0x07, 0x01, 0x01, 0x01, 0x01, 0x01, 0x03, 0x00, 0x01,
				// ChargingID
				0x5e, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x01,
			},
		},
	}

	testutils.Run(t, cases, func(b []byte) (testutils.Marshalable, error) {
		v, err := messages.ParseCreateSessionResponse(b)
		if err != nil {
			return nil, err
		}
		v.Payload = nil
		return v, nil
	})
}
