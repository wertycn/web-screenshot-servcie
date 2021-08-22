package Service

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/device"
)

var ctx context.Context

var deviceMap = make(map[string]chromedp.Device)

func RegisterContext(c context.Context) {
	ctx = c
	initDeviceMap()
}

func GetChromeContext() context.Context {
	return ctx
}

func initDeviceMap() {

	// BlackberryPlayBook is the "Blackberry PlayBook" device.
	deviceMap["BlackberryPlayBook"] = device.BlackberryPlayBook

	// BlackberryPlayBooklandscape is the "Blackberry PlayBook landscape" device.
	deviceMap["BlackberryPlayBooklandscape"] = device.BlackberryPlayBooklandscape

	// BlackBerryZ30 is the "BlackBerry Z30" device.
	deviceMap["BlackBerryZ30"] = device.BlackBerryZ30

	// BlackBerryZ30landscape is the "BlackBerry Z30 landscape" device.
	deviceMap["BlackBerryZ30landscape"] = device.BlackBerryZ30landscape

	// GalaxyNote3 is the "Galaxy Note 3" device.
	deviceMap["GalaxyNote3"] = device.GalaxyNote3

	// GalaxyNote3landscape is the "Galaxy Note 3 landscape" device.
	deviceMap["GalaxyNote3landscape"] = device.GalaxyNote3landscape

	// GalaxyNoteII is the "Galaxy Note II" device.
	deviceMap["GalaxyNoteII"] = device.GalaxyNoteII

	// GalaxyNoteIIlandscape is the "Galaxy Note II landscape" device.
	deviceMap["GalaxyNoteIIlandscape"] = device.GalaxyNoteIIlandscape

	// GalaxySIII is the "Galaxy S III" device.
	deviceMap["GalaxySIII"] = device.GalaxySIII

	// GalaxySIIIlandscape is the "Galaxy S III landscape" device.
	deviceMap["GalaxySIIIlandscape"] = device.GalaxySIIIlandscape

	// GalaxyS5 is the "Galaxy S5" device.
	deviceMap["GalaxyS5"] = device.GalaxyS5

	// GalaxyS5landscape is the "Galaxy S5 landscape" device.
	deviceMap["GalaxyS5landscape"] = device.GalaxyS5landscape

	// IPad is the "iPad" device.
	deviceMap["IPad"] = device.IPad

	// IPadlandscape is the "iPad landscape" device.
	deviceMap["IPadlandscape"] = device.IPadlandscape

	// IPadMini is the "iPad Mini" device.
	deviceMap["IPadMini"] = device.IPadMini

	// IPadMinilandscape is the "iPad Mini landscape" device.
	deviceMap["IPadMinilandscape"] = device.IPadMinilandscape

	// IPadPro is the "iPad Pro" device.
	deviceMap["IPadPro"] = device.IPadPro

	// IPadProlandscape is the "iPad Pro landscape" device.
	deviceMap["IPadProlandscape"] = device.IPadProlandscape

	// IPhone4 is the "iPhone 4" device.
	deviceMap["IPhone4"] = device.IPhone4

	// IPhone4landscape is the "iPhone 4 landscape" device.
	deviceMap["IPhone4landscape"] = device.IPhone4landscape

	// IPhone5 is the "iPhone 5" device.
	deviceMap["IPhone5"] = device.IPhone5

	// IPhone5landscape is the "iPhone 5 landscape" device.
	deviceMap["IPhone5landscape"] = device.IPhone5landscape

	// IPhone6 is the "iPhone 6" device.
	deviceMap["IPhone6"] = device.IPhone6

	// IPhone6landscape is the "iPhone 6 landscape" device.
	deviceMap["IPhone6landscape"] = device.IPhone6landscape

	// IPhone6Plus is the "iPhone 6 Plus" device.
	deviceMap["IPhone6Plus"] = device.IPhone6Plus

	// IPhone6Pluslandscape is the "iPhone 6 Plus landscape" device.
	deviceMap["IPhone6Pluslandscape"] = device.IPhone6Pluslandscape

	// IPhone7 is the "iPhone 7" device.
	deviceMap["IPhone7"] = device.IPhone7

	// IPhone7landscape is the "iPhone 7 landscape" device.
	deviceMap["IPhone7landscape"] = device.IPhone7landscape

	// IPhone7Plus is the "iPhone 7 Plus" device.
	deviceMap["IPhone7Plus"] = device.IPhone7Plus

	// IPhone7Pluslandscape is the "iPhone 7 Plus landscape" device.
	deviceMap["IPhone7Pluslandscape"] = device.IPhone7Pluslandscape

	// IPhone8 is the "iPhone 8" device.
	deviceMap["IPhone8"] = device.IPhone8

	// IPhone8landscape is the "iPhone 8 landscape" device.
	deviceMap["IPhone8landscape"] = device.IPhone8landscape

	// IPhone8Plus is the "iPhone 8 Plus" device.
	deviceMap["IPhone8Plus"] = device.IPhone8Plus

	// IPhone8Pluslandscape is the "iPhone 8 Plus landscape" device.
	deviceMap["IPhone8Pluslandscape"] = device.IPhone8Pluslandscape

	// IPhoneSE is the "iPhone SE" device.
	deviceMap["IPhoneSE"] = device.IPhoneSE

	// IPhoneSElandscape is the "iPhone SE landscape" device.
	deviceMap["IPhoneSElandscape"] = device.IPhoneSElandscape

	// IPhoneX is the "iPhone X" device.
	deviceMap["IPhoneX"] = device.IPhoneX

	// IPhoneXlandscape is the "iPhone X landscape" device.
	deviceMap["IPhoneXlandscape"] = device.IPhoneXlandscape

	// IPhoneXR is the "iPhone XR" device.
	deviceMap["IPhoneXR"] = device.IPhoneXR

	// IPhoneXRlandscape is the "iPhone XR landscape" device.
	deviceMap["IPhoneXRlandscape"] = device.IPhoneXRlandscape

	// IPhone11 is the "iPhone 11" device.
	deviceMap["IPhone11"] = device.IPhone11

	// IPhone11landscape is the "iPhone 11 landscape" device.
	deviceMap["IPhone11landscape"] = device.IPhone11landscape

	// IPhone11Pro is the "iPhone 11 Pro" device.
	deviceMap["IPhone11Pro"] = device.IPhone11Pro

	// IPhone11Prolandscape is the "iPhone 11 Pro landscape" device.
	deviceMap["IPhone11Prolandscape"] = device.IPhone11Prolandscape

	// IPhone11ProMax is the "iPhone 11 Pro Max" device.
	deviceMap["IPhone11ProMax"] = device.IPhone11ProMax

	// IPhone11ProMaxlandscape is the "iPhone 11 Pro Max landscape" device.
	deviceMap["IPhone11ProMaxlandscape"] = device.IPhone11ProMaxlandscape

	// JioPhone2 is the "JioPhone 2" device.
	deviceMap["JioPhone2"] = device.JioPhone2

	// JioPhone2landscape is the "JioPhone 2 landscape" device.
	deviceMap["JioPhone2landscape"] = device.JioPhone2landscape

	// KindleFireHDX is the "Kindle Fire HDX" device.
	deviceMap["KindleFireHDX"] = device.KindleFireHDX

	// KindleFireHDXlandscape is the "Kindle Fire HDX landscape" device.
	deviceMap["KindleFireHDXlandscape"] = device.KindleFireHDXlandscape

	// LGOptimusL70 is the "LG Optimus L70" device.
	deviceMap["LGOptimusL70"] = device.LGOptimusL70

	// LGOptimusL70landscape is the "LG Optimus L70 landscape" device.
	deviceMap["LGOptimusL70landscape"] = device.LGOptimusL70landscape

	// MicrosoftLumia550 is the "Microsoft Lumia 550" device.
	deviceMap["MicrosoftLumia550"] = device.MicrosoftLumia550

	// MicrosoftLumia950 is the "Microsoft Lumia 950" device.
	deviceMap["MicrosoftLumia950"] = device.MicrosoftLumia950

	// MicrosoftLumia950landscape is the "Microsoft Lumia 950 landscape" device.
	deviceMap["MicrosoftLumia950landscape"] = device.MicrosoftLumia950landscape

	// Nexus10 is the "Nexus 10" device.
	deviceMap["Nexus10"] = device.Nexus10

	// Nexus10landscape is the "Nexus 10 landscape" device.
	deviceMap["Nexus10landscape"] = device.Nexus10landscape

	// Nexus4 is the "Nexus 4" device.
	deviceMap["Nexus4"] = device.Nexus4

	// Nexus4landscape is the "Nexus 4 landscape" device.
	deviceMap["Nexus4landscape"] = device.Nexus4landscape

	// Nexus5 is the "Nexus 5" device.
	deviceMap["Nexus5"] = device.Nexus5

	// Nexus5landscape is the "Nexus 5 landscape" device.
	deviceMap["Nexus5landscape"] = device.Nexus5landscape

	// Nexus5X is the "Nexus 5X" device.
	deviceMap["Nexus5X"] = device.Nexus5X

	// Nexus5Xlandscape is the "Nexus 5X landscape" device.
	deviceMap["Nexus5Xlandscape"] = device.Nexus5Xlandscape

	// Nexus6 is the "Nexus 6" device.
	deviceMap["Nexus6"] = device.Nexus6

	// Nexus6landscape is the "Nexus 6 landscape" device.
	deviceMap["Nexus6landscape"] = device.Nexus6landscape

	// Nexus6P is the "Nexus 6P" device.
	deviceMap["Nexus6P"] = device.Nexus6P

	// Nexus6Plandscape is the "Nexus 6P landscape" device.
	deviceMap["Nexus6Plandscape"] = device.Nexus6Plandscape

	// Nexus7 is the "Nexus 7" device.
	deviceMap["Nexus7"] = device.Nexus7

	// Nexus7landscape is the "Nexus 7 landscape" device.
	deviceMap["Nexus7landscape"] = device.Nexus7landscape

	// NokiaLumia520 is the "Nokia Lumia 520" device.
	deviceMap["NokiaLumia520"] = device.NokiaLumia520

	// NokiaLumia520landscape is the "Nokia Lumia 520 landscape" device.
	deviceMap["NokiaLumia520landscape"] = device.NokiaLumia520landscape

	// NokiaN9 is the "Nokia N9" device.
	deviceMap["NokiaN9"] = device.NokiaN9

	// NokiaN9landscape is the "Nokia N9 landscape" device.
	deviceMap["NokiaN9landscape"] = device.NokiaN9landscape

	// Pixel2 is the "Pixel 2" device.
	deviceMap["Pixel2"] = device.Pixel2

	// Pixel2landscape is the "Pixel 2 landscape" device.
	deviceMap["Pixel2landscape"] = device.Pixel2landscape

	// Pixel2XL is the "Pixel 2 XL" device.
	deviceMap["Pixel2XL"] = device.Pixel2XL

	// Pixel2XLlandscape is the "Pixel 2 XL landscape" device.
	deviceMap["Pixel2XLlandscape"] = device.Pixel2XLlandscape

}

func GetDevice(name string) chromedp.Device {
	if _, ok := deviceMap[name]; ok {
		return deviceMap[name]
	}
	return device.IPhone8
}

func GetDeviceList() []string {
	var deviceList  []string
	for k, _ := range deviceMap {
		deviceList = append(deviceList, k)
	}
	return deviceList
}
