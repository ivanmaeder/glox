# Crafting Interpreters, in Go

The original book by Robert Nystrom which uses Java and C is [here](https://craftinginterpreters.com/).

## Notes

| Term           | Lexical grammar | Syntactic grammar |
| -------------- | --------------- | ----------------- |
| Alphabet       | Characters      | Lexeme/token      |
| String         | Lexeme/token    | Expression        |
| Implemented by | Scanner         | Parser            |

### Scanner

- Rules for how characters get grouped into tokens—was called a regular language
- Emits a flat sequence of tokens
- Not enough for arbitrarily-nested structures

```
s.scanTokens()
	s.scanToken() //finds the right token, sometimes by peeking ahead or e.g., looking for the closing string terminator
		s.addToken()
			s.tokens.append()
```

### Parser

- Which strings are valid and which aren't
- Rules = productions
  - Head
  - Body = a list of symbols
    - Two types of symbols
      - Terminal = a "letter" in the grammar (token/lexeme) (no more "moves" in the game)
      - Nonterminal = reference to another rule; play that rule and insert whatever it produces here
- Derivations = generate strings that are in the grammar
- "each rule needs to match expressions at that precedence level or higher"
- "Recursive descent is the simplest way to build a parser, and doesn’t require using complex parser generator tools like Yacc, Bison or ANTLR."
  - GCC, V8
  - "considered a top-down parser because it starts from the top or outermost grammar rule"
  - "In a top-down parser, you reach the lowest-precedence expressions first because they may in turn contain subexpressions of higher precedence."
- "Each method for parsing a grammar rule produces a syntax tree for that rule and returns it to the caller. When the body of the rule contains a nonterminal—a reference to another rule—we call that other rule’s method."

## Go

- In interfaces we define "method sets"
  - All types are `interface{}` (or `any`)
- Is-a relationship:

```
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person // this gives all Androids access to Talk()
	Model string
}

a := new(Android)
a.Talk() //an Android is a person
```

- Another option is for a struct to have the same function(s) as an interface (like duck typing but they call it structural typing because it happens at compile-time and duck typing is usually at runtime),

```
type Shape interface {
	Area() float32
}

type Rectangle struct {
	length float32
	width  float32
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

var shape Shape = &Rectangle{
	length: 10,
	width:  20,
}
```

- Unexported (private) is really package-level
- Packages are just 1-word (not paths)
- Naked returns
- Short var declarations (`:=`) are only available inside functions
- Const <name> (<type>) = <value>
- Ifs can include a short statement: `if <statement>; condition {...}`
  - Vars in this statement are in scope in the if/else block only
- Case statements can include logic:

```
switch {
case a <= 12:
	//...
case a > 12:
	//...
}
```

- A deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns
- Pointers
  - `var p *int`
  - `p //pointer to p`
  - `&p //value`
  - `\*p = 123 //set the value of p`
  - With a pointer to a struct, accessing fields "should" be: `(*s).field`
    - But the `s.field` shorthand is equivalent
    - Similarly, you can call a method with a value receiver with a pointer reference, or a method with a pointer receiver with a value reference (it will all work)
- This is a struct literal: `Vertex{val, val}`
- An array's length is part of its type!
  - "… when you assign or pass around an array value you will make a copy of its contents. (To avoid the copy you could pass a pointer to the array, but then that’s a pointer to an array, not an array.) One way to think about arrays is as a sort of struct but with indexed rather than named fields: a fixed-size composite value."
  - init: `[3]string{"a", "b", "c"}` (or `[...]string{"a", "b", "c"}`)
  - Slices are a dynamically-sized view into the elements of an array
    - Create a slice with `arr[n:m]`
    - Other slices with the same underlying array will see changes
    - This creates a slice (and array), `[]string{"a", "b", "c"}`
    - Slice capacity = number of elements in the underlying array counting from the first element in the slice
      - len(s), cap(s)
  - Use `make([]string, size[, capacity])` to create slices with dynamically-sized arrays
  - Use `slice = append(slice, ...values)` to add to a slice (this will increase the capacity of the array)
- `range` is for slices or maps
- Create maps with `make(map[string]Vertex)`
  - Access values with `m[index]`, delete with `delete(m, key)`, test with `elem, ok = m[key]`
  - Map literals:

```
m := map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```

- "A method is a function with a special receiver argument… a method is just a function with a receiver argument."
  - `func (r *Receiver) f() {}`
  - "You can only declare a method with a receiver whose type is defined in the same package as the method."
  - "Methods with pointer receivers can modify the value to which the receiver points."
  - Doing this also avoids copying values
  - "In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both."
- Type assertions: `variable.(T)` will return the underlying `T` value (or panic unless the two return values are used)
- Type switch:

```
func do(i interface{}) {
	switch v := i.(type) {

	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
```

- Implementing `String()` (from `Stringer`) is like implementing `Object.toString()`

```
type Stringer interface {
    String() string
}

//

func (p Person) String() string { //the receiver cannot be a pointer type
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
```

- The `error` interface:

```
type error interface {
    Error() string
}
```

- `go f(x, y, z)` runs `f()` as a new goroutine; the args are run before in the current thread
