package hashutility

import (
	"crypto/md5"
	"encoding/hex"
)

func getNextHashSeq(text string) string {
	hash := md5.Sum([]byte(text))
	encodedstr := hex.EncodeToString(hash[:])
	return encodedstr[:6]
}
