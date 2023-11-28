# DeviceInfo

## Uni-T UT161D

Basic info (using `lsusb`)

```
ID 1a86:e429 QinHeng Electronics WCH UART TO KB-MS_V1.6
```

From probing the USB devices (using `usbview`) I fished out this info regarding the bundled USB dongle:

```
WCH UART TO KB-MS_V1.6
Manufacturer: WWW.WCH.CN
Serial Number: 20196552E565
Speed: 12Mb/s (full)
Bus:   1
Address:   5
USB Version:  1.10
Device Class: 00(>ifc )
Device Subclass: 00
Device Protocol: 00
Maximum Default Endpoint Size: 8
Number of Configurations: 1
Vendor Id: 1a86
Product Id: e429
Revision Number: 36.01

Config Number: 1
	Number of Interfaces: 1
	Attributes: a0
	MaxPower Needed: 100mA

	Interface Number: 0
		Name: usbhid
		Alternate Number: 0
		Class: 03(HID  ) 
		Sub Class: 00
		Protocol: 00
		Number of Endpoints: 2

			Endpoint Address: 84
			Direction: in
			Attribute: 3
			Type: Int.
			Max Packet Size: 64
			Interval: 1ms

			Endpoint Address: 04
			Direction: out
			Attribute: 3
			Type: Int.
			Max Packet Size: 64
			Interval: 1ms
```

and this (using `usb-devices`):

```
T:  Bus=01 Lev=01 Prnt=01 Port=00 Cnt=01 Dev#=  5 Spd=12  MxCh= 0
D:  Ver= 1.10 Cls=00(>ifc ) Sub=00 Prot=00 MxPS= 8 #Cfgs=  1
P:  Vendor=1a86 ProdID=e429 Rev=36.01
S:  Manufacturer=WWW.WCH.CN
S:  Product=WCH UART TO KB-MS_V1.6
S:  SerialNumber=20196552E565
C:  #Ifs= 1 Cfg#= 1 Atr=a0 MxPwr=100mA
I:  If#= 0 Alt= 0 #EPs= 2 Cls=03(HID  ) Sub=00 Prot=00 Driver=usbhid
E:  Ad=04(O) Atr=03(Int.) MxPS=  64 Ivl=1ms
E:  Ad=84(I) Atr=03(Int.) MxPS=  64 Ivl=1ms
```
