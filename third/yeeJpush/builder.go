package yeeJpush

// Base Builder
type Builder struct {
	Platform interface{} `json:"platform"`
	Audience interface{} `json:"audience"`
	Options  Options     `json:"options"`
}

// MessageBuilder
type MessageBuilder struct {
	Builder
	Message interface{} `json:"message"`
}

// NoticeBuilder
type NoticeBuilder struct {
	Builder
	Notification map[string]interface{} `json:"notification"`
}

// MessageAndNotice
type MessageAndNoticeBuilder struct {
	Builder
	Notification map[string]interface{} `json:"notification"`
	Message      interface{}            `json:"message"`
}

/*---------------------MessageBuilder --------------------*/

func NewMessageBuilder() *MessageBuilder {
	mb := &MessageBuilder{}
	mb.Options = NewOptions()
	return mb
}

func (messageBuilder *MessageBuilder) SetMessage(m *Message) {
	messageBuilder.Message = m
}

func (messageBuilder *MessageBuilder) SetPlatform(pf *Platform) {
	messageBuilder.Platform = pf.Object
}

func (messageBuilder *MessageBuilder) SetAudience(ad *Audience) {
	messageBuilder.Audience = ad.Object
}

func (messageBuilder *MessageBuilder) SetOptions(o Options) {
	messageBuilder.Options = o
}

/*---------------------NoticeBuilder----------------------*/

func NewNoticeBuilder() *NoticeBuilder {
	nb := &NoticeBuilder{}
	nb.Options = NewOptions()
	return nb
}

func (noticeBuilder *NoticeBuilder) SetPlatform(pf *Platform) {
	noticeBuilder.Platform = pf.Object
}

func (noticeBuilder *NoticeBuilder) SetAudience(ad *Audience) {
	noticeBuilder.Audience = ad.Object
}

func (noticeBuilder *NoticeBuilder) SetOptions(o Options) {
	noticeBuilder.Options = o
}
func (noticeBuilder *NoticeBuilder) ClearNotice() {
	noticeBuilder.Notification = nil
}

// 可以为每类Notice设置一个
func (noticeBuilder *NoticeBuilder) SetNotice(o interface{}) {
	if noticeBuilder.Notification == nil {
		noticeBuilder.Notification = make(map[string]interface{})
	}
	switch o.(type) {
	case *NoticeAndroid:
		noticeBuilder.Notification[string(ANDROID)] = o
	case *NoticeWinphone:
		noticeBuilder.Notification[string(WINPHONE)] = o
	case *NoticeIos:
		noticeBuilder.Notification[string(IOS)] = o
	case *NoticeSimple:
		noticeBuilder.Notification["alert"] = o.(*NoticeSimple).Alert
	}
}

/*------------------MessageAndNoticeBuilder------------------*/

func NewMessageAndNoticeBuilder() *MessageAndNoticeBuilder {
	mnb := &MessageAndNoticeBuilder{}
	mnb.Options = NewOptions()
	return mnb
}

func (messageNoticeBuilder *MessageAndNoticeBuilder) SetPlatform(pf *Platform) {
	messageNoticeBuilder.Platform = pf.Object
}

func (messageNoticeBuilder *MessageAndNoticeBuilder) SetAudience(ad *Audience) {
	messageNoticeBuilder.Audience = ad.Object
}

func (messageNoticeBuilder *MessageAndNoticeBuilder) SetOptions(o Options) {
	messageNoticeBuilder.Options = o
}

func (messageNoticeBuilder *MessageAndNoticeBuilder) SetMessage(m *Message) {
	messageNoticeBuilder.Message = m
}
func (messageNoticeBuilder *MessageAndNoticeBuilder) ClearNotice() {
	messageNoticeBuilder.Notification = nil
}

// 可以为每类Notice设置一个
func (messageNoticeBuilder *MessageAndNoticeBuilder) SetNotice(o interface{}) {
	if messageNoticeBuilder.Notification == nil {
		messageNoticeBuilder.Notification = make(map[string]interface{})
	}
	switch o.(type) {
	case NoticeAndroid:
		messageNoticeBuilder.Notification[string(ANDROID)] = o
	case NoticeWinphone:
		messageNoticeBuilder.Notification[string(WINPHONE)] = o
	case NoticeIos:
		messageNoticeBuilder.Notification[string(IOS)] = o
	case NoticeSimple:
		messageNoticeBuilder.Notification["alert"] = o.(NoticeSimple).Alert
	}

}
