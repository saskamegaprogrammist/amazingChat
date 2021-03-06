// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson657d4933DecodeGithubComSaskamegaprogrammistAmazingChatModels(in *jlexer.Lexer, out *IdModel) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson657d4933EncodeGithubComSaskamegaprogrammistAmazingChatModels(out *jwriter.Writer, in IdModel) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v IdModel) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson657d4933EncodeGithubComSaskamegaprogrammistAmazingChatModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v IdModel) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson657d4933EncodeGithubComSaskamegaprogrammistAmazingChatModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *IdModel) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson657d4933DecodeGithubComSaskamegaprogrammistAmazingChatModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *IdModel) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson657d4933DecodeGithubComSaskamegaprogrammistAmazingChatModels(l, v)
}
