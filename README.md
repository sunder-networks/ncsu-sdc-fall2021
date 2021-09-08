## optional go library grammar generation procedure
```
# install
cd /usr/local/lib && curl -O https://www.antlr.org/download/antlr-4.9-complete.jar
export CLASSPATH=".:/usr/local/lib/antlr-4.9-complete.jar:$CLASSPATH"

# alias cli syntax
alias antlr4='java -Xmx500M -cp "/usr/local/lib/antlr-4.9-complete.jar:$CLASSPATH" org.antlr.v4.Tool'
alias grun='java -Xmx500M -cp "/usr/local/lib/antlr-4.9-complete.jar:$CLASSPATH" org.antlr.v4.gui.TestRig'

# generate
mkdir cdlang
antlr4 -o cdlang -Dlanguage=Go -visitor -listener -package cdlang CDLang.g4 

tree cdlang/
cdlang/
├── CDLang.interp
├── CDLang.tokens
├── CDLangLexer.interp
├── CDLangLexer.tokens
├── cdlang_base_listener.go
├── cdlang_base_visitor.go
├── cdlang_lexer.go
├── cdlang_listener.go
├── cdlang_parser.go
└── cdlang_visitor.go

0 directories, 10 files
```