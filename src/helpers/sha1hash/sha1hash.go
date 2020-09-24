package sha1hash

import (
    "crypto/sha1"
    "encoding/hex"
)

func Make(input string) string {
    // To make input string to byte type
    aStringToHash := []byte(input)
    // To encrypt the input byte with sha1 
    sha1Bytes := sha1.Sum(aStringToHash)
    // To convert hexed bytes to string
    result := hex.EncodeToString(sha1Bytes[:])  
    
    return result    
}