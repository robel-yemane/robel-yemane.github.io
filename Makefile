# What do I want to do automate?

# 1. generation of blog html file once text is ready - go run with arguments
# 2. go test, go lint
# 3. build dockerfile
# 4.
article ?= "/Users/snowball/Projects/git/robel-yemane.github.io/automate/blog.txt"

# run locally create blog html  file
gen_blog:
	cd automate && go run cmd/articlegen/articlegen.go -a $(article)



