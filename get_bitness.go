/*usr/bin/env go run $0 $@; exit;*/

/* Return a value for Sec-CH-UA-Bitness.  We exepct to see either "64" or "32". */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// ---==[ Display the bitterness of the machine ]==--------------------------

  // These are the recognized values for GOARCH...
	//   $ go tool dist list | cut -d/ -f2 | sort | uniq
	switch runtime.GOARCH {
	case "386":
		fmt.Println("32") // https://en.wikipedia.org/wiki/X86
	case "amd64":
		fmt.Println("64") // https://en.wikipedia.org/wiki/X86-64
	case "arm":
		fmt.Println("32") // https://en.wikipedia.org/wiki/ARM_architecture_family
	case "arm64":
		fmt.Println("64") // https://en.wikipedia.org/wiki/AArch64
	case "loong64":
		fmt.Println("64") // https://en.wikipedia.org/wiki/Loongson
	case "mips":
		fmt.Println("32")
	case "mips64":
		fmt.Println("64") // https://en.wikipedia.org/wiki/MIPS_architecture
	case "mips64le":
		fmt.Println("64")
	case "mipsle":
		fmt.Println("32")
	case "ppc64":
		fmt.Println("64") // https://en.wikipedia.org/wiki/PowerPC
	case "ppc64le":
		fmt.Println("64")
	case "riscv64":
		fmt.Println("64") // https://en.wikipedia.org/wiki/RISC-V
	case "s390x":
		fmt.Println("64") // ??? https://en.wikipedia.org/wiki/IBM_Z
	case "wasm":
		fmt.Println("🐉") // here be dragons... https://github.com/golang/go/issues/63131
	}
	// NO DEFAULT CASE HERE
}

/// XXX FIXME TODO  I couldn't get any OpenWRT boxen to run binaries built for mips/mips64/mips64le/mipsle!!!

/*
The document at https://wicg.github.io/ua-client-hints defines the following headers...

  Sec-CH-UA                   => ???
  Sec-CH-UA-Arch              => "x86", "arm" or "ARM"
  Sec-CH-UA-Bitness           => "64", "32"
  Sec-CH-UA-Form-Factors      => "Desktop", "Automotive", "Mobile", "Tablet", "XR", "EInk", "Watch"
  Sec-CH-UA-Full-Version      => ???
  Sec-CH-UA-Full-Version-List => ???
  Sec-CH-UA-Mobile            => ?0, ?1
  Sec-CH-UA-Model             => ???
  Sec-CH-UA-Platform          => "Linux"/"Fuchsia", "Android", "iOS", "macOS", "Windows", "Unknown"
  Sec-CH-UA-Platform-Version  => "", "14.5", "11"
  Sec-CH-UA-WoW64             => ?0, ?1

(Tested on an M1 Macbook Pro and a few Linux machines;  No Unixes were attempted)
Here's some shell commands on Linux and macOS and what they display...

  $ [ $((0xffffffff)) -eq -1 ] && echo 32 || echo 64  # "64", "32"
  $ getconf LONG_BIT  # "64", "32"
  $ uname -m  # "x86_64", "arm64", etc.
  $ uname -s  # "Linux", "Darwin", etc.
  $ uname -o  # "GNU/Linux", "Darwin", etc.
  $ uname -p  # "unknown", "arm", etc.
  $ uname -i  # "unknown", illegal option, etc.
  $ arch  # doesn't work on some Linux distros

Things get a little more confusing when you're venturing farther into MIPS land or ARM land...

  - Raspberry Pi 3B a.k.a. Broadcom bcm2837 (arm7???) -> bcm27xx/bcm2710 -> aarch64_cortex-a53 (Whatever-Endian)
  - Raspberry Pi 4B a.k.a. Broadcom bcm2711 (arm8) -> bcm27xx/bcm2711 -> aarch64_cortex-a72 (Whatever-Endian)
  - Raspberry Pi 5B a.k.a. I don't have one to test with -> bcm27xx/bcm2712 -> aarch64_cortex-a76 (Whatever-Endian)
  - Ubiquiti EdgeRouter X a.k.a. MediaTek MT7530/MT7621AT ver1 eco3 a.k.a. mips32 1004 Kc V2.15 -> ramips/mt7621 -> mipsel_24kc (Little-Endian)
  - TP-Link EAP254 v3 US a.k.a. Atheros AR9561/AR8216/AR8236/AR8316 a.k.a. Qualcomm Atheros QCA9563/QCA9982/QCA9560/QCA9990 a.k.a. mips32 74Kc V5.0 -> ath79/generic -> mips_24kc (Big-Endian)
  - etc.
*/
