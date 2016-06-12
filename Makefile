
COMMAND    =  'help clean'
help:
	@echo $(COMMAND)

clean:
	@rm -rf *.pyc *.o
