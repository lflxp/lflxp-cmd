.PHONY: push pull install run

push: clean 
	git add .
	git commit -m "${m}"
	git push origin $(shell git branch|grep '*'|awk '{print $$2}')

pull: clean
	git pull origin $(shell git branch|grep '*'|awk '{print $$2}')

install: clean
	cp cmd/main.go . && go install
	lflxp-cmd -h

run:
	cd cmd && go run main.go

clean:
	rm -f main.go