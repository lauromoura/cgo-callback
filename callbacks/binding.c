#include <stdlib.h>
#include <stdio.h>
#include "binding.h"

#include "_cgo_export.h"

cb_type stored_callback = NULL;
void *stored_data = NULL;

void set_callback(cb_type callback, void *data)
{
    stored_callback = callback;
    stored_data = data;
}

int call_callback(int argument)
{
    int ret = -1;

    if (stored_callback != NULL)
        ret = stored_callback(argument, stored_data);
    else
        printf("ERROR: No callback set\n");

    return ret;
}

int translator(int value, void *data)
{
    long key = (long)data;
    return go_callback_int(key, value);
}

void set_translator(long key)
{
    set_callback(translator, (void*)key);
}
