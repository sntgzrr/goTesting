// Package say leads us say Hello.
package say

import "fmt"

// Say welcome to the person.
func Say(namePerson string) string {
	return fmt.Sprintf("Welcome %s", namePerson)
}
