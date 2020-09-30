package digest

import (
	"net/url"

	"github.com/spikeekips/mitum/util/hint"
)

type Hal interface {
	Interface() interface{}
	Self() HalLink
	Links() map[string]HalLink
	AddLink(rel string, link HalLink) Hal
	SetSelf(HalLink) Hal
	SetInterface(interface{}) Hal
}

var (
	BaseHalType = hint.MustNewType(0xa0, 0x16, "mitum-currency-hal")
	BaseHalHint = hint.MustHint(BaseHalType, "0.0.1")
)

type BaseHal struct {
	ht    hint.Hint
	i     interface{}
	raw   []byte
	self  HalLink
	links map[string]HalLink
}

func NewBaseHal(i interface{}, self HalLink) BaseHal {
	return BaseHal{
		ht:    BaseHalHint,
		i:     i,
		self:  self,
		links: map[string]HalLink{},
	}
}

func (hal BaseHal) Hint() hint.Hint {
	return BaseHalHint
}

func (hal BaseHal) Interface() interface{} {
	return hal.i
}

func (hal BaseHal) RawInterface() []byte {
	return hal.raw
}

func (hal BaseHal) SetInterface(i interface{}) Hal {
	hal.i = i

	return hal
}

func (hal BaseHal) Links() map[string]HalLink {
	return hal.links
}

func (hal BaseHal) Self() HalLink {
	return hal.self
}

func (hal BaseHal) SetSelf(url HalLink) Hal {
	hal.self = url

	return hal
}

func (hal BaseHal) AddLink(rel string, link HalLink) Hal {
	hal.links[rel] = link

	return hal
}

type HalLink struct {
	href       string
	properties map[string]interface{}
}

func NewHalLink(href string, properties map[string]interface{}) HalLink {
	return HalLink{href: href, properties: properties}
}

func (hl HalLink) Href() string {
	return hl.href
}

func (hl HalLink) URL() (*url.URL, error) {
	return url.Parse(hl.href)
}

func (hl HalLink) Properties() map[string]interface{} {
	return hl.properties
}

func (hl HalLink) SetTemplated() HalLink {
	if hl.properties == nil {
		hl.properties = map[string]interface{}{}
	}

	hl.properties["templated"] = true

	return hl
}

func (hl HalLink) SetProperty(key string, value interface{}) HalLink {
	if hl.properties == nil {
		hl.properties = map[string]interface{}{}
	}

	hl.properties[key] = value

	return hl
}
