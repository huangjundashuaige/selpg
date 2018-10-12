
<!-- TOC -->

- [selpg](#selpg)
    - [unix type interface](#unix-type-interface)
    - [how exactly to using pflag](#how-exactly-to-using-pflag)
    - [determine whether the source text comes from](#determine-whether-the-source-text-comes-from)
    - [finally is the main pattern code](#finally-is-the-main-pattern-code)
    - [other stuff](#other-stuff)
    - [last function redirection the output to device like printer](#last-function-redirection-the-output-to-device-like-printer)
    - [some test](#some-test)
    - [simply perfect](#simply-perfect)

<!-- /TOC -->

# selpg
first thing first,i aint gonna lie,this whole project was designed just for homework in the course of service computing.

i wont even learn golang if i aint took that course.but you know,everything just wont workout like we used to think.

so lets go get them,first i am gonna intro some basic part about golang which used during the exper.then some tricky part about this problem.after that some conclusion.

## unix type interface
 since we want that interface not just look like someone just make it for linux daily using,but actually meets all the needs that linux wants.

so we just need to using something like 
```
--name=123 -n 123 
```
this kind of codes

but there comes a dilemma that its kindda hard for us to do all the analys and besides that there will be some bug undected.

so using third-party package seems to be the best solution.

there comes the pflag,but using that you can get easier to do with all the hard work instead of just busy in split the string.

## how exactly to using pflag

```
var input_start_line = flag.IntP("start","s",1,"target paragraph's start line")
var input_end_line = flag.IntP("end","e",1,"target paragraph's end line")
var input_page_size = flag.IntP("line","l",72,"how many line one page contain")
var error_writer = bufio.NewWriterSize(os.Stderr,1024)
```
all you need to do its to assign all the command into the flag.

## determine whether the source text comes from

+ well,by using different source of text,the best way is to desprate the api from the object like the reader.

so here comes that
```
file,err := os.Open(flag.Arg(flag.NArg()-1))
```


## finally is the main pattern code

the main function is easy as fuck,basic just get the range of code from the source text and lead the whole text into the new target text
```
var output_page string
		//the overall result
	for i:=0;;i++{
		line,_ := f.ReadString('\n')
		//check_err(err)
		var s string
		fmt.Sscan(line,&s)
		//read the line one by one
		if(len(s)==0){
			break
		}
		if(i>=( *input_start_line-1)*(*input_page_size) && i<=(*input_end_line)*(*input_page_size)){
			/*

			combine all line into one large string

			*/
			output_page +="\n"
			output_page +=s

			fmt.Println(s)
```

so freaking easy isnt it.

## other stuff

+ the other part of thing is that we have to deal with the error

golang support the error handling by applying something like 
```
panic()
```
```
recovery()
```
which is just like throw and catch in C.

which is pretty easy.

but what doesnt easy is that we have to make sure that we detect all the details that might lead to the bugs.
```
func negative_input(i *int)(error){
	if(*i < 0){
		return errors.New("input cant be negative")
	}
	return nil
}

/*

make sure the start position is definately small than end position

*/
func check_start_lower_end(start *int,end *int)(error){
	if(*start > *end){
		return errors.New("start is larger than end position")
	}
	return nil
}
```

## last function redirection the output to device like printer
```
func load_data(lpin io.WriteCloser,res string){

	defer lpin.Close()
	io.WriteString(lpin,res)

}

func call_go_routine(res string){
	if(*input_destination == ""){
		cmd := exec.Command("lp","-d"+*input_destination)
		lpin,err :=cmd.StdinPipe()
		if err!=nil{
			panic(err)
		}
		go load_data(lpin,res)
	}
}
```
basically we call the command using ``` exec.Command ```
by using something like this,we can get the output just like that after running this command in our machine.

but foucus on this command ``` lp ``` which need the input comes from us.the stdin or somewhere we redirection to.

so we using the commmand pipe to do so.

## some test

using each line with only the line number as a file input,

we can provide some test example and to check for the result with our own.

```
1
2
3
...
1024
```

with the command like ```selpg -s1 -e2 -l 72 numbers.txt```
and then get the result of 
```
1
2
3
...
71
```
which is exactly what i want as result.
## simply perfect