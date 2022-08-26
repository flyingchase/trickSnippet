#include <cstdio>
#include <cstdlib>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/_pthread/_pthread_cond_t.h>
#include <sys/_pthread/_pthread_mutex_t.h>
#include <type_traits>
#define BUFFER_SIZE 10

static int buffer[BUFFER_SIZE] = {0};
static int count = 0;
pthread_t consumer, producer;
pthread_cond_t cond_producer, cond_consumer;
pthread_mutex_t mutex;

void *consume(void *_) {
    while (1) {
        pthread_mutex_lock(&mutex);
        while (count == 0) {
            printf("empty buff, wait producer\n");
            pthread_cond_wait(&cond_consumer, &mutex);
        }

        count--;
        printf("consume a item\n");

        pthread_mutex_unlock(&mutex);
        pthread_cond_signal(&cond_producer);
    }
    pthread_exit(0);
}

void *produce(void *_) {
    while (1) {
        pthread_mutex_lock(&mutex);
        while (count == BUFFER_SIZE) {
            printf("full buffer, wait consumer\n");
        }

        count++;
        printf("produce a item.\n");
        pthread_mutex_unlock(&mutex);
        pthread_cond_signal(&cond_consumer);
    }
    pthread_exit(0);
}

int main() {
    pthread_mutex_init(&mutex, nullptr);
    pthread_cond_init(&cond_consumer, nullptr);
    pthread_cond_init(&cond_producer, nullptr);
    int err = pthread_create(&consumer, nullptr, consume, (void *)nullptr);

    if (err != 0) {
        printf("consumer thread created failed\n");
        exit(1);
    }

    /* err = pthread_create(&produce, nullptr, produce, (void *)nullptr); */
    if (err != 0) {
        printf("consumer thread created failed\n");
        exit(1);
    }

    return 0;
}
