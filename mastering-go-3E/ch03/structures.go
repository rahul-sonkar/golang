package main

import "fmt"

type Empty struct {
	name string 
	surname string
	year int
}

// Initialized by Go
	// return Structer instant
func zeroS() Empty {
	return Empty{}
}
	// return pointer to a structer
func zeroPtoS() *Empty {
	return &Empty{}
}

// Innitialized by User
func initS(n,s string,y int) Empty {
	if y < 2000 {
		return Empty{name: n,surname: s,year: 2000}
	}
	return Empty{name: n,surname: s,year: y}
}

func initPtoS(n,s string,y int) *Empty {
	return &Empty{name: n,surname: s,year: y}
}


func main()  {
	s1 := zeroS()
	p1 := zeroPtoS()
	fmt.Println("s1:",s1," p1:",*p1)

	s2 := initS("rahul","sonkar",2002)
	p2 := initPtoS("name","surname",2000)
	fmt.Println("s2:",s2," p2:",*p2)

	fmt.Println("Years:",s1.year, s2.year, p1.year, p2.year)
	
	pS := new(Empty)
	fmt.Println("pS:",pS)
	fmt.Println("pS:",*pS)

}