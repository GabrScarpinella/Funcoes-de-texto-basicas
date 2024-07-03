package main

import "fmt"

func inverso(pal string){
	for i:=len(pal)-1; i>=0; i--{
		fmt.Print(string(pal[i]));
	}
}
func palindromo(pal string) bool{
	for i:=0; i<len(pal); i++{
		if pal[i]!=pal[(len(pal)-i-1)]{
			return false;
		}
	}
	return true;
}
func vogais(pal string)(int, int, int, int, int, int){//Quantidade de vogais, ases, es, is, os e us
	var v int=0;
	var a int=0;
	var e int=0;
	var i int=0;
	var o int=0;
	var u int=0;

	for j:=0; j<len(pal); j++{
		if string(pal[j])=="a"{
			a++;
		} else if string(pal[j])=="e"{
			e++;
		} else if string(pal[j])=="i"{
			i++;
		} else if string(pal[j])=="o"{
			o++;
		} else if string(pal[j])=="u"{
			u++;
		}
	}
	v=a+e+i+o+u;
	return v, a, e, i, o, u;
}

func main(){
	var t string="arar";
	inverso(t);
	fmt.Println(palindromo(t));
	fmt.Println(vogais(t));
}