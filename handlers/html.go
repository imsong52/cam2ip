// Package handlers.
package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// HTML handler.
type HTML struct {
	Template []byte
}

// NewHTML returns new HTML handler.
func NewHTML(bind string, width, height float64) *HTML {
	h := &HTML{}

	b := strings.Split(bind, ":")
	if b[0] == "" {
		bind = "127.0.0.1" + bind
	}

	html = strings.Replace(html, "{BIND}", bind, -1)
	html = strings.Replace(html, "{WIDTH}", fmt.Sprintf("%.0f", width), -1)
	html = strings.Replace(html, "{HEIGHT}", fmt.Sprintf("%.0f", height), -1)

	h.Template = []byte(html)
	return h
}

// ServeHTTP handles requests on incoming connections.
func (h *HTML) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "HEAD" {
		msg := fmt.Sprintf("405 Method Not Allowed (%s)", r.Method)
		http.Error(w, msg, http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(h.Template)
}

var html = `<html>
    <head>
        <title>cam2ip</title>
        <style>
        body {
            width: 100%;
            height: 100%;
            padding: 0;
            margin: 0;
            background-color: #000000;
        }

        div {
            width: 100%;
            height: 100%;
            position: relative;
        }

        canvas {
            height: auto;
            width: auto;
            max-height: 100%;
            max-width: 100%;
            display: block;
            position: absolute;
            top: 0;
            bottom: 0;
            left: 0;
            right: 0;
            box-sizing: border-box;
            margin: auto;
        }
        </style>
        <script>
        var url = "ws://{BIND}/socket";
        ws = new WebSocket(url);

        ws.onopen = function() {
            console.log("onopen");
        }

        ws.onmessage = function(e) {
			var context = document.getElementById("canvas").getContext("2d");
			
			var image = new Image();
			image.onload = function() {
				context.drawImage(image, 0, 0);
			}

            image.setAttribute("src", "data:image/jpeg;base64," + e.data);
        }

        ws.onclose = function(e) {
            console.log("onclose");
        }

        ws.onerror = function(e) {
            console.log("onerror");
        }
        </script>
    </head>
    <body>
        <div><canvas id="canvas" width="{WIDTH}" height="{HEIGHT}"></canvas></div>
    </body>
</html>`
