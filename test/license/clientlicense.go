package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// LicenseData 结构定义授权信息数据结构
type LicenseData struct {
	SerialNumber  string `json:"serial_number"`
	Signature     string `json:"signature"`
	EncryptedTime string `json:"encrypted_time"`
}

func main() {
	fmt.Println("client======================================")
	serialNumber := client()
	fmt.Println("")
	fmt.Println("server======================================")
	server(serialNumber)
	fmt.Println("")
	fmt.Println("verify======================================")
	verify()
}

func client() string {
	// 1. 生成随机的AES密钥
	aesKey := generateAESKey()
	fmt.Println(len(aesKey), aesKey)
	hardwareInfo := "1234567890"

	// 2. 加密机器硬件信息
	encryptedHardwareInfo, err := encryptAES([]byte(hardwareInfo), aesKey)
	checkError("Failed to encrypt hardware info:", err)

	// 3. 拼接成机器序列号SN
	serialNumber := hex.EncodeToString(aesKey) + hex.EncodeToString(encryptedHardwareInfo)

	fmt.Println("Serial Number:", len(serialNumber), serialNumber)
	return serialNumber
}

func server(serialNumber string) {
	// 读取客户端生成的机器序列号SN
	//clientSerialNumber := serialNumber
	// 读取RSA私钥
	privateKeyBytes, err := ioutil.ReadFile("D:\\work\\code\\go\\src\\HWH\\test\\license\\private.pem")
	checkError("Failed to read private key file:", err)

	// 解析RSA私钥
	privateKey, err := ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	checkError("Failed to parse private key:", err)

	// 计算消息的哈希值
	hashedSerialNumber := sha256.Sum256([]byte(serialNumber))

	// 1. 使用RSA私钥对机器序列号SN进行签名
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashedSerialNumber[:])
	checkError("Failed to sign serial number:", err)

	// 2. 利用SN中携带的AES密钥对授权时间进行加密
	authorizationTime := time.Now().Add(time.Hour * 24 * 30) // 有效期30天
	encryptedTime, err := encryptAES([]byte(authorizationTime.Format(time.RFC3339)), hashedSerialNumber[:])
	checkError("Failed to encrypt authorization time:", err)

	// 3. 构建授权信息数据结构
	licenseData := LicenseData{
		SerialNumber:  serialNumber,
		Signature:     hex.EncodeToString(signature),
		EncryptedTime: hex.EncodeToString(encryptedTime),
	}

	saveLicense := licenseData.SerialNumber + "." + licenseData.Signature + "." + licenseData.EncryptedTime
	// 保存授权信息到文件
	err = ioutil.WriteFile("D:\\work\\code\\go\\src\\HWH\\test\\license\\license.txt", []byte(saveLicense), 0644)
	checkError("Failed to save license to file:", err)

	fmt.Println("License generated and saved successfully.")
}

func verify() {
	// 读取授权文件内容
	licenseBytes, err := ioutil.ReadFile("D:\\work\\code\\go\\src\\HWH\\test\\license\\license.txt")
	checkError("Failed to read license file:", err)
	license := strings.Replace(string(licenseBytes), ".", " ", -1)
	// 解析授权信息
	licenseData := strings.Fields(license)

	serialNumber := licenseData[0]
	signature, err := hex.DecodeString(licenseData[1])
	checkError("Failed to decode signature:", err)
	encryptedTime, err := hex.DecodeString(licenseData[2])
	checkError("Failed to decode encrypted time:", err)

	// 1. 使用RSA公钥验证签名内容
	if !verifySignature(serialNumber, signature) {
		log.Fatal("Signature verification failed")
	}

	// 2. 生成硬件信息
	//hardwareInfo := "1234567890"

	// 3. 比较硬件信息
	//if hardwareInfo != serialNumber[32:] {
	//	log.Fatal("Hardware information does not match")
	//}

	// 4. 使用AES密钥解密授权时间
	aesKey, err := hex.DecodeString(serialNumber[:64])
	checkError("Failed to decode AES key:", err)
	decryptedTime, err := decryptAES(encryptedTime, aesKey)
	checkError("Failed to decrypt authorization time:", err)

	// 5. 判断当前时间是否在授权时间内
	authorizationTime, err := time.Parse(time.RFC3339, string(decryptedTime))
	checkError("Failed to parse authorization time:", err)
	if time.Now().After(authorizationTime) {
		log.Fatal("Authorization expired")
	}

	// 6. 修改最后一次校验的时间为当前时间
	// 在这里你可以执行你的应用程序逻辑

	fmt.Println("License verification passed")
}

// 使用RSA公钥验证签名内容
func verifySignature(serialNumber string, signature []byte) bool {
	// 读取RSA公钥
	publicKeyBytes, err := ioutil.ReadFile("D:\\work\\code\\go\\src\\HWH\\test\\license\\public.pem")
	checkError("Failed to read public key file:", err)
	publicKey, err := ParseRSAPublicKeyFromPEM(publicKeyBytes)
	checkError("Failed to parse public key:", err)

	// 计算SHA-256哈希
	hashed := sha256.Sum256([]byte(serialNumber))

	// 验证签名
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		fmt.Println("Signature verification failed:", err)
		return false
	}

	return true
}

// 生成随机的32字节AES密钥
func generateAESKey() []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatal("Failed to generate AES key:", err)
	}
	return key
}

// ClientEncryptAES 使用AES加密数据，使用PKCS7填充算法
func encryptAES(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// 计算需要填充的字节数
	padding := aes.BlockSize - (len(data) % aes.BlockSize)
	// 创建填充后的数据
	paddedData := make([]byte, len(data)+padding)
	copy(paddedData, data)

	// 填充
	for i := len(data); i < len(paddedData); i++ {
		paddedData[i] = byte(padding)
	}

	ciphertext := make([]byte, aes.BlockSize+len(paddedData))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedData)

	return ciphertext, nil
}

// 使用AES解密数据，移除PKCS7填充
func decryptAES(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(data) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)

	// 移除填充
	padding := int(data[len(data)-1])
	if padding < 1 || padding > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding size")
	}
	if padding > len(data) {
		return nil, fmt.Errorf("padding size is larger than the data length")
	}
	// 检查所有的填充字节是否都等于填充的大小
	for i := 0; i < padding; i++ {
		if data[len(data)-1-i] != byte(padding) {
			return nil, fmt.Errorf("invalid padding")
		}
	}
	return data[:len(data)-padding], nil
}

// ParseRSAPrivateKeyFromPEM parses an RSA private key from a PEM encoded data.
func ParseRSAPrivateKeyFromPEM(pemData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	if x509.IsEncryptedPEMBlock(block) {
		// Encrypted PEM, prompt for password
		return nil, fmt.Errorf("encrypted PEM blocks are not supported")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// ParseRSAPublicKeyFromPEM parses an RSA public key from a PEM encoded data.
func ParseRSAPublicKeyFromPEM(pemData []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("key is not RSA")
	}

	return rsaPub, nil
}

func checkError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
