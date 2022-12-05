default:
	g++ -o day$(day).o $(day)/*.cpp -std=c++11
run:
	@make day=$(day)
	@./day$(day).o $(day)
all:
	@for d in $(shell ls -d */) ; do \
		make day=$${d%/*} ; \
	done
new:
	@mkdir $(day)
	@touch $(day)/main.cpp
	@touch $(day)/part1.h
	@touch $(day)/part2.h
clean:
	rm -rf *.o *.out */*.gch
