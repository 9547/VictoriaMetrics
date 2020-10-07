package auth

import (
	"fmt"
	"strconv"
	"strings"
)

// Token contains settings for request processing
type Token struct {
	ProjectID uint32
	AccountID uint32
}

func (t *Token) String() string {
	return fmt.Sprintf("%d:%d", t.AccountID, t.ProjectID)
}

// NewToken returns new Token for the given authToken
func NewToken(authToken string) (*Token, error) {
	var at Token
	// fast path for empty character
	if authToken == "" {
		return &at, nil
	}
	tmp := strings.Split(authToken, ":")
	if len(tmp) > 2 {
		return nil, fmt.Errorf("unexpected number of items in authToken %q; got %d; want 1 or 2", authToken, len(tmp))
	}
	accountID, err := strconv.Atoi(tmp[0])
	if err != nil {
		return nil, fmt.Errorf("cannot parse accountID from %q: %w", tmp[0], err)
	}
	at.AccountID = uint32(accountID)
	if len(tmp) > 1 {
		projectID, err := strconv.Atoi(tmp[1])
		if err != nil {
			return nil, fmt.Errorf("cannot parse projectID from %q: %w", tmp[1], err)
		}
		at.ProjectID = uint32(projectID)
	}
	return &at, nil
}
