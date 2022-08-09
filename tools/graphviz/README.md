```shell
go install github.com/davidschlachter/embedded-struct-visualizer@latest

embedded-struct-visualizer -out demo.gv .

# https://graphviz.org/download/
dot -Tps fancy.gv -o fancy.ps

```