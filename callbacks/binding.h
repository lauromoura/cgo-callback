#ifndef BINDING_H
#define BINDING_H


typedef int (*cb_type)(int, void *data);

void set_callback(cb_type callback, void *data);

int call_callback(int argument);

int translator(int value, void *data);
void set_translator(long key);

#endif
