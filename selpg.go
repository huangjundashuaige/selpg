package main

import(
	flag "github.com/spf13/pflag"
	"fmt"
	"bufio"
	"os"
	"errors"
	"os/exec"
	"io"
)
/*

set up for the pflag 

*/

var input_start_line = flag.IntP("start","s",1,"target paragraph's start line")
var input_end_line = flag.IntP("end","e",1,"target paragraph's end line")
var input_page_size = flag.IntP("line","l",72,"how many line one page contain")
var input_destination = flag.StringP("destination","d","","make sure the output can be sent to the target")
var error_writer = bufio.NewWriterSize(os.Stderr,1024)


/*

make sure the inputs are all legal

*/
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

/*

whole checking process

*/
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

/*

standard mechaism for handling the error

*/
func check_err(err error){
	if err != nil{
		error_writer.WriteString(err.Error())
		panic(err)
	}
}

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


func main(){
	flag.Parse() 
	/*

	dealing with the input, make sure which is which

	*/
	check_legal_input()
	

	var f *bufio.Reader //creat the reader object
	// to simplfy the read afterwards option
	
	
	if(flag.NArg()!=0){
	/*

		not providing the input file address

	*/
		file,err := os.Open(flag.Arg(flag.NArg()-1))
		// openning the file
		check_err(err)
		f = bufio.NewReader(file)
		//inititialize the object as the trash can	
	
	}else{
		/*

		normal procedure everything comes from stand input

		*/
		f = bufio.NewReader(os.Stdin)
	}
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
		}
	}
	if(*input_destination==""){
		/*

		redirection the output to the divice we want

		*/
		call_go_routine(output_page)
	}
}