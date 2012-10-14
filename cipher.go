package main 

import ( 
        "crypto/aes" 
        "crypto/cipher"        
        "io" 
        "os" 
        "log"
) 

func EncryptAndGzip(dstfile, srcfile string, key, iv []byte) { 
        r, _ := os.Open(srcfile) 
        var w io.Writer 
        w, _ = os.Create(dstfile) 
        c, err := aes.NewCipher(key)
        if err != nil {
        	log.Fatal(err)
        } 
        w = &cipher.StreamWriter{S: cipher.NewOFB(c, iv), W: w}          
        io.Copy(w, r)        
} 

func DecryptAndGunzip(dstfile, srcfile string, key, iv []byte) { 
        f, _ := os.Open(srcfile) 
        defer f.Close() 
        c, err := aes.NewCipher(key)
        if err != nil {
        	log.Fatal(err)
        }         
        r := &cipher.StreamReader{S: cipher.NewOFB(c, iv), R: f}         
        w, _ := os.Create(dstfile) 
        defer w.Close() 
        io.Copy(w, r) 
} 

func main() { 
        EncryptAndGzip("/tmp/passwd", "/etc/passwd", make([]byte, 16), 
make([]byte, 16)) 
        DecryptAndGunzip("/dev/stdout", "/tmp/passwd", make([]byte, 16), 
make([]byte, 16)) 
}
