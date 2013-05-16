package main

import (
	"code.google.com/p/go.net/websocket"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.Handle("/echo", websocket.Handler(echo))
	http.Handle("/data", websocket.Handler(data))
	log.Print("The server is listening...")
	http.ListenAndServe("0.0.0.0:8000", nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	rootTemplate.Execute(w, rootTemplate)
}

func echo(c *websocket.Conn) {
	c.Write([]byte("mama are mere"))
}

func data(c *websocket.Conn) {
	imgs, _ := filepath.Glob("*.jpg")
	for _, img := range imgs {
		log.Print(img)
		data, _ := ioutil.ReadFile(img)
		log.Print(data)
		c.Write(data)
	}
}

var rootTemplate = template.Must(template.New("root").Parse(`
<!DOCTYPE html>
<html>
<body>
<h3>Image sequence</h3>
<canvas id="myCanvas" width="640" height="272"></canvas>
<script src="http://notmasteryet.github.com/jpgjs/jpg.js"></script>
<script>
  var s = new WebSocket("ws://rif.homelinux.org/data");
  s.binaryType = "arraybuffer";
  var canvas = document.getElementById('myCanvas');
  var context = canvas.getContext('2d');

  s.onmessage = function(e) {
     var j = new JpegImage();
     j.parse(new Uint8Array(e.data));
     var id = context.getImageData(0,0,j.width, j.height);
     j.copyToImageData(id);
     context.putImageData(id, 0,0);
  }  
  s.onopen = function() {
    s.send("start");
  }  

</script>
</body>
</html>

`))
