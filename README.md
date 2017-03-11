# cgo-callback
Simple library showing how to call Go functions from cgo callbacks.

The cgo command is a handy tool allowing to call C code easily and to
export Go functions to be called from C with the export directive.
But there are times that you need to pass a Go function during the
runtime to be called later. And due to the managed nature of Go
pointers, it is not possible to pass a function pointer directly to C.

Based on this[1] Stack overflow answer, one option is to use some
kind of "id" that can be passed to the C callback for each Go callback,
and when it is called, the id is used to retrieve the original Go
function. This example shows how to achieve this.

[1] http://stackoverflow.com/a/37641451/347889
