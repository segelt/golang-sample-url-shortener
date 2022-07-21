package hashutility

import (
	"crypto/md5"
	"encoding/hex"
)

func HashStr(text string) string {
	hash := md5.Sum([]byte(text))
	encodedstr := hex.EncodeToString(hash[:])
	return encodedstr
}

func CompareHashAndPassword(providedPassword string, existingHash string) bool {
	hashedPassword := HashStr(providedPassword)
	return hashedPassword == existingHash
}

func getNextHashSeq(text string) string {
	encodedstr := HashStr(text)
	return encodedstr[:6]
}
