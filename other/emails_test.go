package other

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqueEmails(t *testing.T) {
	{
		emails := []string{
			"test.email@gmail.com",
			"test.email@mail.ru",
			"test.email....@mail.ru",
			"test+body@gmail.com",
			"test+well@gmail.com",
			"test+fake@g.mail.com",
			"test+fake@g.mail+com",
		}
		require.Equal(t, 5, numUniqueEmails(emails))
	}
	{
		emails := []string{
			"test.email+alex@leetcode.com",
			"test.e.mail+bob.cathy@leetcode.com",
			"testemail+david@lee.tcode.com",
		}
		require.Equal(t, 2, numUniqueEmails(emails))
	}
	{
		emails := []string{
			"a@leetcode.com", "b@leetcode.com", "c@leetcode.com",
		}
		require.Equal(t, 3, numUniqueEmails(emails))
	}
}
