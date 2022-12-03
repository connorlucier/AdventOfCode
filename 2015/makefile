default:
	g++ -o day$(day).o $(day)/*.cpp
run:
	@make day=$(day)
	@./day$(day).o $(day)
all:
	@for d in $(shell ls -d */) ; do \
		make day=$${d%/*} ; \
	done
clean:
	rm -rf *.o *.out */*.gch
