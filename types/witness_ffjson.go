// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: witness.go

package types

import (
	"bytes"
	"errors"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *Witness) MarshalJSON() ([]byte, error) {
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
func (j *Witness) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"id":`)

	{

		obj, err = j.ID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"witness_account":`)

	{

		obj, err = j.WitnessAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"last_aslot":`)
	fflib.FormatBits2(buf, uint64(j.LastAslot), 10, false)
	buf.WriteString(`,"signing_key":`)

	{

		obj, err = j.SigningKey.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte(',')
	if j.PayVestingBalance != nil {
		if true {
			buf.WriteString(`"pay_vb":`)

			{

				obj, err = j.PayVestingBalance.MarshalJSON()
				if err != nil {
					return err
				}
				buf.Write(obj)

			}
			buf.WriteByte(',')
		}
	}
	buf.WriteString(`"vote_id":`)

	{

		obj, err = j.VoteID.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"total_votes":`)
	fflib.FormatBits2(buf, uint64(j.TotalVotes), 10, false)
	buf.WriteString(`,"url":`)

	{

		obj, err = j.URL.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"total_missed":`)
	fflib.FormatBits2(buf, uint64(j.TotalMissed), 10, false)
	buf.WriteString(`,"last_confirmed_block_num":`)
	fflib.FormatBits2(buf, uint64(j.LastConfirmedBlockNum), 10, false)
	if j.WorkStatus {
		buf.WriteString(`,"work_status":true`)
	} else {
		buf.WriteString(`,"work_status":false`)
	}
	buf.WriteString(`,"next_maintenance_time":`)

	{

		obj, err = j.NextMaintenanceTime.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"supporters":`)

	{

		obj, err = j.Supporters.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjtWitnessbase = iota
	ffjtWitnessnosuchkey

	ffjtWitnessID

	ffjtWitnessWitnessAccount

	ffjtWitnessLastAslot

	ffjtWitnessSigningKey

	ffjtWitnessPayVestingBalance

	ffjtWitnessVoteID

	ffjtWitnessTotalVotes

	ffjtWitnessURL

	ffjtWitnessTotalMissed

	ffjtWitnessLastConfirmedBlockNum

	ffjtWitnessWorkStatus

	ffjtWitnessNextMaintenanceTime

	ffjtWitnessSupporters
)

var ffjKeyWitnessID = []byte("id")

var ffjKeyWitnessWitnessAccount = []byte("witness_account")

var ffjKeyWitnessLastAslot = []byte("last_aslot")

var ffjKeyWitnessSigningKey = []byte("signing_key")

var ffjKeyWitnessPayVestingBalance = []byte("pay_vb")

var ffjKeyWitnessVoteID = []byte("vote_id")

var ffjKeyWitnessTotalVotes = []byte("total_votes")

var ffjKeyWitnessURL = []byte("url")

var ffjKeyWitnessTotalMissed = []byte("total_missed")

var ffjKeyWitnessLastConfirmedBlockNum = []byte("last_confirmed_block_num")

var ffjKeyWitnessWorkStatus = []byte("work_status")

var ffjKeyWitnessNextMaintenanceTime = []byte("next_maintenance_time")

var ffjKeyWitnessSupporters = []byte("supporters")

// UnmarshalJSON umarshall json - template of ffjson
func (j *Witness) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *Witness) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtWitnessbase
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
				currentKey = ffjtWitnessnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'i':

					if bytes.Equal(ffjKeyWitnessID, kn) {
						currentKey = ffjtWitnessID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'l':

					if bytes.Equal(ffjKeyWitnessLastAslot, kn) {
						currentKey = ffjtWitnessLastAslot
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWitnessLastConfirmedBlockNum, kn) {
						currentKey = ffjtWitnessLastConfirmedBlockNum
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'n':

					if bytes.Equal(ffjKeyWitnessNextMaintenanceTime, kn) {
						currentKey = ffjtWitnessNextMaintenanceTime
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'p':

					if bytes.Equal(ffjKeyWitnessPayVestingBalance, kn) {
						currentKey = ffjtWitnessPayVestingBalance
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 's':

					if bytes.Equal(ffjKeyWitnessSigningKey, kn) {
						currentKey = ffjtWitnessSigningKey
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWitnessSupporters, kn) {
						currentKey = ffjtWitnessSupporters
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 't':

					if bytes.Equal(ffjKeyWitnessTotalVotes, kn) {
						currentKey = ffjtWitnessTotalVotes
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWitnessTotalMissed, kn) {
						currentKey = ffjtWitnessTotalMissed
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'u':

					if bytes.Equal(ffjKeyWitnessURL, kn) {
						currentKey = ffjtWitnessURL
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'v':

					if bytes.Equal(ffjKeyWitnessVoteID, kn) {
						currentKey = ffjtWitnessVoteID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyWitnessWitnessAccount, kn) {
						currentKey = ffjtWitnessWitnessAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyWitnessWorkStatus, kn) {
						currentKey = ffjtWitnessWorkStatus
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffjKeyWitnessSupporters, kn) {
					currentKey = ffjtWitnessSupporters
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWitnessNextMaintenanceTime, kn) {
					currentKey = ffjtWitnessNextMaintenanceTime
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessWorkStatus, kn) {
					currentKey = ffjtWitnessWorkStatus
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessLastConfirmedBlockNum, kn) {
					currentKey = ffjtWitnessLastConfirmedBlockNum
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessTotalMissed, kn) {
					currentKey = ffjtWitnessTotalMissed
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyWitnessURL, kn) {
					currentKey = ffjtWitnessURL
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessTotalVotes, kn) {
					currentKey = ffjtWitnessTotalVotes
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWitnessVoteID, kn) {
					currentKey = ffjtWitnessVoteID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyWitnessPayVestingBalance, kn) {
					currentKey = ffjtWitnessPayVestingBalance
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessSigningKey, kn) {
					currentKey = ffjtWitnessSigningKey
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessLastAslot, kn) {
					currentKey = ffjtWitnessLastAslot
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyWitnessWitnessAccount, kn) {
					currentKey = ffjtWitnessWitnessAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyWitnessID, kn) {
					currentKey = ffjtWitnessID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtWitnessnosuchkey
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

				case ffjtWitnessID:
					goto handle_ID

				case ffjtWitnessWitnessAccount:
					goto handle_WitnessAccount

				case ffjtWitnessLastAslot:
					goto handle_LastAslot

				case ffjtWitnessSigningKey:
					goto handle_SigningKey

				case ffjtWitnessPayVestingBalance:
					goto handle_PayVestingBalance

				case ffjtWitnessVoteID:
					goto handle_VoteID

				case ffjtWitnessTotalVotes:
					goto handle_TotalVotes

				case ffjtWitnessURL:
					goto handle_URL

				case ffjtWitnessTotalMissed:
					goto handle_TotalMissed

				case ffjtWitnessLastConfirmedBlockNum:
					goto handle_LastConfirmedBlockNum

				case ffjtWitnessWorkStatus:
					goto handle_WorkStatus

				case ffjtWitnessNextMaintenanceTime:
					goto handle_NextMaintenanceTime

				case ffjtWitnessSupporters:
					goto handle_Supporters

				case ffjtWitnessnosuchkey:
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

handle_ID:

	/* handler: j.ID type=types.WitnessID kind=struct quoted=false*/

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

handle_WitnessAccount:

	/* handler: j.WitnessAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.WitnessAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LastAslot:

	/* handler: j.LastAslot type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.LastAslot.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_SigningKey:

	/* handler: j.SigningKey type=types.PublicKey kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.SigningKey.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_PayVestingBalance:

	/* handler: j.PayVestingBalance type=types.VestingBalanceID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.PayVestingBalance = nil

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			if j.PayVestingBalance == nil {
				j.PayVestingBalance = new(VestingBalanceID)
			}

			err = j.PayVestingBalance.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_VoteID:

	/* handler: j.VoteID type=types.VoteID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.VoteID.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TotalVotes:

	/* handler: j.TotalVotes type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.TotalVotes.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_URL:

	/* handler: j.URL type=types.String kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.URL.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_TotalMissed:

	/* handler: j.TotalMissed type=types.UInt64 kind=uint64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.TotalMissed.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_LastConfirmedBlockNum:

	/* handler: j.LastConfirmedBlockNum type=types.UInt32 kind=uint32 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.LastConfirmedBlockNum.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WorkStatus:

	/* handler: j.WorkStatus type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.WorkStatus = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.WorkStatus = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_NextMaintenanceTime:

	/* handler: j.NextMaintenanceTime type=types.Time kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.NextMaintenanceTime.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Supporters:

	/* handler: j.Supporters type=types.SupporterType kind=map quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Supporters.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
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
