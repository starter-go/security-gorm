package vericode

import (
	"bytes"
	"crypto/sha256"
	"strconv"

	"github.com/starter-go/base/lang"
)

type computer struct {
	addr  string
	code  Code
	nonce []byte
	salt  []byte
	t1    lang.Time
	t2    lang.Time
}

// 计算经验和
func (inst *computer) sum() []byte {

	tn1 := inst.t1.Int()
	tn2 := inst.t2.Int()

	builder := &bytes.Buffer{}
	builder.WriteString(strconv.FormatInt(tn1, 10))
	builder.WriteString(inst.addr)
	builder.WriteString(string(inst.code))
	builder.WriteString(strconv.FormatInt(tn2, 10))
	builder.Write(inst.nonce)
	builder.Write(inst.salt)

	sum := sha256.Sum256(builder.Bytes())
	return sum[:]
}
