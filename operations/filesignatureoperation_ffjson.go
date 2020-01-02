// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: filesignatureoperation.go

package operations

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *FileSignatureOperation) MarshalJSON() ([]byte, error) {
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
func (j *FileSignatureOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "signature_account":`)

	{

		obj, err = j.SignatureAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"file_id":`)

	{

		obj, err = j.FileID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"signature":`)
	fflib.WriteJsonString(buf, string(j.Signature))
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			/* Struct fall back. type=types.AssetAmount kind=struct */
			buf.WriteString(`"fee":`)
			err = buf.Encode(j.Fee)
			if err != nil {
				return err
			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtFileSignatureOperationbase = iota
	ffjtFileSignatureOperationnosuchkey

	ffjtFileSignatureOperationSignatureAccount

	ffjtFileSignatureOperationFileID

	ffjtFileSignatureOperationSignature

	ffjtFileSignatureOperationFee
)

var ffjKeyFileSignatureOperationSignatureAccount = []byte("signature_account")

var ffjKeyFileSignatureOperationFileID = []byte("file_id")

var ffjKeyFileSignatureOperationSignature = []byte("signature")

var ffjKeyFileSignatureOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *FileSignatureOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *FileSignatureOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtFileSignatureOperationbase
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
				currentKey = ffjtFileSignatureOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'f':

					if bytes.Equal(ffjKeyFileSignatureOperationFileID, kn) {
						currentKey = ffjtFileSignatureOperationFileID
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFileSignatureOperationFee, kn) {
						currentKey = ffjtFileSignatureOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyFileSignatureOperationSignatureAccount, kn) {
						currentKey = ffjtFileSignatureOperationSignatureAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyFileSignatureOperationSignature, kn) {
						currentKey = ffjtFileSignatureOperationSignature
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyFileSignatureOperationFee, kn) {
					currentKey = ffjtFileSignatureOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyFileSignatureOperationSignature, kn) {
					currentKey = ffjtFileSignatureOperationSignature
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyFileSignatureOperationFileID, kn) {
					currentKey = ffjtFileSignatureOperationFileID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyFileSignatureOperationSignatureAccount, kn) {
					currentKey = ffjtFileSignatureOperationSignatureAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtFileSignatureOperationnosuchkey
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

				case ffjtFileSignatureOperationSignatureAccount:
					goto handle_SignatureAccount

				case ffjtFileSignatureOperationFileID:
					goto handle_FileID

				case ffjtFileSignatureOperationSignature:
					goto handle_Signature

				case ffjtFileSignatureOperationFee:
					goto handle_Fee

				case ffjtFileSignatureOperationnosuchkey:
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

handle_SignatureAccount:

	/* handler: j.SignatureAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.SignatureAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FileID:

	/* handler: j.FileID type=types.FileID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.FileID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Signature:

	/* handler: j.Signature type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.Signature = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		/* Falling back. type=types.AssetAmount kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &j.Fee)
		if err != nil {
			return fs.WrapErr(err)
		}
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
