# logDMM

A tool to log data from a Digital Multi Meter (DMM) in Linux - and in the end also Mac and possibly Windows.

I have a Uni-T UT161D with logging ability over USB - however the software provided is for Windows only. Which won't help me one iota as I use Linux and Mac only for more 'serious' work. Windows is for games and VR fun, not productivity.

Ambition is to have this app be quite modular and support several makes, brands and models of DMMs, so built with this in mind. It will also - once I got it working properly in terminal - have a GUI, probably using ebiten engine for Go, though Fyne is an alternative I am looking at.

## Build

Will require libusb-1.0 or newer installed on the system. For Windows (or Mac) builds, please reference the guide from the gousb library README:

https://github.com/google/gousb

Linux should have this installed already, and if not is easily accesible from your distros package manager. Again reference the above link.

Non standard libraries used in this project:

- github.com/google/gousb

---

## Binary

In Linux it will have to be run using sudo. In Windows, you'll probably have to agree to it being not signed and potentially even run as admin. I'll know when getting it compiled in Windows. Which as it turns out is a bit of a hurdle that I'll have to figure out.

Reason for these elevated priveliges is without root/sudo, it won't grab the usb interface.

Mac and/or Windows may or may not be more forgiving.

## USB info

See file `DeviceInfo.md`

---

**Contact:**

location   | name/handle
-----------|---------
github:    | rDybing
twitter:   | @DybingRoy
Linked In: | Roy Dybing

---

## Releases

- Version format: [major release].[new feature(s)].[bugfix patch-version]
- Date format: yyyy-mm-dd

### v.0.0.1: 2023-11-28

- Added loading of (hardcoded) config file
- Added connecting to USB interface of DMM

---

## License: MIT

Full license text found in LICENCE file

## Copyright © 2023 Roy Dybing

---

ʕ◔ϖ◔ʔ