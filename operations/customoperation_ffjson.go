// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: customoperation.go

package operations

import (
	"bytes"
	"fmt"
	"github.com/gkany/graph-sdk/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *CustomOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *CustomOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "payer":`)

	{

		obj, err = j.Payer.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"required_auths":`)
	if j.RequiredAuths != nil {
		buf.WriteString(`[`)
		for i, v := range j.RequiredAuths {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				obj, err = v.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"id":`)
	fflib.FormatBits2(buf, uint64(j.ID), 10, false)
	buf.WriteString(`,"data":`)

	{

		obj, err = j.Data.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			buf.WriteString(`"fee":`)

			{

				err = j.Fee.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtCustomOperationbase = iota
	ffjtCustomOperationnosuchkey

	ffjtCustomOperationPayer

	ffjtCustomOperationRequiredAuths

	ffjtCustomOperationID

	ffjtCustomOperationData

	ffjtCustomOperationFee
)

var ffjKeyCustomOperationPayer = []byte("payer")

var ffjKeyCustomOperationRequiredAuths = []byte("required_auths")

var ffjKeyCustomOperationID = []byte("id")

var ffjKeyCustomOperationData = []byte("data")

var ffjKeyCustomOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *CustomOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *CustomOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtCustomOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtCustomOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'd':

					if bytes.Equal(ffjKeyCustomOperationData, kn) {
						currentKey = ffjtCustomOperationData
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyCustomOperationFee, kn) {
						currentKey = ffjtCustomOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'i':

					if bytes.Equal(ffjKeyCustomOperationID, kn) {
						currentKey = ffjtCustomOperationID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyCustomOperationPayer, kn) {
						currentKey = ffjtCustomOperationPayer
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'r':

					if bytes.Equal(ffjKeyCustomOperationRequiredAuths, kn) {
						currentKey = ffjtCustomOperationRequiredAuths
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyCustomOperationFee, kn) {
					currentKey = ffjtCustomOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCustomOperationData, kn) {
					currentKey = ffjtCustomOperationData
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCustomOperationID, kn) {
					currentKey = ffjtCustomOperationID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyCustomOperationRequiredAuths, kn) {
					currentKey = ffjtCustomOperationRequiredAuths
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCustomOperationPayer, kn) {
					currentKey = ffjtCustomOperationPayer
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtCustomOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtCustomOperationPayer:
					goto handle_Payer

				case ffjtCustomOperationRequiredAuths:
					goto handle_RequiredAuths

				case ffjtCustomOperationID:
					goto handle_ID

				case ffjtCustomOperationData:
					goto handle_Data

				case ffjtCustomOperationFee:
					goto handle_Fee

				case ffjtCustomOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_Payer:

	/* handler: j.Payer type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Payer.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_RequiredAuths:

	/* handler: j.RequiredAuths type=types.AccountIDs kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for AccountIDs", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.RequiredAuths = nil
		} else {

			j.RequiredAuths = []types.AccountID{}

			wantVal := true

			for {

				var tmpJRequiredAuths types.AccountID

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJRequiredAuths type=types.AccountID kind=struct quoted=false*/

				{
					if tok == fflib.FFTok_null {

					} else {

						tbuf, err := fs.CaptureField(tok)
						if err != nil {
							return fs.WrapErr(err)
						}

						err = tmpJRequiredAuths.UnmarshalJSON(tbuf)
						if err != nil {
							return fs.WrapErr(err)
						}
					}
					state = fflib.FFParse_after_value
				}

				j.RequiredAuths = append(j.RequiredAuths, tmpJRequiredAuths)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ID:

	/* handler: j.ID type=types.UInt16 kind=uint16 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.ID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Data:

	/* handler: j.Data type=types.Buffer kind=slice quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Data.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.Fee = nil

		} else {

			if j.Fee == nil {
				j.Fee = new(types.AssetAmount)
			}

			err = j.Fee.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
