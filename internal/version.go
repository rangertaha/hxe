/*
Copyright © 2025 Rangertaha <rangertaha@gmail.com>
*/
package internal

import "fmt"

const NAME = "hxe"
const VERSION = "0.0.1"
const COMPILED = "2025-06-18"
const BANNER = `   
  _   _          
 | | | |_  _____ 
 | |_| \ \/ / _ \
 |  _  |>  <  __/
 |_| |_/_/\_\___|
                         
 Host eXecution Engine
______________________________________________
AUTHOR:  Rangertaha
VERSION: v%s
DATE:    %s

`

func Banner() string {
	return fmt.Sprintf(BANNER, VERSION, COMPILED)
}

func PrintBanner() {
	fmt.Print(Banner())
}
