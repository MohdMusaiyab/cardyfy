package handlers

import "net/http"

func GenerateImageHandler(w http.ResponseWriter, r *http.Request) {
    // Write your image generation logic here
    w.Write([]byte("Image generated!"))
}
