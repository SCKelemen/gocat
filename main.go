package main

import (
	"fmt"	
	"github.com/fatih/color"
	"go/token"
	"github.com/golang-collections/collections/stack" 
)

func main() {

printf("%s%s%s\n", BgGreenString("Success!"), grayscale(14)("/"), color.CyanString("%d.%d %s", 1,1,"fuck"))
printf("%s%s%s\n", BgYellowString("Warning!"), grayscale(14)("/"), color.CyanString("%d.%d %s", 1,1,"fuck"))
printf("%s%s%s\n", BgRedString("Error!"), grayscale(14)("/"), color.CyanString("%d.%d %s", 1,1,"fuck"))
printf("%s%s%s\n", SuccessString("Success!"), grayscale(14)("/"), color.CyanString("%d.%d %s", 1,1,"fuck"))
printf("%s%s%s\n", WarningString("Warning!"), grayscale(14)("/"), color.CyanString("%d.%d %s", 1,1,"fuck"))
printf("%s%s%s\n", ErrorString("Error!"), grayscale(14)("/"), color.CyanString("%d.%d %s", 1,1,"fuck"))
println(StatusColor(-1))
printf("Status: %s \n", StatusColor(0))
printf("Status: %s \n", StatusColor(1))
printf("Status: %n \n", StatusColor(99))
printf("Status: %n \n", StatusColor(100))
printf("Status: %n \n", StatusColor(101))
printf("Status: %n\n", StatusColor(199))
println("halp")
println(StatusColor(200))
println("/halp")
printf("Status: %n\n", StatusColor(201))
printf("Status: %n\n", StatusColor(299))
printf("Status: %n\n", StatusColor(300))
printf("Status: %n\n", StatusColor(301))
printf("Status: %n\n", StatusColor(399))
printf("Status: %n\n", StatusColor(400))
printf("Status: %n\n", StatusColor(401))
printf("Status: %n\n", StatusColor(499))
printf("Status: %n\n", StatusColor(500))
printf("Status: %n\n", StatusColor(501))
printf("Status: %n\n", StatusColor(599))
printf("Status: %n\n", StatusColor(600))

print(BgBlueString("Tas"), BgCyanString("te t"), BgGreenString("he r"), BgYellowString("ain"), BgRedString("bow"), BgMagentaString("!!!"), "\r\n")

for i := 0; i < 15; i++ {
	print(nthBracketColor(i))
	print(" ")
    }
print("\r\n")

}

func colorString(format string, p color.Attribute, a ...interface{}) string {
	c := color.New(p)

	if len(a) == 0 {
		return c.SprintFunc()(format)
	}

	return c.SprintfFunc()(format, a...)
}

func printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(color.Output, format, a...)
}

func grayscale(code color.Attribute) func(string, ...interface{}) string {
	return color.New(code + 232).SprintfFunc()
}

func BgGreenString(format string, a ...interface{}) string { return colorString(format, color.BgGreen, a...) }
func BgYellowString(format string, a ...interface{}) string { return colorString(format, color.BgYellow, a...) }
func BgRedString(format string, a ...interface{}) string { return colorString(format, color.BgRed, a...) }
func BgBlueString(format string, a ...interface{}) string { return colorString(format, color.BgBlue, a...) }
func BgCyanString(format string, a ...interface{}) string { return colorString(format, color.BgCyan, a...) }
func BgMagentaString(format string, a ...interface{}) string { return colorString(format, color.BgMagenta, a...) }

func SuccessString(format string, a ...interface{}) string { return colorString(format, color.BgGreen, a...) }
func WarningString(format string, a ...interface{}) string { return colorString(format, color.BgYellow, a...) }
func ErrorString(format string, a ...interface{}) string { return colorString(format, color.BgRed, a...) }


func statusColor(code int) string {
	switch {
                case code >= 100 && code < 200:
                        return "blue" 		//informational
		case code >= 200 && code < 300:
			return "green" 		//success
                case code >= 300 && code < 400:
                        return "cyan" 		//redirects
                case code >= 400 && code < 500:
                        return "yellow"		//client errors
                case code >= 500 && code < 600:
                        return "red" 		//server errors
		default:
			return "magenta"	//non standard
		
	}
}

func StatusColor(code int) string {
        switch {
                case code >= 100 && code < 200:
                        return colorString("%n", color.BgBlue, code)           //informational
                case code >= 200 && code < 300:
                        return colorString("%n", color.BgGreen, code)              //success
                case code >= 300 && code < 400:
                        return colorString("%n", color.BgCyan, code)               //redirects
                case code >= 400 && code < 500:
                        return colorString("%n", color.BgYellow, code)             //client errors
                case code >= 500 && code < 600:
                        return colorString("%n", color.BgRed, code)                //server errors
                default:
                        return colorString("%s", color.BgMagenta, code)            //non standard

        }
}


func methodColor(method string) string {
//	switch {
//		case "GET":
//			return "blue"
//		case "POST":
//			return "cyan"
//		case "PUT":
//			return "yellow"
//		case "PATCH":
//			return "green"
//		case "DELETE":
//			return "red"
//		case "HEAD":
//			return "magenta"
//		case "OPTIONS":
//			return "white"
//		case "LINK":
//		case "UNLINK":
//		case "PURGE":
//			return "hired"
//		case "COPY":
//		case "LOCK":
//		case "UNLOCK":
//		case "PROPFIND":
//		case "VIEW":
//	}
return "lllll"
}



//nthBracketColor returns the colored string
//for the pair of matching braces such that
//every level deep is highlighted with 
//a different color, otherwise returns 
//"FUUUUUCK". Example: 
// blue{  cyan{  green{  }green  }cyan  }blue
func nthBracketColor(c int) string {
	switch c % 6 {
		case 0: return BgBlueString("{")
		case 1: return BgCyanString("{")
		case 2: return BgGreenString("{")
		case 3: return BgYellowString("{")
		case 4: return BgRedString("{")
		case 5: return BgMagentaString("{")
		default:
			return "FUUUUUUUCK"
	}
	
}


func colorBrackets(code string) {
	s := stack.New()
	for i := 0; i < len(code); i++ {
        	switch code[i] {
		 case '{': 
			stack.Push(i)
			break
		 case '}':
		}
    	} 
}


