package comp

import (
	"bytes"
)

/*
  StringBuilder struct.
	  Usage:
		builder := NewStringBuilder()
		builder.Append("a").Append("b").Append("c")
		s := builder.String()
		println(s)
	  print:
		abc
啊啊啊
*/

// StringBuilder
type StringBuilder struct {
	buffer bytes.Buffer
}

//NewStringBuilderCap 带容器长度的创建
func NewStringBuilderCap(capnum int) *StringBuilder {
	builder := StringBuilder{
		buffer: *bytes.NewBuffer(make([]byte, 0, capnum)),
	}
	return &builder
}

func NewStringBuilder() *StringBuilder {
	var builder StringBuilder
	return &builder
}

func NewStringBuilderString(str *String) *StringBuilder {
	var builder StringBuilder
	builder.buffer.WriteString(str.ToString())
	return &builder
}

func (builder *StringBuilder) Appendln(s string) *StringBuilder {
	builder.buffer.WriteString(s + "\n")
	return builder
}

func (builder *StringBuilder) Append(s string) *StringBuilder {
	builder.buffer.WriteString(s)
	return builder
}

func (builder *StringBuilder) AppendRune(s rune) *StringBuilder {
	builder.buffer.WriteRune(s)
	return builder
}

func (builder *StringBuilder) AppendStrings(ss ...string) *StringBuilder {
	for i := range ss {
		builder.buffer.WriteString(ss[i])
	}
	return builder
}

func (builder *StringBuilder) AppendInt(i int) *StringBuilder {
	builder.buffer.WriteString(NewStringInt(i).ToString())
	return builder
}

func (builder *StringBuilder) AppendInt64(i int64) *StringBuilder {
	builder.buffer.WriteString(NewStringInt64(i).ToString())
	return builder
}

func (builder *StringBuilder) AppendFloat64(f float64) *StringBuilder {
	builder.buffer.WriteString(NewStringFloat64(f).ToString())
	return builder
}

func (builder *StringBuilder) Replace(old, new string) *StringBuilder {
	str := NewString(builder.ToString()).Replace(old, new)
	builder.Clear()
	builder.buffer.WriteString(str.ToString())
	return builder
}

func (builder *StringBuilder) RemoveLast() *StringBuilder {
	str1 := NewString(builder.ToString())
	builder.Clear()
	str2 := str1.Substring(0, str1.Len()-1)
	builder.buffer.WriteString(str2.ToString())
	return builder
}

func (builder *StringBuilder) Clear() *StringBuilder {
	builder.buffer.Reset()
	return builder
}

func (builder *StringBuilder) ToString() string {
	return builder.buffer.String()
}

//IsEmpty 是否为空字符串
func (builder *StringBuilder) IsEmpty() bool {
	return builder.buffer.Len() == 0
}
