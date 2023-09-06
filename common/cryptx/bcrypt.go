package cryptx

import "golang.org/x/crypto/bcrypt"

func BcryptHash(rawPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func IsBcryptHashed(str string) bool {
	// cryptx 加密后的长度等于 60
	return len(str) == 60
}

func BcyptCheck(rawPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err == nil
}
