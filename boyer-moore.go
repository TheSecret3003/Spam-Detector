//test Boyer-Moore
//boyer-moore.go
package main
import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"log"
)
type bm struct {
	pattern string
	badchar[256] int
	suffix []int
}
//fungsi pembuat string finder Boyer-Moore
func BMstringfinder(pattern string) * bm {
	f := &bm{
		pattern: pattern,
		suffix: make([]int,len(pattern)),
	}
	//buat tabel badchar
	last := len(pattern)-1
	for i := range f.badchar {
		f.badchar[i] = len(pattern)
	}
	for i:=0;i<last;i++ {
		f.badchar[pattern[i]] = last-i
	}
	//buat tabel suffix
	lastprefix := last
	for i:=last;i>=0;i--{
		if strings.HasPrefix(pattern,pattern[i+1:]) {
			lastprefix = i+1
		}
		f.suffix[i] = lastprefix+last-i
	}
	for i:=0;i<last;i++ {
		lens := longestCommonSuffix(pattern,pattern[1:i+1])
		if pattern[i-lens]!=pattern[last-lens] {
			f.suffix[last-lens] = lens+last-i
		}
	}
	return f
}
//fungsi penentu suffix terpanjang
func longestCommonSuffix(a,b string) (i int) {
	for ; i<len(a)&&i<len(b);i++ {
		if a[len(a)-1-i]!=b[len(b)-1-i] {
			break
		}
	}
	return
}
//fungsi untuk maju ke index selanjutnya
func (f *bm) next(text string) int {
	i := len(f.pattern)-1
	for i<len(text) {
		j := len(f.pattern)-1
		for j>=0&&text[i] == f.pattern[j] {
			i--
			j--
		}
		if j < 0 {
			return i+1 //match
		}
		i += max(f.badchar[text[i]],f.suffix[j])
	}
	return -1
}
//fungsi periksa maksimum
func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}
//fungsi mendapatkan jumlah kata
func getwords(s,searching string) int {
	count := 0
	sf := BMstringfinder(searching)
	idx := sf.next(s)
	for idx!=-1 {
		s = s[idx+len(searching):]
		idx = sf.next(s)
		count++
	}
	return count
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
	//test boyer-moore
	for i:=0;i<len(texts);i++ {
		amount:= 0 //jumlah kata yang ditemukan
		for j:=0;j<len(keywords);j++ {
			//assignment agar pencarian dapat dilakukan
			a1 := texts[i]
			a2 := keywords[j]
			amount = getwords(a1,a2) //cari kata dengan Boyer-Moore
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