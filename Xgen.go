package Xgen

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var tpl =  `package {{.Package}}
type Generic struct {

Value []{{.MyType}}
}


func From(source []{{.MyType}}) Generic{

return Generic{
Value:source,
}

}

func(c *Generic) Add(item {{.MyType}}){
c.Value=append(c.Value,item)
}

func(c *Generic) RemoveAt(index int){
c.Value[len(c.Value)-1],c.Value[index]=c.Value[index],c.Value[len(c.Value)-1]
c.Value=c.Value[:len(c.Value)-1]

}
func (c *Generic) AddRange(item []{{.MyType}}){
c.Value=append(c.Value,item...)
}

func (c Generic) First() (item {{.MyType}}){
return c.Value[0]
}

func (c Generic) Last() (item {{.MyType}}){
return c.Value[len(c.Value)-1]
}
func (c Generic) Contains(item {{.MyType}}) (result bool){
for i:=0;i< len(c.Value);i++{
if item==c.Value[i]{
result=true
return
}
}

return false
}

func (c *Generic) Reverse() {

	for i, j := 0, len(c.Value)-1; i < j; i, j = i+1, j-1 {
		c.Value[i], c.Value[j] = c.Value[j], c.Value[i]
	}

}

func (c Generic) Skip(val int) Generic{
	return Generic{
		Value:c.Value[val:],
	}
}
func (c Generic) Take(val int) Generic{
	return Generic{
		Value:c.Value[:val],
	}
}`


func F(vals []string) {

	tt := template.Must(template.New("xgeneric").Parse(tpl))

	for i:=0;i<len(vals);i++{
		dest:=strings.ToLower(vals[i])+"_xgen.go"
		file,err:=os.Create(dest)
		check(err)
		sals:=map[string]string{
			"MyType": vals[i],
			"Package": os.Getenv("GOPACKAGE"),
		}
		err=tt.Execute(file, sals)
		check(err)
		err=file.Close()
		check(err)
	}
}
func check(err error){
	if err!=nil{
		fmt.Println(err)
	}

}

