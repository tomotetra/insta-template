setup:
	dep ensure
	go install
	cp common_tags.txt $(HOME)/.instatemplate_tags

run:
	go run main.go -T $(TITLE) -t $(TAGS) $(TARGET)
