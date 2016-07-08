
COMMAND    =  'help clean'
help:
	@echo $(COMMAND)

test:
	@py.test python/tests/$(name)

clean:
	@find . -name '*.pyc' -and -name '*.o' -type f -delete

unittest:
	@py.test --capture=no --exitfirst python/tests/$(name)
