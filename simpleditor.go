package main;
import (
	"os";
	"fmt";
	"strings";
	"bufio";
	);
func readln() string {
	scanner := bufio.NewScanner(os.Stdin);
	var standardinput string;
	if scanner.Scan() {
		standardinput = scanner.Text();
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err);
		return "";
	}
	standardinput = strings.ReplaceAll(standardinput,"\n","");
	return standardinput;
}
func main() {
	var command rune;
	var selection int;
	var characterinput rune;
	var file string;
	var input string;
	
	var buffer string;
	if len(os.Args) > 1{
		tmp,err := os.ReadFile(os.Args[1]);
		if(err != nil){
			fmt.Println(err);
		}else {
			file = os.Args[1];
			buffer = string(tmp)
			fmt.Println("#",len(buffer));
		}
	} else {
		fmt.Println("no file");
	}
	for command != 'q' {
		input = readln();
		if input == "" {
			continue;
		} else if []rune(input)[0] == '>'{
			var patt,replacement string;
			var moreparseableinput = []rune(input);
			var i int;
			if(moreparseableinput[0] == '>'){
				for i = 1; i<len(moreparseableinput); i++ {
					if moreparseableinput[i] == '\\' && i+1 < len(moreparseableinput){
						if moreparseableinput[i+1] == 'n' {
							patt+="\n";
						} else {
							patt+=string(moreparseableinput[i+1]);
						}
						i++;
						continue;
					}
					if moreparseableinput[i] == '/' {
						break;
					}
					patt+=string(moreparseableinput[i]);
				}
				if i == len(moreparseableinput)-1 {
					fmt.Println("expected pattern");
				} else {
					
					for i = i+1; i<len(moreparseableinput); i++ {
						if moreparseableinput[i] == '\\' && i+1 < len(moreparseableinput){
							if moreparseableinput[i+1] == 'n' {
								replacement+="\n";
							} else {
								replacement+=string(moreparseableinput[i+1]);
							}
							i++;
							continue;
						}
						replacement+=string(moreparseableinput[i]);
					}
					buffer = strings.ReplaceAll(buffer,patt,replacement);
				}
			}
		} else if n,err :=fmt.Sscanf(input,"%d+%c+%c",&selection,&command,&characterinput); n >=  3{
			if err != nil {
				fmt.Println(err);
			} else {
				if(selection <= len(buffer)){
					switch(command){
						case 'a':
							buffer += string(characterinput);
						case 's':
							var tmp string;
							for i,char := range buffer {
								if(selection-1 == i){
									tmp += string(characterinput);
								} else {
									tmp+= string(char);
								}
							}
							buffer = tmp;
					}
				}
			}
		}else if n,err :=fmt.Sscanf(input,"%d+%c",&selection,&command);  n >= 2 {
			if err != nil {
				fmt.Println(err);
			} else {
				if selection <= len(buffer){
					if command == 'd' {
						var tmp string;
						
						for i,char := range buffer {
							if(i == selection-1){
								continue;
							} else {
								tmp+=string(char);
							}
						}
						buffer = tmp;
					}
				}
			}
		} else if n,err :=fmt.Sscanf(input,"%c+%s",&command,&file); n >= 2 {
			if err != nil {
				fmt.Println(err);
			} else {
				if(command == 'w'){
					//resets permissions
					//not best practice
					err := os.WriteFile(file,[]byte(buffer),0644);
					if err != nil {
						fmt.Println(err);
					} else {
						fmt.Println("wrote #",len(buffer),"to",file);

					}
				} else if command == 'r' {
					tmp,err := os.ReadFile(file);
					if err != nil  {
						fmt.Println(err);
					} else {
						buffer = string(tmp);
						fmt.Println("#",len(buffer));
					}
				} else {
					fmt.Println("aint reading or writing");
				}
			}
		} else if n,err :=fmt.Sscanf(input,"%c",&command); n >= 1 {
			if err != nil {
				fmt.Println(err);
			} else {
				if command == 'w' && file != "" {
					//resets permissions
					//not best practice
					err := os.WriteFile(file,[]byte(buffer),0644);
					if err != nil {
						fmt.Println(err);
					} else {
						fmt.Println("wrote #",len(buffer),"to",file);
					}
				} else if command == 'q'{
					break;
				} else if command == 'p'{
					fmt.Println(buffer);
				} else {
					fmt.Println("?",string(input));
				}
			}
		} else {
			
			fmt.Println("?");
		}
	}


}
