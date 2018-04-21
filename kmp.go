//test KMP
//kmp.go
package main
import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"log"
)

//proses list panjang dengan prefix/suffix terpanjang
func preprocess(s string)[]int{
	l := len(s)
	var out []int
	out = make([]int,l,l)
	for i:=1;i<=l;i++ {
		max := 0
		for j:=1;j<i;j++ {
			if strings.HasSuffix(s[0:i],s[0:j]) == true {
				max = j
			}
		}
		out[i-1] = max
	}
	return out
}

//kembalikan jumlah karakter
func kmpSearch(pattern string, text string) int {
	var pre[] int = preprocess(pattern)
	var pos[] int
	j := 0
	for i:=0;i<len(text);i++ {
		if text[i] == pattern[j] {
			if j == (len(pattern)-1) {
				pos = append(pos,i-j)
				j = 0
			} else {
				j++
			}
		} else {
			if j > 0 {
				for j > 0 && text[i] != pattern[j] {
					j = pre[j-1]
				}
				if text[i] == pattern[j] {
					j++
				}
			}
		}
	}
	return len(pos)
}

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

func main() {
	//read file external
	in,erri := read("test.txt")
	if erri != nil {
		log.Fatalf("error: %s",erri)
	}
	//pecah hasil read file jadi keyword dan text
	var keywords,texts []string
	keywords = strings.Split(in[0],";")
	texts = strings.Split(in[1],";")

	//test print keyword
	fmt.Println("\nKeywords:")
	fmt.Println(len(keywords))
	for i:=0;i<len(keywords);i++{
		fmt.Print("("+keywords[i]+") ")
	}
	//test print text
	fmt.Println()
	fmt.Println("Texts:")
	fmt.Println(len(texts))
	for i:=0;i<len(texts);i++{
		fmt.Print(texts[i]+" ")
	}
	fmt.Println()

	var check []int //periksa jumlah kata yang ditemukan
	var spam []int //list isi text yang merupakan spam
	var notspam []int //list isi text yang bukan spam
	for i:=0;i<len(texts);i++{
		check = append(check,0)
	}
	//test KMP
	for i:=0;i<len(texts);i++ {
		amount:= 0 //jumlah kata yang ditemukan
		for j:=0;j<len(keywords);j++ {
			//assignment agar pencarian dapat dilakukan
			a1 := texts[i]
			a2 := keywords[j]
			amount = kmpSearch(a2,a1) //cari kata dengan KMP
		}
		check[i] = amount //jumlah kata yang ditemukan pada text ke-i
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
	out,erro := os.Create("result-kmp.txt")
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