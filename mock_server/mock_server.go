package main
import (
	"io"
	"net/http"
)

func MockHandler(w http.ResponseWriter, r *http.Request) {
	msg := `{"me":{"type":"human","size":3},"start":{"type":"start","x":0,"y":0},"end":{"type":"end","x":100,"y":100},"obstacles":{"stable":[{"type":"block","sx":1,"sy":2,"ex":10,"ey":5},{"type":"circle","x":3,"y":5,"size":10}],"floating":[{}]},"background":{"type":"background","height":100,"width":100},"controls":[{"type":"up"},{"type":"down"},{"type":"left"},{"type":"right"}],"resources":{"human":{"url":"45.76.13.24:12321/smart/v0.1/resource/human.svg"},"start":{"url":"45.76.13.24:12321/smart/v0.1/resource/start.svg"},"end":{"url":"45.76.13.24:12321/smart/v0.1/resource/end.svg"},"block":{"url":"45.76.13.24:12321/smart/v0.1/resource/block.svg"},"up":{"url":"45.76.13.24:12321/smart/v0.1/resource/up.svg"},"down":{"url":"45.76.13.24:12321/smart/v0.1/resource/down.svg"},"left":{"url":"45.76.13.24:12321/smart/v0.1/resource/left.svg"},"right":{"url":"45.76.13.24:12321/smart/v0.1/resource/right.svg"},"background":{"url":"45.76.13.24:12321/smart/v0.1/resource/background.svg"}}}`

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, msg)
}

func main() {
	http.HandleFunc("/", MockHandler)
	http.ListenAndServe("0.0.0.0:12300", nil)
}
