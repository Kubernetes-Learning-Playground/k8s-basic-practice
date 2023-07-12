package main

import (
	"fmt"
	"github.com/vishvananda/netlink"
	"log"
)

// Netlink: linux提供的通信方式，用于内核与用户态之前的通信
/*
	1、/proc 文件
  	譬如  cat /proc/net/route   (对照 route -n)

	2、 ioctl  (ifconfig)

	3、netlink (ip命令  ip addr, ip link )

*/

func main() {

	//getLink("eth0")

	createBridge("mybr", "10.16.0.1/16")

}

// getLink 获取设备对象
func getLink(name string) {

	link, err := netlink.LinkByName(name)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(link.Type())
	addr, err := netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(addr)
}

// createBridge 创建网桥
func createBridge(name, bridgeAddr string) {
	br := &netlink.Bridge{
		LinkAttrs: netlink.LinkAttrs{Name: name},
	}

	err := netlink.LinkAdd(br)
	if err != nil {
		log.Fatalln(err)
	}

	addr, _ := netlink.ParseAddr(bridgeAddr)

	// 加入地址
	err = netlink.AddrAdd(br, addr)
	if err != nil {
		log.Fatalln(err)
	}

	// 启动
	err = netlink.LinkSetUp(br)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
[root@VM-0-16-centos netlink_device]# ip link set mybr up     # 启动网桥设备
[root@VM-0-16-centos netlink_device]# ip link set mybr down   # 关闭网桥设备
[root@VM-0-16-centos ~]# ip link del mybr	# 删除设备
[root@VM-0-16-centos ~]# ifconfig
cni0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1450
		inet 10.244.0.1  netmask 255.255.255.0  broadcast 10.244.0.255
		inet6 fe80::94dc:38ff:fe78:3835  prefixlen 64  scopeid 0x20<link>
		ether 96:dc:38:78:38:35  txqueuelen 1000  (Ethernet)
		RX packets 817728698  bytes 1373876343337 (1.2 TiB)
		RX errors 0  dropped 0  overruns 0  frame 0
		TX packets 800353446  bytes 292087351888 (272.0 GiB)
		TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

docker0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
		inet 172.17.0.1  netmask 255.255.0.0  broadcast 172.17.255.255
		inet6 fe80::42:93ff:fe82:d7c4  prefixlen 64  scopeid 0x20<link>
		ether 02:42:93:82:d7:c4  txqueuelen 0  (Ethernet)
		RX packets 375643402  bytes 155234949132 (144.5 GiB)
		RX errors 0  dropped 0  overruns 0  frame 0
		TX packets 396627618  bytes 1670146268173 (1.5 TiB)
		TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
		inet 10.0.0.16  netmask 255.255.252.0  broadcast 10.0.3.255
		inet6 fe80::5054:ff:fe84:3e8c  prefixlen 64  scopeid 0x20<link>
		ether 52:54:00:84:3e:8c  txqueuelen 1000  (Ethernet)
		RX packets 110006503  bytes 18352095000 (17.0 GiB)
		RX errors 0  dropped 0  overruns 0  frame 0
		TX packets 142199644  bytes 19654729529 (18.3 GiB)
		TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

flannel.1: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1450
		inet 10.244.0.0  netmask 255.255.255.255  broadcast 0.0.0.0
		inet6 fe80::e83d:7bff:fe70:dc33  prefixlen 64  scopeid 0x20<link>
		ether ea:3d:7b:70:dc:33  txqueuelen 0  (Ethernet)
		RX packets 0  bytes 0 (0.0 B)
		RX errors 0  dropped 0  overruns 0  frame 0
		TX packets 0  bytes 0 (0.0 B)
		TX errors 0  dropped 8 overruns 0  carrier 0  collisions 0
 */
