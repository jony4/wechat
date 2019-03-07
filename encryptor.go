package wechat

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"time"
)

// EncryptMsg EncryptMsg
func EncryptMsg(random, rawXMLMsg []byte, appID, aesKey string) (encrtptMsg []byte, err error) {
	var key []byte
	key, err = aesKeyDecode(aesKey)
	if err != nil {
		return []byte{}, err
	}
	ciphertext := AESEncryptMsg(random, rawXMLMsg, appID, key)
	encrtptMsg = []byte(base64.StdEncoding.EncodeToString(ciphertext))
	return
}

// AESEncryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
func AESEncryptMsg(random, rawXMLMsg []byte, appID string, aesKey []byte) (ciphertext []byte) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	appIDOffset := 20 + len(rawXMLMsg)
	contentLen := appIDOffset + len(appID)
	amountToPad := BlockSize - contentLen&BlockMask
	plaintextLen := contentLen + amountToPad

	plaintext := make([]byte, plaintextLen)

	copy(plaintext[:16], random)
	encodeNetworkByteOrder(plaintext[16:20], uint32(len(rawXMLMsg)))
	copy(plaintext[20:], rawXMLMsg)
	copy(plaintext[appIDOffset:], appID)

	for i := contentLen; i < plaintextLen; i++ {
		plaintext[i] = byte(amountToPad)
	}

	block, err := aes.NewCipher(aesKey[:])
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, aesKey[:16])
	mode.CryptBlocks(plaintext, plaintext)

	ciphertext = plaintext
	return
}

// DecryptMsg DecryptMsg
func DecryptMsg(appID, encryptedMsg, aesKey string) (random, rawMsgXMLBytes []byte, err error) {
	var encryptedMsgBytes, key, getAppIDBytes []byte
	encryptedMsgBytes, err = base64.StdEncoding.DecodeString(encryptedMsg)
	if err != nil {
		return
	}
	key, err = aesKeyDecode(aesKey)
	if err != nil {
		return
	}
	random, rawMsgXMLBytes, getAppIDBytes, err = AESDecryptMsg(encryptedMsgBytes, key)
	if err != nil {
		err = fmt.Errorf("DecryptMsg failed, %v", err)
		return
	}
	if appID != string(getAppIDBytes) {
		err = fmt.Errorf("DecryptMsg appId caused failed")
		return
	}
	return
}

func aesKeyDecode(encodedAESKey string) (key []byte, err error) {
	if len(encodedAESKey) != 43 {
		err = fmt.Errorf("the length of encodedAESKey must be equal to 43")
		return
	}
	key, err = base64.StdEncoding.DecodeString(encodedAESKey + "=")
	if err != nil {
		return
	}
	if len(key) != 32 {
		err = fmt.Errorf("encodingAESKey invalid")
		return
	}
	return
}

// AESDecryptMsg ciphertext = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
func AESDecryptMsg(ciphertext []byte, aesKey []byte) (random, rawXMLMsg, appID []byte, err error) {
	const (
		BlockSize = 32            // PKCS#7
		BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
	)

	if len(ciphertext) < BlockSize {
		err = fmt.Errorf("the length of ciphertext too short: %d", len(ciphertext))
		return
	}
	if len(ciphertext)&BlockMask != 0 {
		err = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(ciphertext))
		return
	}

	plaintext := make([]byte, len(ciphertext)) // len(plaintext) >= BLOCK_SIZE

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return
	}
	mode := cipher.NewCBCDecrypter(block, aesKey[:16])
	mode.CryptBlocks(plaintext, ciphertext)

	amountToPad := int(plaintext[len(plaintext)-1])
	if amountToPad < 1 || amountToPad > BlockSize {
		err = fmt.Errorf("the amount to pad is incorrect: %d", amountToPad)
		return
	}
	plaintext = plaintext[:len(plaintext)-amountToPad]

	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(plaintext) <= 20 {
		err = fmt.Errorf("plaintext too short, the length is %d", len(plaintext))
		return
	}
	rawXMLMsgLen := int(decodeNetworkByteOrder(plaintext[16:20]))
	if rawXMLMsgLen < 0 {
		err = fmt.Errorf("incorrect msg length: %d", rawXMLMsgLen)
		return
	}
	appIDOffset := 20 + rawXMLMsgLen
	if len(plaintext) <= appIDOffset {
		err = fmt.Errorf("msg length too large: %d", rawXMLMsgLen)
		return
	}

	random = plaintext[:16:20]
	rawXMLMsg = plaintext[20:appIDOffset:appIDOffset]
	appID = plaintext[appIDOffset:]
	return
}

// 把整数 n 格式化成 4 字节的网络字节序
func encodeNetworkByteOrder(orderBytes []byte, n uint32) {
	orderBytes[0] = byte(n >> 24)
	orderBytes[1] = byte(n >> 16)
	orderBytes[2] = byte(n >> 8)
	orderBytes[3] = byte(n)
}

// 从 4 字节的网络字节序里解析出整数
func decodeNetworkByteOrder(orderBytes []byte) (n uint32) {
	return uint32(orderBytes[0])<<24 |
		uint32(orderBytes[1])<<16 |
		uint32(orderBytes[2])<<8 |
		uint32(orderBytes[3])
}

// MD5Sum 计算 32 位长度的 MD5 sum
func MD5Sum(txt string) (sum string) {
	h := md5.New()
	buf := bufio.NewWriterSize(h, 128)
	buf.WriteString(txt)
	buf.Flush()
	sign := make([]byte, hex.EncodedLen(h.Size()))
	hex.Encode(sign, h.Sum(nil))
	sum = string(bytes.ToUpper(sign))
	return
}

// Signature sha1签名
func Signature(params ...string) string {
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

// RandomStr random string
func RandomStr(length int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}
