article  ?= "/Users/snowball/Projects/git/robel-yemane.github.io/automate/blog.txt"
EMAIL    ?=
TWITTER  ?=
LINKEDIN ?=

# run locally create blog html  file
gen_blog:
	@cd automate && go run cmd/articlegen/articlegen.go -a $(article) \
		--email "$(EMAIL)" --twitter "$(TWITTER)" --linkedin "$(LINKEDIN)"



