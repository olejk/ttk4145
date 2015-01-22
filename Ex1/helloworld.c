// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld.c -lpthread

#include <pthread.h>
#include <stdio.h>

int i=0;

void* tellOpp(){
	int j;
	for (j=0; j<1000000; j++)
		i++;
	return NULL
}

void* tellNed(){
	int j;
	for (j=0; j<1000000; j++)
		i--;
	return NULL
}

int main(){
	pthread_t ink;
	pthread_t dek;

	pthread_create(&ink, NULL, tellOpp, NULL);
	pthread_create(&ink, NULL, tellOpp, NULL);

	pthread_join(ink, NULL);
	pthread_join(dek, NULL);
	printf("%d\n", i);

	return 0;
}
