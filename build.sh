#!/bin/bash
TARGET=`basename "$PWD"`
echo building $TARGET...
mkdir -p build
GOOS=js GOARCH=wasm go build -o build/$TARGET.wasm .
#brotli -f -Z --suffix=-brotli build/$TARGET.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js build
touch build/favicon.ico
cat << EOF > build/index.html
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <script src="wasm_exec.js"></script>
    <script>
      const go = new Go()
      WebAssembly.instantiateStreaming(
        fetch('$TARGET.wasm'),
        go.importObject
      ).then(result => {
        go.run(result.instance)
      })
    </script>
  </head>
  <body></body>
</html>
EOF
xdg-open http://localhost:8000
cd build && python3 -m http.server 
