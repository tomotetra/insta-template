setup:
	go install
	go build
	cp instagenTags.txt $(HOME)/.instagenTags
	mv instagen /usr/local/bin/

dev/setup:
	dep ensure

dev/run:
	go run main.go -T $(TITLE) -t $(TAGS) $(TARGET)
