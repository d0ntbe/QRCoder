package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/skip2/go-qrcode"
)

func main() {

	//link := "deeplink://aplication/api/v1/ppp"
	//link := "deeplink://aplication/api/v1/ppp?title=hello#popla=gopla#index=123#toto=3v"

	if len(os.Args) <= 1 {
		fmt.Println(" ")
		fmt.Println("You forget the args. See Usage! Type: h/-h/help")
		fmt.Println("With Love from ZL...")
		return
	}

	if os.Args[1] == "-h" || os.Args[1] == "-help" || os.Args[1] == "help" || os.Args[1] == "h" {

		fmt.Println("USAGE: ")
		fmt.Println("Example: go run main.go deeplink://aplication/api/v1/ppp")
		fmt.Println("Example: go run main.go deeplink://aplication/api/v1/ppp?title=hello#popla=gopla#index=123#toto=3v")
		fmt.Println("Example: QRCoder.exe superbank://linktoapp/p2p")
		fmt.Println("If use Windows - do not forgot double quotes. Example: QRCoder.exe ''ddd://ooooo?dsdsasd=123&dddaq=1'' ")
		fmt.Println("Payloads are in payloads.txt. But you can also add your specified payloads, line by line.")
		fmt.Println("--<-@  With Love from ZL...")
		return
	}

	link := os.Args[1]

	dirr := "qrs"

	_, err := os.Stat(dirr)
	if err != nil {
		err = os.Mkdir(dirr, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	files, err := filepath.Glob(filepath.Join(dirr, "*"))
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.OpenFile("payloads.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		fmt.Println("Check 0777 ...")
		log.Fatalf("open error: %v", err)
		return
	}
	defer f.Close()

	payloads := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		payloads = append(payloads, sc.Text())
		if err := sc.Err(); err != nil {
			log.Fatalf("Get payloads file error: %v", err)
			return
		}
	}

	query := make([]string, 0)
	params := make([]string, 0)
	rezstrikes := make([]string, 0)
	fl := 0
	if strings.Contains(link, "?") {
		query = strings.SplitAfter(link, "?")
		params = strings.SplitAfter(query[1], "&")
		fl = 1

	}
	if strings.Contains(link, "#") {
		query = strings.SplitAfter(link, "?")
		params = strings.SplitAfter(query[1], "#")
		fl = 1
	}
	if fl == 0 {

		if strings.Index(link, "://") == -1 {
			fmt.Println(" ")
			fmt.Println("Your link must be with ://. Link formarts  are often with pampampam:// for QR ")
			return
		}
		p3 := link[:strings.Index(link, "://")]

		for j := range payloads {
			p1 := payloads[j]

			p2 := link + p1
			err = qrcode.WriteColorFile(p2, qrcode.Medium, 256, color.Black, color.White, dirr+"\\secondfile"+strconv.Itoa(int(time.Now().UnixNano()))+".png") // strconv.Itoa(n)
			if err != nil {
				fmt.Printf("Sorry couldn't create qrcode:,%v", err)
			}
			rezstrikes = append(rezstrikes, p2)

			p2 = link + "?" + p1
			err = qrcode.WriteColorFile(p2, qrcode.Medium, 256, color.Black, color.White, dirr+"\\secondfile"+strconv.Itoa(int(time.Now().UnixNano()))+".png")
			if err != nil {
				fmt.Printf("Sorry couldn't create qrcode:,%v", err)
			}
			rezstrikes = append(rezstrikes, p2)

			err = qrcode.WriteColorFile(p3+"://"+p1, qrcode.Medium, 256, color.Black, color.White, dirr+"\\secondfile"+strconv.Itoa(int(time.Now().UnixNano()))+".png")
			if err != nil {
				fmt.Printf("Sorry couldn't create qrcode:,%v", err)
			}
			rezstrikes = append(rezstrikes, p3+"://"+p1)
		}
	}
	paramsTrue := make([]string, len(params))
	copy(paramsTrue, params)
	if fl == 1 {
		for i := range params {

			for j := range payloads {
				p1 := payloads[j]

				k := strings.Index(params[i], "=")
				params[i] = params[i][:k+1] + p1 + "&"

				if string(params[len(params)-1][len(params[len(params)-1])-1]) == "&" || string(params[len(params)-1][len(params[len(params)-1])-1]) == "#" {

					params[i] = string(params[i][:(len(params[i]) - 1)]) //delete & or # when last symbol in last word inside query payloads ... lol

				}

				rez := strings.Join(params, "")

				err = qrcode.WriteColorFile(query[0]+rez, qrcode.Medium, 256, color.Black, color.White, dirr+"\\secondfile"+strconv.Itoa(int(time.Now().UnixNano()))+".png") //strconv.Itoa(n)
				if err != nil {
					fmt.Printf("Sorry couldn't create qrcode:,%v", err)

				}
				rezstrikes = append(rezstrikes, query[0]+rez)

			}
			copy(params, paramsTrue)
		}
	}

	//pngs, err := os.ReadDir("qrs")
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(" ")
	fmt.Println("--<-@  QR examples are created inside qrs folder.....")
	fmt.Println("--<-@  Generating HTML page....")

	pngs := make([]string, 0)
	filepath.Walk("qrs", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		pngs = append(pngs, path)
		return nil
	})
	htmlstart := "<!DOCTYPE html><html><head><meta charset='UTF-8'><title>QRCoder</title><style>body { background-color: black; } table { width: 100%; border-collapse: collapse; } th, td { padding: 15px; overflow: hidden; max-width: 100px; text-align: center; border: 1px solid yellow; color: yellow; width: 200px; left: 0;} h1 {color: yellow; text-align: center;} div {color: yellow; font-size: 30px; text-align: right;} a {color: yellow; font-size: 30px; text-align: center;}</style></head><div >Follow: <a href=https://t.me/zeropticum>[tg]</a><a href=https://zeropticum.org>[zeroticum.org]</a></div><h1>QRCoder ▩ Zeropticum</h1><table>"
	htmlend := "</table><body></body></html>"

	addpng := make([]string, 0)

	for i := range pngs {

		if i == 0 {
			addpng = append(addpng, "<td><img src='"+pngs[i]+"'><h3>"+rezstrikes[i]+"</h3></td>")

		}

		if i%4 == 0 && i != 0 {
			addpng = append(addpng, "<tr></tr>")
			addpng = append(addpng, "<td><img src='"+pngs[i]+"'><h3>"+rezstrikes[i]+"</h3></td>")

		}
		if i%4 != 0 {

			addpng = append(addpng, "<td><img src='"+pngs[i]+"'><h3>"+rezstrikes[i]+"</h3></td>")

		}

	}

	//html constructor (lego) =)
	rezaddpng := strings.Join(addpng, "") //
	html := htmlstart + rezaddpng + htmlend

	file, err := os.Create("qrs.html")
	if err != nil {
		log.Fatal("Error File qrs.html Create")
	}
	defer file.Close()

	_, err = file.WriteString(html)

	fmt.Println("--<-@  DONE! Look:   qrs.html")
	fmt.Println("--<-@  With Love from ZL...")

}
