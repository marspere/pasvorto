package pasvorto

import "github.com/nbutton23/zxcvbn-go"

// Score used to evaluate the password strength.
func Score(pwd string) int {
	return zxcvbn.PasswordStrength(pwd, nil).Score
}

func (pwd *Password) scoring() {
	pwd.Score = zxcvbn.PasswordStrength(pwd.Value, nil).Score
}
