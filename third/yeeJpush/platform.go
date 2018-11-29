package yeeJpush

type PlatformType string

const (
	IOS      PlatformType = "ios"
	ANDROID  PlatformType = "android"
	WINPHONE PlatformType = "winphone"
)

type Platform struct {
	Object interface{}
}

func AllPlatform() *Platform {
	return &Platform{
		Object: "all",
	}
}

func (platform *Platform) Add(os PlatformType) {
	if platform.Object == nil {
		s := []string{string(os)}
		platform.Object = s
	} else {
		switch platform.Object.(type) {
		case []string:
			platform.Object = append(platform.Object.([]string), string(os))
		default:
		}

	}
}
