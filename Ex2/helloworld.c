// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o helloworld_c helloworld.c -lpthread

#include <pthread.h>
#include <stdio.h>

int i=0;

void* tellOpp(void* mutex){
	int j;
	for (j=0; j<1000000; j++){
		pthread_mutex_lock(mutex);
		i++;
		pthread_mutex_unlock(mutex);
	}
	return NULL;
}

void* tellNed(void* mutex){
	int j;
	for (j=0; j<1000000; j++){
		pthread_mutex_lock(mutex);
		i--;
		pthread_mutex_unlock(mutex);
	}
	return NULL;
}

int main(){
	pthread_mutex_t mutex;
	pthread_mutex_init(&mutex, NULL);

	pthread_t ink;
	pthread_t dek;

	pthread_create(&ink, NULL, tellOpp, (void*) &mutex);
	pthread_create(&dek, NULL, tellNed, (void*) &mutex);

	pthread_join(ink, NULL);
	pthread_join(dek, NULL);
	printf("%d\n", i);

	pthread_mutex_destroy(&mutex);

	return 0;
}
