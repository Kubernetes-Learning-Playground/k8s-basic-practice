package main

/*
	进入：
	[root@VM-0-16-centos ~]# cd /sys/fs/cgroup/
	[root@VM-0-16-centos cgroup]# pwd
	/sys/fs/cgroup
	[root@VM-0-16-centos cgroup]# ls
	blkio  cpu  cpuacct  cpu,cpuacct  cpuset  devices  freezer  hugetlb  memory  net_cls  net_cls,net_prio  net_prio  perf_event  pids  systemd

	[root@VM-0-16-centos cpu]# ls
	cgroup.clone_children  cgroup.procs          cpuacct.stat   cpuacct.usage_percpu  cpu.cfs_quota_us  cpu.rt_runtime_us  cpu.stat        myapp              release_agent  tasks
	cgroup.event_control   cgroup.sane_behavior  cpuacct.usage  cpu.cfs_period_us     cpu.rt_period_us  cpu.shares         kubepods.slice  notify_on_release  system.slice   user.slice
	[root@VM-0-16-centos cpu]# cd myapp/
	[root@VM-0-16-centos myapp]# ls
	cgroup.clone_children  cgroup.event_control  cgroup.procs  cpuacct.stat  cpuacct.usage  cpuacct.usage_percpu  cpu.cfs_period_us  cpu.cfs_quota_us  cpu.rt_period_us  cpu.rt_runtime_us  cpu.shares  cpu.stat  notify_on_release  tasks

	echo 10000 > cpu.cfs_period_us
	echo 1000 > cpu.cfs_quota_us

	cgexec -g cpu:myapp  ./myapp

*/

func main() {
	i := 1
	for {
		i++
	}
}
