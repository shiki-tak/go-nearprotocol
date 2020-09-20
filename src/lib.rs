mod standalone;

use libc;
use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn greet(name: *const libc::c_char) {
    let buf_name = unsafe { CStr::from_ptr(name).to_bytes() };
    let str_name = String::from_utf8(buf_name.to_vec()).unwrap();
    println!("Hello, {}!", str_name);
}

fn arg_warapper(param: *const libc::c_char) -> Option<String> {
    let buf_name = unsafe { CStr::from_ptr(param).to_bytes() };
    let param_str = String::from_utf8(buf_name.to_vec()).unwrap();
    if param_str == "" {
        return None;
    }
    return Some(param_str)
}

#[no_mangle]
pub extern "C" fn run_with_standalone(vm_kind_c: *const libc::c_char, context_c: *const libc::c_char,
    context_file_c: *const libc::c_char, input_c: *const libc::c_char,
    state_c: *const libc::c_char, method_name_c: *const libc::c_char, config_c: *const libc::c_char,
    config_file_c: *const libc::c_char, wasm_file_c: *const libc::c_char,) {

        let vm_kind = arg_warapper(vm_kind_c);
        let context = arg_warapper(context_c);
        let context_file = arg_warapper(context_file_c);
        let input = arg_warapper(input_c);
        let state = arg_warapper(state_c);
        let method_name = arg_warapper(method_name_c);
        let config = arg_warapper(config_c);
        let config_file = arg_warapper(config_file_c);
        let wasm_file = arg_warapper(wasm_file_c);

        standalone::run(vm_kind, context, 
            context_file, input, state,
            method_name, config, config_file,
            wasm_file);
}

// pub fn run(vm_kind: Option<String>
//     context: Option<String>
//     context_file: Option<String>
//     input: Option<String>
//     state: Option<String>
//     method_name: Option<String>
//     config: Option<String>
//     config_file: Option<String>
//     wasm_file: Option<String>)
