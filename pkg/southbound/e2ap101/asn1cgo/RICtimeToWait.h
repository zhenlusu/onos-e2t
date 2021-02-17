/*
 * Generated by asn1c-0.9.29 (http://lionet.info/asn1c)
 * From ASN.1 module "E2AP-IEs"
 * 	found in "e2ap-v01.01.00.asn1"
 * 	`asn1c -fcompound-names -fincludes-quoted -fno-include-deps -findirect-choice -gen-PER -no-gen-OER -D.`
 */

#ifndef	_RICtimeToWait_H_
#define	_RICtimeToWait_H_


#include "asn_application.h"

/* Including external dependencies */
#include "NativeEnumerated.h"

#ifdef __cplusplus
extern "C" {
#endif

/* Dependencies */
typedef enum RICtimeToWait {
	RICtimeToWait_zero	= 0,
	RICtimeToWait_w1ms	= 1,
	RICtimeToWait_w2ms	= 2,
	RICtimeToWait_w5ms	= 3,
	RICtimeToWait_w10ms	= 4,
	RICtimeToWait_w20ms	= 5,
	RICtimeToWait_w30ms	= 6,
	RICtimeToWait_w40ms	= 7,
	RICtimeToWait_w50ms	= 8,
	RICtimeToWait_w100ms	= 9,
	RICtimeToWait_w200ms	= 10,
	RICtimeToWait_w500ms	= 11,
	RICtimeToWait_w1s	= 12,
	RICtimeToWait_w2s	= 13,
	RICtimeToWait_w5s	= 14,
	RICtimeToWait_w10s	= 15,
	RICtimeToWait_w20s	= 16,
	RICtimeToWait_w60s	= 17
	/*
	 * Enumeration is extensible
	 */
} e_RICtimeToWait;

/* RICtimeToWait */
typedef long	 RICtimeToWait_t;

/* Implementation */
extern asn_per_constraints_t asn_PER_type_RICtimeToWait_constr_1;
extern asn_TYPE_descriptor_t asn_DEF_RICtimeToWait;
extern const asn_INTEGER_specifics_t asn_SPC_RICtimeToWait_specs_1;
asn_struct_free_f RICtimeToWait_free;
asn_struct_print_f RICtimeToWait_print;
asn_constr_check_f RICtimeToWait_constraint;
ber_type_decoder_f RICtimeToWait_decode_ber;
der_type_encoder_f RICtimeToWait_encode_der;
xer_type_decoder_f RICtimeToWait_decode_xer;
xer_type_encoder_f RICtimeToWait_encode_xer;
per_type_decoder_f RICtimeToWait_decode_uper;
per_type_encoder_f RICtimeToWait_encode_uper;
per_type_decoder_f RICtimeToWait_decode_aper;
per_type_encoder_f RICtimeToWait_encode_aper;

#ifdef __cplusplus
}
#endif

#endif	/* _RICtimeToWait_H_ */
#include "asn_internal.h"
