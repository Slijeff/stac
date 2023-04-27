package parser

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Webhook struct {
	// X-Github-Delivery
	Delivery string

	// X-GitHub-Event
	Event string

	// X-Hub-Signature-256 (using this only)
	Signature_256 string

	// Payload
	Payload []byte
}

const sigPrefix = "sha256="
const sigLength = 32 // len(hex(signature))

func sign(secret, body []byte) []byte {
	computed := hmac.New(sha256.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}

func (h *Webhook) Verify(secret []byte) bool {
	if len(h.Signature_256) != (2*sigLength+len(sigPrefix)) || !strings.HasPrefix(h.Signature_256, sigPrefix) {
		fmt.Println("Format doesn't match")
		return false
	}

	given := make([]byte, sigLength)
	// convert string to hex
	hex.Decode(given, []byte(h.Signature_256[len(sigPrefix):]))

	computed := sign(secret, h.Payload)

	return hmac.Equal(computed, given)
}

func (h *Webhook) ConvertPayload(dst interface{}) error {
	return json.Unmarshal(h.Payload, dst)
}

func LoadWebhook(r *http.Request) (hook *Webhook, err error) {
	hook = new(Webhook)
	if !strings.EqualFold(r.Method, "POST") {
		return nil, errors.New("request is not POST")
	}
	if hook.Delivery = r.Header.Get("X-GitHub-Delivery"); len(hook.Delivery) == 0 {
		return nil, errors.New("no X-GitHub-Delivery field")
	}
	if hook.Event = r.Header.Get("X-GitHub-Event"); len(hook.Event) == 0 {
		return nil, errors.New("no X-GitHub-Event field")
	}
	if hook.Signature_256 = r.Header.Get("X-Hub-Signature-256"); len(hook.Signature_256) == 0 {
		return nil, errors.New("no X-Hub-Signature-256 field")
	}

	hook.Payload, err = io.ReadAll(r.Body)
	return hook, err
}

func Parse(secret []byte, r *http.Request) (hook *Webhook, err error) {
	h, e := LoadWebhook(r)
	if e == nil && !h.Verify(secret) {
		e = errors.New("invalid signature")
	}
	return h, e
}
