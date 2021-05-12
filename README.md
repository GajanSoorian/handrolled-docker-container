# handrolled-docker-container

Using Linux namespaces to build our own isolated "continaer". Inspired by "Building a container from scratch in Go" by Liz Rice


SysProcAttr flasgs explaination:

CLONE_NEWUTS: if this is set, then create the process in a new
              UTS namespace, whose identifiers are initialized by
              duplicating the identifiers from the UTS namespace of the
              calling process.  If this flag is not set, then (as with
              fork(2)) the process is created in the same UTS namespace
              as the calling process.

CLONE_NEWPID: If CLONE_NEWPID is set, then create the process in a new
              PID namespace.  If this flag is not set, then (as with
              fork(2)) the process is created in the same PID namespace
              as the calling process.


Fork and exec explaination:

fork : create a new process and returns the process id of this child process.

Mounting a linux file system for our container:


1) mkinitramfs -o initrd.img
   unmkinitramfs -v initrd.img <maybe in a temp folder>

2)
sudo dd if=/dev/zero of=rootfs.img bs=1024 count=500000  (500 MB)
sudo mkfs.ext4 -F rootfs.img

Mount the fs:
sudo mount rootfs.img /mnt -t ext4

Now copy the contents of initramfs
sudo cp -rf tmp/main/* /mnt/containerFS<or any mount path>
sudo cp -rf /bin/bash /mnt/bin/   #for good measure since intird wont have it.
sudo mkdir /mnt/sys /mnt/proc /mnt/dev



Troubleshooting: 
if you see the below error: 
panic: fork/exec /bin/bash: operation not permitted
goroutine 1 [running]:

Build the code and run it with sudo since this is not a rootless container.


Take home:
Namespaces are what you can see: UNIX Timesharing System, Process IDs, File system (mount points), Users, IPC, Networking
Control groups are what you can use: CPU, Memory, Disk I/O, N/W, Disk Permissions. Path on system: /sys/fs/cgroup/

Bonus: Image layers(created with docker build) are the file systems the container sees, with additional things like ENV variables!