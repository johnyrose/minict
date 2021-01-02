# Minict

Minict is a small and minimal container runtime written in Go. It was made mainly for learning purposes and is intended to be as simple as possible. 

It's main intention is to be easily understandable to anyone who wishes to read it's code and see what goes into running containers in famous runtimes such as Containerd and full container-running platforms like Docker.

Minict runs OCI standard images and supports pulling images from existing registries. 

## Prerequisites
 * The `gpgme-devel` package must be installed on your system.
   * Run `sudo dnf install gpgme-devel` on RHEL-based distros (RHEL, CentOS, Fedora, etc.)
   * Run `sudo apt install libgpgme-dev` on Debian-based distros (Debian, Ubuntu, etc.)
 * Have `golang` and `git` installed.

## Building & Running
 * Clone this repository into your `$HOME/go/src` directory and run `cd $HOME/go/src/minict`
 * Run the `go get` command.
 * Run the `go build` command.
 * You should now have a `minict` executable in your directory. Run `chmod a+x minict` and it can now be used.
 * To use it from anywhere easily, move the `minict` executable to a directory that is in yout `PATH` variable. One such directory should be `/usr/bin`.