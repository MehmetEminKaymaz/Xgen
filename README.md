# Xgen
Create type-safe collections with Xgen for Go programming language (metaprogramming)

### Generate go code with go (Meta programming)


### Usage

```Go

type MyInt int

func main() {

	var typeNames =[]string{
		"MyInt",
		//other types...
	}
	Xgen.F(typeNames)//the function generate a go file that has generics for MyInt type


	
}

```

Or use go generate command


Reflection is usefull but some case you need to use metaprogramming.
