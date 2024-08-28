#ifndef _H_LIB_YOYO_ACTIVATE_H_
#define _H_LIB_YOYO_ACTIVATE_H_

#ifdef YOYO_ACTIVATE_EXPORT
#define YOYOACTIVATE _declspec(dllexport)
#else
#define YOYOACTIVATE
#endif

#define STATUS_ACTIVATE_SUCCESS (0)
#define STATUS_ACTIVATE_ERROR (1)

#ifdef __cplusplus
extern "C" {
#endif
	/*
	generate the unique machine code
	@machine_code output which length is 1024
	*/
	YOYOACTIVATE char* yoyo_machine_code();

	/*
	check the saved key validation
	@userSnCode the device code
	*/
	YOYOACTIVATE int yoyo_activate_key(const char* keyCode, const char* suffix);

	YOYOACTIVATE char* get_activate_version();

	YOYOACTIVATE const char *get_query();

	#ifdef _WIN32
    YOYOACTIVATE char *get_sn_query();
    
	YOYOACTIVATE char *get_mac_query();
	#endif

#ifdef __cplusplus
}
#endif

#endif
