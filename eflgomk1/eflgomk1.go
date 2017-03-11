package main

/*
#cgo CFLAGS: -DEFL_BETA_API_SUPPORT=1 -DEFL_EO_API_SUPPORT=1
#cgo pkg-config: elementary
#include <Elementary.h>

// efl_add is a macro and cgo has no CPP support
static inline Eo* _efl_new(const Efl_Class *klass, Eo* parent)
{
    return efl_add(klass, parent);
}

*/
import "C"
import "unsafe"

func main() {
    C.elm_init(0, nil);

    title := C.CString("Hello, World!");
    defer C.free(unsafe.Pointer(title));
    message := C.CString("Click me!");
    defer C.free(unsafe.Pointer(message));

    win := C._efl_new(C.efl_ui_win_standard_class_get(), nil);

    C.elm_policy_set(C.ELM_POLICY_QUIT, C.ELM_POLICY_QUIT_LAST_WINDOW_CLOSED);

    C.efl_ui_win_autodel_set(win, 1);
    C.efl_text_set(win, title);

    box := C._efl_new(C.efl_ui_box_class_get(), win);
    C.efl_content_set(win, box);
    C.efl_gfx_size_set(box, 640, 480);
    C.efl_gfx_size_hint_weight_set(box, 1, 1);
    C.efl_gfx_visible_set(box, 1);

    button := C._efl_new(C.elm_button_class_get(), box);
    C.efl_pack(box, button);
    C.efl_gfx_size_hint_weight_set(button, 0, 0);
    C.efl_text_set(button, message);
    C.efl_gfx_visible_set(button, 1);

    C.efl_gfx_size_set(win, 640, 480);
    C.efl_gfx_visible_set(win, 1);

    loop := C.efl_loop_main_get(C.efl_loop_class_get());
    C.efl_loop_begin(loop);

}

