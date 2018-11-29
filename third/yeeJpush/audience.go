package yeeJpush

const (
	TAG     = "tag"
	TAG_AND = "tag_and"
	ALIAS   = "alias"
	ID      = "registration_id"
)

type Audience struct {
	Object interface{}
}

func AllAudience() *Audience {
	return &Audience{
		Object: "all",
	}
}

func (audience *Audience) SetID(Object []string) {
	audience.set(ID, Object)
}

func (audience *Audience) SetTag(Object []string) {
	audience.set(TAG, Object)
}

func (audience *Audience) SetTagAnd(Object []string) {
	audience.set(TAG_AND, Object)
}

func (audience *Audience) SetAlias(Object []string) {
	audience.set(ALIAS, Object)
}

func (audience *Audience) set(key string, Object []string) {
	if audience.Object == nil {
		audience.Object = map[string][]string{key: Object}
	} else {
		switch audience.Object.(type) {
		case map[string][]string:
			audience.Object.(map[string][]string)[key] = Object
		default:
		}
	}
}
