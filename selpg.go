package main

import(
	flag "github.com/spf13/pflag"
	"fmt"
	"bufio"
	"os"
	"errors"
)

var input_start_line = flag.IntP("start","s",1,"target paragraph's start line")
var input_end_line = flag.IntP("end","e",1,"target paragraph's end line")
var input_page_size = flag.IntP("line","l",72,"how many line one page contain")
var error_writer = bufio.NewWriterSize(os.Stderr,1024)

func negative_input(i *int)(error){
	if(*i < 0){
		return errors.New("input cant be negative")
	}
	return nil
}

func check_start_lower_end(start *int,end *int)(error){
	if(*start > *end){
		return errors.New("start is larger than end position")
	}
	return nil
}

func check_legal_input(){
	if err:=negative_input(input_start_line);err!=nil{
		error_writer.WriteString(err.Error())
		error_writer.Flush()
		panic(err)
	}
	if err:=negative_input(input_end_line);err!=nil{
		error_writer.WriteString(err.Error())
		error_writer.Flush()
		panic(err)
	}	
	if err:=negative_input(input_page_size);err!=nil{
		error_writer.WriteString(err.Error())
		error_writer.Flush()
		panic(err)
	}
	if err:=check_start_lower_end(input_start_line,input_end_line);err!=nil{
		error_writer.WriteString(err.Error())
		error_writer.Flush()
		panic(err)
	}
}

func check_err(err error){
	if err != nil{
		error_writer.WriteString(err.Error())
		panic(err)
	}
}

func main(){
	flag.Parse()
	check_legal_input()
	var f *bufio.Reader
	if(flag.NArg()!=0){
		file,err := os.Open(flag.Arg(flag.NArg()-1))
		check_err(err)
		f = bufio.NewReader(file)	
	}else{
		f = bufio.NewReader(os.Stdin)
	}
		var output_page string
	for i:=0;;i++{
		line,_ := f.ReadString('\n')
		//check_err(err)
		var s string
		fmt.Sscan(line,&s)
		if(len(s)==0){
			break
		}
		if(i>=( *input_start_line-1)*(*input_page_size) && i<=(*input_end_line)*(*input_page_size)){
			output_page +="\n"
			output_page +=s

			fmt.Println(s)
		}
	}
}