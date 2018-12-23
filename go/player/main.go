package main

import "fmt"
import "errors"
import "strings"
import "bufio"
import "os"
import "music"
import "strconv"

/*****
<menu>
- L: list all music
    Show info: <index>:<song name>
- P <index>: play some music with the <index>
	Playing logic(virtual): show the information of music
	#Playing logic(real): 
	#  - Once play one music, the time is fixed(3 min=180 sec)
	#  - While playing, the info that shown in the screen is: <song name>: playing <n> secs; remain <m> secs;
	#  - When playing over, Return to the main menu(show the <main menu info>)
	#Only one operation is allowed in the playing, that is pause; when music is pausing, the playing time is stopped;
- A <singer>:<song name>
	Add some music into the library
- D <index>: Deleting the song from the library
- X: exit the program

<main menu info>: L(list music); P <index>(playing); A <singer>:<song name>; D <index>(deleting music); X(exit the program)

*****/

var MAIN_MENU string = "L(list music); P <index>(playing); A <singer>:<song name>; D <index>(deleting music); X(exit the program)"

func showMainMenu() {
    fmt.Println(MAIN_MENU)
}

func extractValidParams(s string) (params []string, err error){
    if len(s) == 0 {
        return params, err
    }
    if (string(s[0]) == "A") {
        //把一串字符串提取出有效的成分
        parts := strings.Split(s, " ")
        if (len(parts) >= 3) {
            if (parts[0] == "A") {
                params = append(params, parts[0], parts[1], parts[2])
                return
            }
        }
        err = errors.New("Param input error")
        return
    }
    if (string(s[0]) == "D") {
        //把一串字符串提取出有效的成分
        parts := strings.Split(s, " ")
        if (len(parts) >= 2) {
            if (parts[0] == "D") {
                params = append(params, parts[0], parts[1])
                return
            }
        }
        err = errors.New("Param input error")
        return
    }
    if (string(s[0]) == "P") {
        //把一串字符串提取出有效的成分
        parts := strings.Split(s, " ")
        if (len(parts) >= 2) {
            if (parts[0] == "P") {
                params = append(params, parts[0], parts[1])
                return
            }
        }
        err = errors.New("Param input error")
        return        
    }
    if (string(s[0]) == "L") {
        if (s == "L") {
            return
        } else {
            err = errors.New("Param error")
            return
        }
    }
    return params, err
}

func main() {
    //首先输出main menu tips
    showMainMenu()
    
    recorder := music.NewLib()

    var f string
    var fb []byte
	for {
        fmt.Println("Please input:")
        fb, _, _ = bufio.NewReader(os.Stdin).ReadLine()
        f = string(fb)
        //fmt.Println(f)
        if (len(f) < 1) {
            continue
        }
        if (string(f[0]) == "X") {
            break
        }
        if (string(f[0]) == "A") {
            //add
            params, err := extractValidParams(f)
            if err == nil {
                recorder.Add(params[1], params[2])
            }
            continue
        }
        
        if (string(f[0]) == "P") {
            //Play
            params, err := extractValidParams(f)
            if err == nil {
                fmt.Println(params)
                i, _ := strconv.Atoi(params[1])
                recorder.Play(i)
            }
            continue
        }
        
        if (string(f[0]) == "D") {
            //Delete
            params, err := extractValidParams(f)
            if err == nil {
                i, _ := strconv.Atoi(params[1])
                recorder.Delete(i)
            }
            continue
        }
        
        if (string(f[0]) == "L") {
            //list
            recorder.Show()
            continue
        }

        break
    }
    fmt.Println("It is over, and program exits!")
}