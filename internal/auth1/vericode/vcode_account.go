package vericode

import (
	"bytes"
	"context"
	"fmt"
	"strconv"

	"github.com/starter-go/base/lang"
	"github.com/starter-go/i18n"
	"github.com/starter-go/keyvalues"
	"github.com/starter-go/mails"
)

// Account ... 包含关于某个具体账号的信息
type Account struct {
	vc        *Context
	cc        context.Context
	entry     keyvalues.JSONEntry
	addr      string
	languages []i18n.Language
}

// Put ...
func (inst *Account) Put(item *DataObject) error {
	if item == nil {
		return fmt.Errorf("item is nil")
	}
	opt := &keyvalues.Options{
		MaxAge: inst.vc.maxAge,
	}
	return inst.entry.Put(item, opt)
}

// Fetch ...
func (inst *Account) Fetch() (*DataObject, error) {
	item := &DataObject{}
	err := inst.entry.Get(item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Verify 验证参数中的验证码是否正确
func (inst *Account) Verify(v *Verification) error {

	verifier := &codeVerifier{
		acc:   inst,
		code:  v.Code,
		nonce: v.Nonce,
	}

	steps := make([]func() error, 0)
	steps = append(steps, verifier.fetch)
	steps = append(steps, verifier.checkTTL)
	steps = append(steps, verifier.checkTime)
	steps = append(steps, verifier.checkSum)
	steps = append(steps, verifier.update)

	for _, step := range steps {
		err := step()
		if err != nil {
			return err
		}
	}
	return nil
}

// SendNewCode 生成并发送验证码;
// 注意：出于安全原因，返回的 Verification 结构中不包含验证码
func (inst *Account) SendNewCode() (*Verification, error) {
	cm := &codeMaker{acc: inst}
	steps := make([]func() error, 0)
	steps = append(steps, cm.prepare)
	steps = append(steps, cm.generate)
	steps = append(steps, cm.compute)
	steps = append(steps, cm.send)
	steps = append(steps, cm.save)
	for _, step := range steps {
		err := step()
		if err != nil {
			return nil, err
		}
	}
	return cm.result()
}

////////////////////////////////////////////////////////////////////////////////

type codeVerifier struct {
	acc   *Account
	code  Code
	nonce []byte
	data  *DataObject
}

func (inst *codeVerifier) fetch() error {
	have, err := inst.acc.Fetch()
	if err == nil {
		inst.data = have
	}
	return err
}

func (inst *codeVerifier) checkSum() error {
	have := inst.data
	com := &computer{
		addr:  inst.acc.addr,
		code:  inst.code,
		nonce: inst.nonce,
		salt:  have.Salt.Bytes(),
		t1:    have.Time1,
		t2:    have.Time2,
	}
	sum1 := com.sum()
	sum2 := have.Sum.Bytes()
	if !bytes.Equal(sum1, sum2) {
		return fmt.Errorf("bad verification code")
	}
	return nil
}

func (inst *codeVerifier) checkTime() error {
	now := lang.Now()
	t1 := inst.data.Time1
	t2 := inst.data.Time2
	if (t1 <= now) && (now <= t2) {
		return nil
	}
	return fmt.Errorf("expired")
}

func (inst *codeVerifier) checkTTL() error {
	// 每次检查都要 ttl--
	ttl := inst.data.TTL
	if ttl > 0 {
		inst.data.TTL--
		return nil
	}
	return fmt.Errorf("timeout")
}

func (inst *codeVerifier) update() error {
	return inst.acc.Put(inst.data)
}

////////////////////////////////////////////////////////////////////////////////

type codeMaker struct {
	acc *Account

	sum   []byte
	salt  []byte
	nonce []byte
	code  Code
	t1    lang.Time
	t2    lang.Time
}

func (inst *codeMaker) prepare() error {

	random := inst.acc.vc.random

	// nonce
	nonce := make([]byte, 16)
	random.Reader().Read(nonce)

	// salt
	salt := make([]byte, 32)
	random.Reader().Read(salt)

	// time
	age := inst.acc.vc.maxAge
	now := lang.Now()

	// hold
	inst.t1 = now
	inst.t2 = now.Add(age)
	inst.nonce = nonce
	inst.salt = salt
	return nil
}

// 生成验证码
func (inst *codeMaker) generate() error {
	length := 6
	random := inst.acc.vc.random
	n := random.NextInt64()
	if n < 0 {
		n = -n
	}
	str := strconv.FormatInt(n, 10)
	if len(str) > length {
		str = str[0:length]
	}
	inst.code = Code(str)
	return nil
}

func (inst *codeMaker) compute() error {
	com := &computer{
		addr:  inst.acc.addr,
		code:  inst.code,
		t1:    inst.t1,
		t2:    inst.t2,
		salt:  inst.salt,
		nonce: inst.nonce,
	}
	sum := com.sum()
	inst.sum = sum
	return nil
}

func (inst *codeMaker) save() error {
	item := &DataObject{
		Addr:  inst.acc.addr,
		Salt:  lang.Base64FromBytes(inst.salt),
		Sum:   lang.HexFromBytes(inst.sum),
		Time1: inst.t1,
		Time2: inst.t2,
		TTL:   inst.acc.vc.maxTTL,
	}
	return inst.acc.Put(item)
}

func (inst *codeMaker) send() error {

	addr := inst.acc.addr
	lang := inst.acc.languages
	code := inst.code

	msg, err := inst.acc.vc.NewMessage(code, lang...)
	if err != nil {
		return err
	}

	msg.ToAddresses = []mails.Address{mails.Address(addr)}
	cc := inst.acc.cc
	sender := inst.acc.vc.sender
	return sender.Send(cc, msg)
}

func (inst *codeMaker) result() (*Verification, error) {
	ver := &Verification{
		Nonce: inst.nonce,
		Addr:  inst.acc.addr,
		Code:  "", // 不包含验证码
	}
	return ver, nil
}
