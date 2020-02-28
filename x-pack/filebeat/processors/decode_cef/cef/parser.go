// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//line cef.rl:1
// Code generated by ragel DO NOT EDIT.
package cef

import (
	"fmt"
	"strconv"

	"go.uber.org/multierr"
)

//line parser.go:15
var _cef_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 16, 16, 0, 0, 0, 0,
	19, 22, 22, 22, 19, 22, 22, 22,
	0,
}

const cef_start int = 1
const cef_first_final int = 31
const cef_error int = 0

const cef_en_gobble_extension int = 28
const cef_en_main int = 1
const cef_en_main_cef_extensions int = 24

//line cef.rl:16

// unpack unpacks a CEF message.
func (e *Event) unpack(data string) error {
	cs, p, pe, eof := 0, 0, len(data), len(data)
	mark := 0

	// Extension key.
	var extKey string

	// Extension value start and end indices.
	extValueStart, extValueEnd := 0, 0

	// recoveredErrs are problems with the message that the parser was able to
	// recover from (though the parsing might not be "correct").
	var recoveredErrs []error

	e.init(data)

//line parser.go:55
	{
		cs = cef_start
	}

//line parser.go:60
	{
		if (p) == (pe) {
			goto _test_eof
		}
		if cs == 0 {
			goto _out
		}
	_resume:
		switch cs {
		case 1:
			if data[(p)] == 67 {
				goto tr0
			}
			goto tr1
		case 0:
			goto _out
		case 2:
			if data[(p)] == 69 {
				goto tr2
			}
			goto tr1
		case 3:
			if data[(p)] == 70 {
				goto tr3
			}
			goto tr1
		case 4:
			if data[(p)] == 58 {
				goto tr4
			}
			goto tr1
		case 5:
			if 48 <= data[(p)] && data[(p)] <= 57 {
				goto tr5
			}
			goto tr1
		case 6:
			if data[(p)] == 124 {
				goto tr7
			}
			if 48 <= data[(p)] && data[(p)] <= 57 {
				goto tr6
			}
			goto tr1
		case 7:
			switch data[(p)] {
			case 92:
				goto tr9
			case 124:
				goto tr10
			}
			goto tr8
		case 8:
			switch data[(p)] {
			case 92:
				goto tr12
			case 124:
				goto tr13
			}
			goto tr11
		case 9:
			switch data[(p)] {
			case 92:
				goto tr11
			case 124:
				goto tr11
			}
			goto tr1
		case 10:
			switch data[(p)] {
			case 92:
				goto tr15
			case 124:
				goto tr16
			}
			goto tr14
		case 11:
			switch data[(p)] {
			case 92:
				goto tr18
			case 124:
				goto tr19
			}
			goto tr17
		case 12:
			switch data[(p)] {
			case 92:
				goto tr17
			case 124:
				goto tr17
			}
			goto tr1
		case 13:
			switch data[(p)] {
			case 92:
				goto tr21
			case 124:
				goto tr22
			}
			goto tr20
		case 14:
			switch data[(p)] {
			case 92:
				goto tr24
			case 124:
				goto tr25
			}
			goto tr23
		case 15:
			switch data[(p)] {
			case 92:
				goto tr23
			case 124:
				goto tr23
			}
			goto tr1
		case 16:
			switch data[(p)] {
			case 92:
				goto tr27
			case 124:
				goto tr28
			}
			goto tr26
		case 17:
			switch data[(p)] {
			case 92:
				goto tr30
			case 124:
				goto tr31
			}
			goto tr29
		case 18:
			switch data[(p)] {
			case 92:
				goto tr29
			case 124:
				goto tr29
			}
			goto tr1
		case 19:
			switch data[(p)] {
			case 92:
				goto tr33
			case 124:
				goto tr34
			}
			goto tr32
		case 20:
			switch data[(p)] {
			case 92:
				goto tr36
			case 124:
				goto tr37
			}
			goto tr35
		case 21:
			switch data[(p)] {
			case 92:
				goto tr35
			case 124:
				goto tr35
			}
			goto tr1
		case 22:
			switch data[(p)] {
			case 45:
				goto tr38
			case 124:
				goto tr39
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr38
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr38
				}
			default:
				goto tr38
			}
			goto tr1
		case 23:
			switch data[(p)] {
			case 45:
				goto tr40
			case 124:
				goto tr41
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr40
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr40
				}
			default:
				goto tr40
			}
			goto tr1
		case 31:
			switch data[(p)] {
			case 32:
				goto tr42
			case 95:
				goto tr43
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr43
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr43
				}
			default:
				goto tr43
			}
			goto tr1
		case 24:
			switch data[(p)] {
			case 32:
				goto tr42
			case 95:
				goto tr43
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr43
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr43
				}
			default:
				goto tr43
			}
			goto tr1
		case 25:
			switch data[(p)] {
			case 44:
				goto tr44
			case 46:
				goto tr44
			case 61:
				goto tr45
			case 93:
				goto tr44
			case 95:
				goto tr44
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr44
				}
			case data[(p)] > 91:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr44
				}
			default:
				goto tr44
			}
			goto tr1
		case 32:
			switch data[(p)] {
			case 32:
				goto tr54
			case 61:
				goto tr46
			case 92:
				goto tr55
			}
			goto tr53
		case 33:
			switch data[(p)] {
			case 32:
				goto tr56
			case 61:
				goto tr46
			case 92:
				goto tr57
			}
			goto tr48
		case 34:
			switch data[(p)] {
			case 32:
				goto tr56
			case 61:
				goto tr46
			case 92:
				goto tr57
			case 95:
				goto tr58
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr58
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr58
				}
			default:
				goto tr58
			}
			goto tr48
		case 35:
			switch data[(p)] {
			case 32:
				goto tr56
			case 44:
				goto tr59
			case 46:
				goto tr59
			case 61:
				goto tr60
			case 92:
				goto tr57
			case 95:
				goto tr59
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr59
				}
			case data[(p)] > 93:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr59
				}
			default:
				goto tr59
			}
			goto tr48
		case 36:
			switch data[(p)] {
			case 32:
				goto tr62
			case 61:
				goto tr46
			case 92:
				goto tr63
			}
			goto tr61
		case 37:
			switch data[(p)] {
			case 32:
				goto tr64
			case 61:
				goto tr46
			case 92:
				goto tr65
			}
			goto tr47
		case 38:
			switch data[(p)] {
			case 32:
				goto tr64
			case 61:
				goto tr46
			case 92:
				goto tr65
			case 95:
				goto tr66
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr66
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr66
				}
			default:
				goto tr66
			}
			goto tr47
		case 39:
			switch data[(p)] {
			case 32:
				goto tr64
			case 44:
				goto tr67
			case 46:
				goto tr67
			case 61:
				goto tr60
			case 92:
				goto tr65
			case 95:
				goto tr67
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr67
				}
			case data[(p)] > 93:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr67
				}
			default:
				goto tr67
			}
			goto tr47
		case 26:
			switch data[(p)] {
			case 61:
				goto tr47
			case 92:
				goto tr47
			}
			goto tr46
		case 27:
			switch data[(p)] {
			case 61:
				goto tr48
			case 92:
				goto tr48
			}
			goto tr46
		case 28:
			if data[(p)] == 32 {
				goto tr50
			}
			goto tr49
		case 29:
			switch data[(p)] {
			case 32:
				goto tr50
			case 95:
				goto tr51
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr51
				}
			case data[(p)] > 90:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr51
				}
			default:
				goto tr51
			}
			goto tr49
		case 30:
			switch data[(p)] {
			case 32:
				goto tr50
			case 44:
				goto tr51
			case 46:
				goto tr51
			case 61:
				goto tr52
			case 93:
				goto tr51
			case 95:
				goto tr51
			}
			switch {
			case data[(p)] < 65:
				if 48 <= data[(p)] && data[(p)] <= 57 {
					goto tr51
				}
			case data[(p)] > 91:
				if 97 <= data[(p)] && data[(p)] <= 122 {
					goto tr51
				}
			default:
				goto tr51
			}
			goto tr49
		case 40:
			if data[(p)] == 32 {
				goto tr50
			}
			goto tr49
		}

	tr1:
		cs = 0
		goto _again
	tr46:
		cs = 0
		goto f15
	tr0:
		cs = 2
		goto _again
	tr2:
		cs = 3
		goto _again
	tr3:
		cs = 4
		goto _again
	tr4:
		cs = 5
		goto _again
	tr6:
		cs = 6
		goto _again
	tr5:
		cs = 6
		goto f0
	tr7:
		cs = 7
		goto f1
	tr11:
		cs = 8
		goto _again
	tr8:
		cs = 8
		goto f0
	tr12:
		cs = 9
		goto _again
	tr9:
		cs = 9
		goto f0
	tr10:
		cs = 10
		goto f2
	tr13:
		cs = 10
		goto f3
	tr17:
		cs = 11
		goto _again
	tr14:
		cs = 11
		goto f0
	tr18:
		cs = 12
		goto _again
	tr15:
		cs = 12
		goto f0
	tr16:
		cs = 13
		goto f4
	tr19:
		cs = 13
		goto f5
	tr23:
		cs = 14
		goto _again
	tr20:
		cs = 14
		goto f0
	tr24:
		cs = 15
		goto _again
	tr21:
		cs = 15
		goto f0
	tr22:
		cs = 16
		goto f6
	tr25:
		cs = 16
		goto f7
	tr29:
		cs = 17
		goto _again
	tr26:
		cs = 17
		goto f0
	tr30:
		cs = 18
		goto _again
	tr27:
		cs = 18
		goto f0
	tr28:
		cs = 19
		goto f8
	tr31:
		cs = 19
		goto f9
	tr35:
		cs = 20
		goto _again
	tr32:
		cs = 20
		goto f0
	tr36:
		cs = 21
		goto _again
	tr33:
		cs = 21
		goto f0
	tr34:
		cs = 22
		goto f10
	tr37:
		cs = 22
		goto f11
	tr40:
		cs = 23
		goto _again
	tr38:
		cs = 23
		goto f0
	tr42:
		cs = 24
		goto _again
	tr44:
		cs = 25
		goto _again
	tr43:
		cs = 25
		goto f0
	tr65:
		cs = 26
		goto _again
	tr63:
		cs = 26
		goto f20
	tr57:
		cs = 27
		goto _again
	tr55:
		cs = 27
		goto f20
	tr49:
		cs = 28
		goto _again
	tr50:
		cs = 29
		goto f0
	tr51:
		cs = 30
		goto _again
	tr39:
		cs = 31
		goto f12
	tr41:
		cs = 31
		goto f13
	tr45:
		cs = 32
		goto f14
	tr48:
		cs = 33
		goto f16
	tr53:
		cs = 33
		goto f19
	tr56:
		cs = 34
		goto f16
	tr54:
		cs = 34
		goto f19
	tr59:
		cs = 35
		goto f16
	tr58:
		cs = 35
		goto f22
	tr60:
		cs = 36
		goto f14
	tr47:
		cs = 37
		goto f16
	tr61:
		cs = 37
		goto f19
	tr64:
		cs = 38
		goto f16
	tr62:
		cs = 38
		goto f19
	tr67:
		cs = 39
		goto f16
	tr66:
		cs = 39
		goto f23
	tr52:
		cs = 40
		goto f17

	f0:
//line cef.rl:37

		mark = p

		goto _again
	f1:
//line cef.rl:40

		e.Version, _ = strconv.Atoi(data[mark:p])

		goto _again
	f3:
//line cef.rl:43

		e.DeviceVendor = replaceHeaderEscapes(data[mark:p])

		goto _again
	f5:
//line cef.rl:46

		e.DeviceProduct = replaceHeaderEscapes(data[mark:p])

		goto _again
	f7:
//line cef.rl:49

		e.DeviceVersion = replaceHeaderEscapes(data[mark:p])

		goto _again
	f9:
//line cef.rl:52

		e.DeviceEventClassID = replaceHeaderEscapes(data[mark:p])

		goto _again
	f11:
//line cef.rl:55

		e.Name = replaceHeaderEscapes(data[mark:p])

		goto _again
	f13:
//line cef.rl:58

		e.Severity = data[mark:p]

		goto _again
	f14:
//line cef.rl:61

		// A new extension key marks the end of the last extension value.
		if len(extKey) > 0 && extValueStart <= mark-1 {
			e.pushExtension(extKey, replaceExtensionEscapes(data[extValueStart:mark-1]))
			extKey, extValueStart, extValueEnd = "", 0, 0
		}
		extKey = data[mark:p]

		goto _again
	f20:
//line cef.rl:69

		extValueStart = p
		extValueEnd = p

		goto _again
	f16:
//line cef.rl:73

		extValueEnd = p + 1

		goto _again
	f15:
//line cef.rl:83

		recoveredErrs = append(recoveredErrs, fmt.Errorf("malformed value for %s at pos %d", extKey, p+1))
		(p)--
		cs = 28
		goto _again
	f17:
//line cef.rl:87

		extKey, extValueStart, extValueEnd = "", 0, 0
		// Resume processing at p, the start of the next extension key.
		p = mark
		cs = 24
		goto _again
	f2:
//line cef.rl:37

		mark = p

//line cef.rl:43

		e.DeviceVendor = replaceHeaderEscapes(data[mark:p])

		goto _again
	f4:
//line cef.rl:37

		mark = p

//line cef.rl:46

		e.DeviceProduct = replaceHeaderEscapes(data[mark:p])

		goto _again
	f6:
//line cef.rl:37

		mark = p

//line cef.rl:49

		e.DeviceVersion = replaceHeaderEscapes(data[mark:p])

		goto _again
	f8:
//line cef.rl:37

		mark = p

//line cef.rl:52

		e.DeviceEventClassID = replaceHeaderEscapes(data[mark:p])

		goto _again
	f10:
//line cef.rl:37

		mark = p

//line cef.rl:55

		e.Name = replaceHeaderEscapes(data[mark:p])

		goto _again
	f12:
//line cef.rl:37

		mark = p

//line cef.rl:58

		e.Severity = data[mark:p]

		goto _again
	f23:
//line cef.rl:37

		mark = p

//line cef.rl:73

		extValueEnd = p + 1

		goto _again
	f19:
//line cef.rl:69

		extValueStart = p
		extValueEnd = p

//line cef.rl:73

		extValueEnd = p + 1

		goto _again
	f22:
//line cef.rl:73

		extValueEnd = p + 1

//line cef.rl:37

		mark = p

		goto _again

	_again:
		if cs == 0 {
			goto _out
		}
		if (p)++; (p) != (pe) {
			goto _resume
		}
	_test_eof:
		{
		}
		if (p) == eof {
			switch _cef_eof_actions[cs] {
			case 22:
//line cef.rl:76

				// Reaching the EOF marks the end of the final extension value.
				if len(extKey) > 0 && extValueStart <= extValueEnd {
					e.pushExtension(extKey, replaceExtensionEscapes(data[extValueStart:extValueEnd]))
					extKey, extValueStart, extValueEnd = "", 0, 0
				}

			case 16:
//line cef.rl:83

				recoveredErrs = append(recoveredErrs, fmt.Errorf("malformed value for %s at pos %d", extKey, p+1))
				(p)--
				cs = 28
				goto _again

			case 19:
//line cef.rl:69

				extValueStart = p
				extValueEnd = p

//line cef.rl:76

				// Reaching the EOF marks the end of the final extension value.
				if len(extKey) > 0 && extValueStart <= extValueEnd {
					e.pushExtension(extKey, replaceExtensionEscapes(data[extValueStart:extValueEnd]))
					extKey, extValueStart, extValueEnd = "", 0, 0
				}

//line parser.go:847
			}
		}

	_out:
		{
		}
	}

//line cef.rl:145

	// Check if state machine completed.
	if cs < cef_first_final {
		// Reached an early end.
		if p == pe {
			return multierr.Append(multierr.Combine(recoveredErrs...), fmt.Errorf("unexpected end of CEF event"))
		}

		// Encountered invalid input.
		return multierr.Append(multierr.Combine(recoveredErrs...), fmt.Errorf("error in CEF event at pos %d", p+1))
	}

	return multierr.Combine(recoveredErrs...)
}
