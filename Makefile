setup:
	go install
	cp instagenTags.txt $(HOME)/.instagenTags

dev/setup:
	dep ensure

dev/run:
	go run main.go -T $(TITLE) -t $(TAGS) $(TARGET)
