package grandpa

import "fmt"

// SearchKey Very dumb search TODO improve
// also TODO test
func SearchKey(key Key, changes []PendingChange) (int, error) {
	for i, change := range changes {
		changeKey := Key{
			effectiveNumber:   change.EffectiveNumber(),
			signalBlockNumber: change.canonHeight,
		}
		if key.Equals(changeKey) {
			return i, nil
		}
	}
	return 0, fmt.Errorf("key not found")
}
