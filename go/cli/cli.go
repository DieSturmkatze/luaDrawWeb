package cli

import (
	"fmt"
	"os"
	"strings"
)

const luaList = `
	lua function list:
		* draw:Println(text)     - prints text to console
		* draw:Rect(x, y, w, h, color)  - draws rectangle outline
		* draw:FillRect(x, y, w, h, color) - draws filled rectangle
		* draw:Line(x1, y1, x2, y2, color) - draws line
		* draw:Circle(x, y, radius, color) - draws circle outline
		* draw:ClearScreen() - manually clears screen (not needed with autoclear)
		* draw:KeyDown(key) - check if a key is pressed (follows https://www.toptal.com/developers/keycode)

	lua variable list:
		* Interval int - determines drawing fps (default: 30)
		* AutoClearScreen bool - sets auto clearing of screen (default: true)

	color list:
		* green
		* red
		* blue
		* black
		* white
		* yellow
		* purple
		* orange
`

const help = `
arguments: 
		foo.lua - file to execute
		help - displays this help
		genExample / ge - generates an example script.lua with documentation
` + luaList

const exampleLua = "--[[" + luaList + `]]

function init()
	draw.Interval = 30
	draw:Println("Hello from Lua")
end

local i = 0
local x = 0

function update()
	i = i + 1
	draw:Rect(100 + i , 100, 100, 100, "green")

	if draw:KeyDown(68) then
		x = x+2
	elseif draw:KeyDown(65) then
		x = x-2
	end

	draw:Circle(100+x, 100, 50, "blue")
end
`

func Setup() string {
	if len(os.Args) > 1 {
		arg1 := os.Args[1]

		if arg1 != "" && strings.HasSuffix(arg1, ".lua") {
			fmt.Println("Running given Script ", arg1)
			return arg1
		} else if arg1 == "help" {
			fmt.Println(help)
			os.Exit(0)
		} else if arg1 == "genExample" || arg1 == "ge" {
			fmt.Println("Generating example script.lua")
			if _, err := os.Stat("script.lua"); os.IsNotExist(err) {
				d1 := []byte(exampleLua)
				os.WriteFile("./script.lua", d1, 0644)
			} else {
				fmt.Println("script.lua already exists, aborting")
				os.Exit(1)
			}
		}
	}
	fmt.Println("Using default script ./script.lua (invalid or no argument given)")
	return "script.lua"
}
