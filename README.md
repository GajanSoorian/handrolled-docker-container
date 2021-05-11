# handrolled-docker-container

Inspired by "Building a container from scratch in Go" by Liz Rice


SysProcAttr flasgs explaination:

CLONE_NEWUTS: if this is set, then create the process in a new
UTS namespace, whose identifiers are initialized by
duplicating the identifiers from the UTS namespace of the
calling process.  If this flag is not set, then (as with
fork(2)) the process is created in the same UTS namespace
as the calling process.

Note: 
if you see the below error: 
panic: fork/exec /bin/bash: operation not permitted
goroutine 1 [running]:

Then, the user is probably does not have cap_sys_admin capapbilities. Follow the instruction here:
https://unix.stackexchange.com/questions/454708/how-do-you-add-cap-sys-admin-permissions-to-user-in-centos-7
Then enter:
su - <user>
Now confirm this with:
>capsh --print
Current: = cap_sys_admin+i

Then, sudo setcap cap_sys_admin+ie inside-docker-containers
