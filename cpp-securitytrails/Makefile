CC = g++ 
 
INCLUDES = -I/includes
 
LIBINCLUDES = -L/includes -Lbuild/includes
 
INC = $(INCLUDES) $(LIBINCLUDES) 
CFLAGS= -std=c++11
 
OBJS =
 
LIBS =


securitytrails: main.o
	$(CC) $(CFLAGS) $(INC) main.o $(OBJS) $(LIBS)  -o $@
 
main.o:	main.cpp $(OBJS)
	$(CC) $(CFLAGS) $(INC) -c main.cpp

clean:
	rm -f *.o *~ securitytrails
	rm -f *.out
	rm -f securitytrails