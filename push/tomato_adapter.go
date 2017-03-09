package push

import "github.com/freeznet/tomato/types"

type tomatoPushAdapter struct {
	validPushTypes []string
}

func newTomatoPush() *tomatoPushAdapter {
	t := &tomatoPushAdapter{
		validPushTypes: []string{"ios", "android"},
	}
	return t
}

func (t *tomatoPushAdapter) send(data types.M, installations types.S, objectID string) []types.M {
	return []types.M{}
}

func (t *tomatoPushAdapter) getValidPushTypes() []string {
	return t.validPushTypes
}
