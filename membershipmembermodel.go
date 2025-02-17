package membership

import (
	"fmt"
	"strings"
	"time"
)

func demo(){
	fmt.Println("ds")
	strings.ToLower("dskjv ")
	t, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))
	fmt.Println("t",t)

	
}