package main

import (
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
	"log"
)

/*
	测试手动搭建容器veth。
	docker 默认使用bridge模式，会自动为容器端与宿主机端创建veth设备，并连接到docker0网桥上。
	使用此程序需要关闭默认网桥，使用none模式，即：不会创建出任何网络设备或ip等
	ex: docker run -d --name test --net=none nginx:1.18-alpine

	# 只有回还设备
	[root@VM-0-16-centos netlink_device]# docker exec -it test sh
	/ # ip a
	1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1000
		link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
		inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever

	docker inspect test | grep Pid   # 查看pid进程号

	ip netns 默认目录：/var/run/netns 下
	docker 需要在 /proc/[pid]/ns/net 下找到

/ #
 */

func main() {
	// veth设备对
	vethpeer := &netlink.Veth{
		LinkAttrs: netlink.LinkAttrs{Name: "myveth-host"},
		PeerName:  "myveth-docker", // 对端（容器端）
	}
	// ip link add
	err := netlink.LinkAdd(vethpeer)
	if err != nil {
		log.Fatalln(err)
	}
	ns, err := netns.GetFromPath("/proc/31916/ns/net")
	if err != nil {
		log.Fatalln(err)
	}
	defer ns.Close()
	myveth_docker, err := netlink.LinkByName("myveth-docker")
	if err != nil {
		log.Fatalln(err)
	}

	// 把其中一端放入容器端namespace
	// ip link set xxx netns dockerns
	err = netlink.LinkSetNsFd(myveth_docker, int(ns))
	if err != nil {
		log.Fatalln(err)
	}


	// 设置当前网络命名空间，进入当前命名空间执行命令
	err = netns.Set(ns)
	if err != nil {
		log.Fatalln(err)
	}
	// 下面就是容器端内的操作


	myveth_docker, err = netlink.LinkByName("myveth-docker")
	if err != nil {
		log.Fatalln(err)
	}

	addr, _ := netlink.ParseAddr("10.18.0.10/16")

	// 设置IP地址
	err = netlink.AddrAdd(myveth_docker, addr)
	if err != nil {
		log.Fatalln(err)
	}

	// 修改容器端设备名称
	_ = netlink.LinkSetName(myveth_docker, "eth0")
	err = netlink.LinkSetUp(myveth_docker)
	if err != nil {
		log.Fatalln(err)
	}

}
