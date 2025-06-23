/*
 * HXE - Host-based Process Execution Engine
 * Copyright (C) 2025 rangertaha@gmail.com
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

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
