package main 

import{
	"fmt"
	"connect"
}

func main(){
	var c connect
	c.open("pop.nyx.net","110")
	c.close()
	
}