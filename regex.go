//test regex
//regex.go
package main
import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"regexp"
	"log"
)

//fungsi membaca file eksternal
//format:
//keyword1;keyword2;keyword3
//text1;text2;text3
func read(path string) ([]string,error) {
	file,err := os.Open(path)
	if err != nil {
		return nil,err
	}
	defer file.Close()
	var lines []string
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		lines = append(lines,scan.Text())
	}
	return lines,scan.Err()
}

//fungsi periksa maksimum
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	
	//read file external
	in,erri := read("input.txt")
	if erri != nil {
		log.Fatalf("error: %s",erri)
	}
	//ambil keyword
	keywords := strings.Split(in[0],";")
	//print keyword
	for i:=0;i<len(keywords);i++ {
		fmt.Print(keywords[i]+" ")
	}
	fmt.Println()
	//ambil text per baris
	var line []string
	for i:=1;i<len(in);i++ {
		line = append(line,in[i])
	}
	//gabungkan text per baris
	text := strings.Join(line[:],"\n")
	fmt.Println(text) // print
	//pisahkan text per baris menurut separator 
	texts := strings.Split(text,";")
	//print text
	for i:=0;i<len(texts);i++ {
		fmt.Println(texts[i])
	}


	var check []int //periksa jumlah kata yang ditemukan
	var spam []int //list isi text yang merupakan spam
	var notspam []int //list isi text yang bukan spam
	for i:=0;i<len(texts);i++{
		check = append(check,0)
	}
	//test regex
	for i:=0;i<len(texts);i++ {
		amount:= 0 //jumlah kata yang ditemukan
		max_amount := 0
		for j:=0;j<len(keywords);j++ {
			//assignment agar pencarian dapat dilakukan
			a1 := texts[i]
			a2 := keywords[j]
			pattern := "(?i)"+a2
			r := regexp.MustCompile(pattern)
			match := r.FindAllStringIndex(a1,-1)
			amount =  len(match)//cari kata dengan regex
			max_amount = max(amount,max_amount)
		}
		check[i] = max_amount //jumlah kata yang ditemukan pada text ke-i
		fmt.Println(check[i])
		//jika tidak ditemukan, maka text bukan spam
		if check[i] == 0 {
			notspam = append(notspam,i+1)
		}
		//jika ada lebih dari 1, maka text tersebut spam
		if check[i] > 0 {
			spam = append(spam,i+1)
		}
	}
	//print text spam dan bukan spam
	fmt.Println("Spam: ",spam)
	fmt.Println("Not spam: ",notspam)

	//write ke file eksternal
	//format: spam1,spam2,spam3;notspam1,notspam2,notspam3
	out,erro := os.Create("result.txt")
	if erro != nil {
		log.Fatalf("error making file: %s",erro)
	}
	defer out.Close()
	//tulis yang spam
	if len(spam) > 0 {
		for i:=0;i<len(spam);i++{
			out.WriteString(fmt.Sprint(spam[i]))
			if i!=len(spam)-1 {
				out.WriteString(",")
			}
		}
	}
	//tulis yang bukan spam
	out.WriteString(";")
	if len(notspam) > 0 {
		for i:=0;i<len(notspam);i++{
			out.WriteString(fmt.Sprint(notspam[i]))
			if i!=len(notspam)-1 {
				out.WriteString(",")
			}
		}
	}

}