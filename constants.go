// Copyright 2019 Canonical Ltd.
// Licensed under the LGPLv3 with static-linking exception.
// See LICENCE file for details.

package tpm2

import (
	"math"
)

const (
	StartupClear StartupType = iota
	StartupState
)

const (
	OpEq         ArithmeticOp = 0x0000 // TPM_EO_EQ
	OpNeq        ArithmeticOp = 0x0001 // TPM_EO_NEQ
	OpSignedGT   ArithmeticOp = 0x0002 // TPM_EO_SIGNED_GT
	OpUnsignedGT ArithmeticOp = 0x0003 // TPM_EO_UNSIGNED_GT
	OpSignedLT   ArithmeticOp = 0x0004 // TPM_EO_SIGNED_LT
	OpUnsignedLT ArithmeticOp = 0x0005 // TPM_EO_UNSIGNED_LT
	OpSignedGE   ArithmeticOp = 0x0006 // TPM_EO_SIGNED_GE
	OpUnsignedGE ArithmeticOp = 0x0007 // TPM_EO_UNSIGNED_GE
	OpSignedLE   ArithmeticOp = 0x0008 // TPM_EO_SIGNED_LE
	OpUnsignedLE ArithmeticOp = 0x0009 // TPM_EO_UNSIGNED_LE
	OpBitset     ArithmeticOp = 0x000a // TPM_EO_BITSET
	OpBitclear   ArithmeticOp = 0x000b // TPM_EO_BITCLEAR
)

const (
	TagNoSessions         StructTag = 0x8001 // TPM_ST_NO_SESSIONS
	TagSessions           StructTag = 0x8002 // TPM_ST_SESSIONS
	TagAttestNV           StructTag = 0x8014 // TPM_ST_ATTEST_NV
	TagAttestCommandAudit StructTag = 0x8015 // TPM_ST_ATTEST_COMMAND_AUDIT
	TagAttestSessionAudit StructTag = 0x8016 // TPM_ST_ATTEST_SESSION_AUDIT
	TagAttestCertify      StructTag = 0x8017 // TPM_ST_ATTEST_CERTIFY
	TagAttestQuote        StructTag = 0x8018 // TPM_ST_ATTEST_QUOTE
	TagAttestTime         StructTag = 0x8019 // TPM_ST_ATTEST_TIME
	TagAttestCreation     StructTag = 0x801a // TPM_ST_ATTEST_CREATION
	TagCreation           StructTag = 0x8021 // TPM_ST_CREATION
	TagVerified           StructTag = 0x8022 // TPM_ST_VERIFIED
	TagAuthSecret         StructTag = 0x8023 // TPM_ST_AUTH_SECRET
	TagHashcheck          StructTag = 0x8024 // TPM_ST_HASHCHECK
	TagAuthSigned         StructTag = 0x8025 // TPM_ST_AUTH_SIGNED
)

const (
	TPMGeneratedValue TPMGenerated = 0xff544347 // TPM_GENERATED_VALUE
)

const (
	CommandFirst CommandCode = 0x0000011A

	CommandNVUndefineSpaceSpecial     CommandCode = 0x0000011F // TPM_CC_NV_UndefineSpaceSpecial
	CommandEvictControl               CommandCode = 0x00000120 // TPM_CC_EvictControl
	CommandNVUndefineSpace            CommandCode = 0x00000122 // TPM_CC_NV_UndefineSpace
	CommandClear                      CommandCode = 0x00000126 // TPM_CC_Clear
	CommandClearControl               CommandCode = 0x00000127 // TPM_CC_ClearControl
	CommandHierarchyChangeAuth        CommandCode = 0x00000129 // TPM_CC_HierarchyChangeAuth
	CommandNVDefineSpace              CommandCode = 0x0000012A // TPM_CC_NV_DefineSpace
	CommandCreatePrimary              CommandCode = 0x00000131 // TPM_CC_CreatePrimary
	CommandNVGlobalWriteLock          CommandCode = 0x00000132 // TPM_CC_NV_GlobalWriteLock
	CommandNVIncrement                CommandCode = 0x00000134 // TPM_CC_NV_Increment
	CommandNVSetBits                  CommandCode = 0x00000135 // TPM_CC_NV_SetBits
	CommandNVExtend                   CommandCode = 0x00000136 // TPM_CC_NV_Extend
	CommandNVWrite                    CommandCode = 0x00000137 // TPM_CC_NV_Write
	CommandNVWriteLock                CommandCode = 0x00000138 // TPM_CC_NV_WriteLock
	CommandDictionaryAttackLockReset  CommandCode = 0x00000139 // TPM_CC_DictionaryAttackLockReset
	CommandDictionaryAttackParameters CommandCode = 0x0000013A // TPM_CC_DictionaryAttackParameters
	CommandNVChangeAuth               CommandCode = 0x0000013B // TPM_CC_NV_ChangeAuth
	CommandPCREvent                   CommandCode = 0x0000013C // TPM_CC_PCR_Event
	CommandIncrementalSelfTest        CommandCode = 0x00000142 // TPM_CC_IncrementalSelfTest
	CommandSelfTest                   CommandCode = 0x00000143 // TPM_CC_SelfTest
	CommandStartup                    CommandCode = 0x00000144 // TPM_CC_Startup
	CommandShutdown                   CommandCode = 0x00000145 // TPM_CC_Shutdown
	CommandStirRandom                 CommandCode = 0x00000146 // TPM_CC_StirRandom
	CommandActivateCredential         CommandCode = 0x00000147 // TPM_CC_ActivateCredential
	CommandPolicyNV                   CommandCode = 0x00000149 // TPM_CC_PolicyNV
	CommandCertifyCreation            CommandCode = 0x0000014A // TPM_CC_CertifyCreation
	CommandNVRead                     CommandCode = 0x0000014E // TPM_CC_NV_Read
	CommandNVReadLock                 CommandCode = 0x0000014F // TPM_CC_NV_ReadLock
	CommandObjectChangeAuth           CommandCode = 0x00000150 // TPM_CC_ObjectChangeAuth
	CommandPolicySecret               CommandCode = 0x00000151 // TPM_CC_PolicySecret
	CommandCreate                     CommandCode = 0x00000153 // TPM_CC_Create
	CommandLoad                       CommandCode = 0x00000157 // TPM_CC_Load
	CommandSign                       CommandCode = 0x0000015D // TPM_CC_Sign
	CommandUnseal                     CommandCode = 0x0000015E // TPM_CC_Unseal
	CommandPolicySigned               CommandCode = 0x00000160 // TPM_CC_PolicySigned
	CommandContextLoad                CommandCode = 0x00000161 // TPM_CC_ContextLoad
	CommandContextSave                CommandCode = 0x00000162 // TPM_CC_ContextSave
	CommandFlushContext               CommandCode = 0x00000165 // TPM_CC_FlushContext
	CommandLoadExternal               CommandCode = 0x00000167 // TPM_CC_LoadExternal
	CommandMakeCredential             CommandCode = 0x00000168 // TPM_CC_MakeCredential
	CommandNVReadPublic               CommandCode = 0x00000169 // TPM_CC_NV_ReadPublic
	CommandPolicyAuthValue            CommandCode = 0x0000016B // TPM_CC_PolicyAuthValue
	CommandPolicyCommandCode          CommandCode = 0x0000016C // TPM_CC_PolicyCommandCode
	CommandPolicyOR                   CommandCode = 0x00000171 // TPM_CC_PolicyOR
	CommandPolicyTicket               CommandCode = 0x00000172 // TPM_CC_PolicyTicket
	CommandReadPublic                 CommandCode = 0x00000173 // TPM_CC_ReadPublic
	CommandStartAuthSession           CommandCode = 0x00000176 // TPM_CC_StartAuthSession
	CommandVerifySignature            CommandCode = 0x00000177 // TPM_CC_VerifySignature
	CommandGetCapability              CommandCode = 0x0000017A // TPM_CC_GetCapability
	CommandGetRandom                  CommandCode = 0x0000017B // TPM_CC_GetRandom
	CommandGetTestResult              CommandCode = 0x0000017C // TPM_CC_GetTestResult
	CommandPCRRead                    CommandCode = 0x0000017E // TPM_CC_PCR_Read
	CommandPolicyPCR                  CommandCode = 0x0000017F // TPM_CC_PolicyPCR
	CommandPolicyRestart              CommandCode = 0x00000180 // TPM_CC_PolicyRestart
	CommandReadClock                  CommandCode = 0x00000181 // TPM_CC_ReadClock
	CommandPCRExtend                  CommandCode = 0x00000182 // TPM_CC_PCR_Extend
	CommandPolicyGetDigest            CommandCode = 0x00000189 // TPM_CC_PolicyGetDigest
	CommandPolicyPassword             CommandCode = 0x0000018C // TPM_CC_PolicyPassword
	CommandCreateLoaded               CommandCode = 0x00000191 // TPM_CC_CreateLoaded
)

const (
	Success ResponseCode = 0
)

const (
	ErrorInitialize      ErrorCode0 = 0x00 // TPM_RC_INITIALIZE
	ErrorFailure         ErrorCode0 = 0x01 // TPM_RC_FAILURE
	ErrorSequence        ErrorCode0 = 0x03 // TPM_RC_SEQUENCE
	ErrorDisabled        ErrorCode0 = 0x20 // TPM_RC_DISABLED
	ErrorExclusive       ErrorCode0 = 0x21 // TPM_RC_EXCLUSIVE
	ErrorAuthType        ErrorCode0 = 0x24 // TPM_RC_AUTH_TYPE
	ErrorAuthMissing     ErrorCode0 = 0x25 // TPM_RC_AUTH_MISSING
	ErrorPolicy          ErrorCode0 = 0x26 // TPM_RC_POLICY
	ErrorPCR             ErrorCode0 = 0x27 // TPM_RC_PCR
	ErrorPCRChanged      ErrorCode0 = 0x28 // TPM_RC_PCR_CHANGED
	ErrorUpgrade         ErrorCode0 = 0x2d // TPM_RC_UPGRADE
	ErrorTooManyContexts ErrorCode0 = 0x2e // TPM_RC_TOO_MANY_CONTEXTS
	ErrorAuthUnavailable ErrorCode0 = 0x2f // TPM_RC_AUTH_UNAVAILABLE
	ErrorReboot          ErrorCode0 = 0x30 // TPM_RC_REBOOT
	ErrorUnbalanced      ErrorCode0 = 0x31 // TPM_RC_UNBALANCED
	ErrorCommandSize     ErrorCode0 = 0x42 // TPM_RC_COMMAND_SIZE
	ErrorCommandCode     ErrorCode0 = 0x43 // TPM_RC_COMMAND_CODE
	ErrorAuthsize        ErrorCode0 = 0x44 // TPM_RC_AUTHSIZE
	ErrorAuthContext     ErrorCode0 = 0x45 // TPM_RC_AUTH_CONTEXT
	ErrorNVRange         ErrorCode0 = 0x46 // TPM_RC_NV_RANGE
	ErrorNVSize          ErrorCode0 = 0x47 // TPM_RC_NV_SIZE
	ErrorNVLocked        ErrorCode0 = 0x48 // TPM_RC_NV_LOCKED
	ErrorNVAuthorization ErrorCode0 = 0x49 // TPM_RC_NV_AUTHORIZATION
	ErrorNVUninitialized ErrorCode0 = 0x4a // TPM_RC_NV_UNINITIALIZED
	ErrorNVSpace         ErrorCode0 = 0x4b // TPM_RC_NV_SPACE
	ErrorNVDefined       ErrorCode0 = 0x4c // TPM_RC_NV_DEFINED
	ErrorBadContext      ErrorCode0 = 0x50 // TPM_RC_BAD_CONTEXT
	ErrorCpHash          ErrorCode0 = 0x51 // TPM_RC_CPHASH
	ErrorParent          ErrorCode0 = 0x52 // TPM_RC_PARENT
	ErrorNeedsTest       ErrorCode0 = 0x53 // TPM_RC_NEEDS_TEST
	ErrorNoResult        ErrorCode0 = 0x54 // TPM_RC_NO_RESULT
	ErrorSensitive       ErrorCode0 = 0x55 // TPM_RC_SENSITIVE
)

const (
	ErrorAsymmetric   ErrorCode1 = 0x01 // TPM_RC_ASYMMETRIC
	ErrorAttributes   ErrorCode1 = 0x02 // TPM_RC_ATTRIBUTES
	ErrorHash         ErrorCode1 = 0x03 // TPM_RC_HASH
	ErrorValue        ErrorCode1 = 0x04 // TPM_RC_VALUE
	ErrorHierarchy    ErrorCode1 = 0x05 // TPM_RC_HIERARCHY
	ErrorKeySize      ErrorCode1 = 0x07 // TPM_RC_KEY_SIZE
	ErrorMGF          ErrorCode1 = 0x08 // TPM_RC_MGF
	ErrorMode         ErrorCode1 = 0x09 // TPM_RC_MODE
	ErrorType         ErrorCode1 = 0x0a // TPM_RC_TYPE
	ErrorHandle       ErrorCode1 = 0x0b // TPM_RC_HANDLE
	ErrorKDF          ErrorCode1 = 0x0c // TPM_RC_KDF
	ErrorRange        ErrorCode1 = 0x0d // TPM_RC_RANGE
	ErrorAuthFail     ErrorCode1 = 0x0e // TPM_RC_AUTH_FAIL
	ErrorNonce        ErrorCode1 = 0x0f // TPM_RC_NONCE
	ErrorPP           ErrorCode1 = 0x10 // TPM_RC_PP
	ErrorScheme       ErrorCode1 = 0x12 // TPM_RC_SCHEME
	ErrorSize         ErrorCode1 = 0x15 // TPM_RC_SIZE
	ErrorSymmetric    ErrorCode1 = 0x16 // TPM_RC_SYMMETRIC
	ErrorTag          ErrorCode1 = 0x17 // TPM_RC_TAG
	ErrorSelector     ErrorCode1 = 0x18 // TPM_RC_SELECTOR
	ErrorInsufficient ErrorCode1 = 0x1a // TPM_RC_INSUFFICIENT
	ErrorSignature    ErrorCode1 = 0x1b // TPM_RC_SIGNATURE
	ErrorKey          ErrorCode1 = 0x1c // TPM_RC_KEY
	ErrorPolicyFail   ErrorCode1 = 0x1d // TPM_RC_POLICY_FAIL
	ErrorIntegrity    ErrorCode1 = 0x1f // TPM_RC_INTEGRITY
	ErrorTicket       ErrorCode1 = 0x20 // TPM_RC_TICKET
	ErrorReservedBits ErrorCode1 = 0x21 // TPM_RC_RESERVED_BITS
	ErrorBadAuth      ErrorCode1 = 0x22 // TPM_RC_BAD_AUTH
	ErrorExpired      ErrorCode1 = 0x23 // TPM_RC_EXPIRED
	ErrorPolicyCC     ErrorCode1 = 0x24 // TPM_RC_POLICY_CC
	ErrorBinding      ErrorCode1 = 0x25 // TPM_RC_BINDING
	ErrorCurve        ErrorCode1 = 0x26 // TPM_RC_CURVE
	ErrorECCPoint     ErrorCode1 = 0x27 // TPM_RC_ECC_POINT
)

const (
	WarningContextGap     WarningCode = 0x01 // TPM_RC_CONTEXT_GAP
	WarningObjectMemory   WarningCode = 0x02 // TPM_RC_OBJECT_MEMORY
	WarningSessionMemory  WarningCode = 0x03 // TPM_RC_SESSION_MEMORY
	WarningMemory         WarningCode = 0x04 // TPM_RC_MEMORY
	WarningSessionHandles WarningCode = 0x05 // TPM_RC_SESSION_HANDLES
	WarningObjectHandles  WarningCode = 0x06 // TPM_RC_OBJECT_HANDLES
	WarningLocality       WarningCode = 0x07 // TPM_RC_LOCALITY
	WarningYielded        WarningCode = 0x08 // TPM_RC_YIELDED
	WarningCanceled       WarningCode = 0x09 // TPM_RC_CANCELED
	WarningTesting        WarningCode = 0x0a // TPM_RC_TESTING
	WarningReferenceH0    WarningCode = 0x10 // TPM_RC_REFERENCE_H0
	WarningReferenceH1    WarningCode = 0x11 // TPM_RC_REFERENCE_H1
	WarningReferenceH2    WarningCode = 0x12 // TPM_RC_REFERENCE_H2
	WarningReferenceH3    WarningCode = 0x13 // TPM_RC_REFERENCE_H3
	WarningReferenceH4    WarningCode = 0x14 // TPM_RC_REFERENCE_H4
	WarningReferenceH5    WarningCode = 0x15 // TPM_RC_REFERENCE_H5
	WarningReferenceH6    WarningCode = 0x16 // TPM_RC_REFERENCE_H6
	WarningReferenceS0    WarningCode = 0x18 // TPM_RC_REFERENCE_S0
	WarningReferenceS1    WarningCode = 0x19 // TPM_RC_REFERENCE_S1
	WarningReferenceS2    WarningCode = 0x1a // TPM_RC_REFERENCE_S2
	WarningReferenceS3    WarningCode = 0x1b // TPM_RC_REFERENCE_S3
	WarningReferenceS4    WarningCode = 0x1c // TPM_RC_REFERENCE_S4
	WarningReferenceS5    WarningCode = 0x1d // TPM_RC_REFERENCE_S5
	WarningReferenceS6    WarningCode = 0x1e // TPM_RC_REFERENCE_S6
	WarningNVRate         WarningCode = 0x20 // TPM_RC_NV_RATE
	WarningLockout        WarningCode = 0x21 // TPM_RC_LOCKOUT
	WarningRetry          WarningCode = 0x22 // TPM_RC_RETRY
	WarningNVUnavailable  WarningCode = 0x23 // TPM_RC_NV_UNAVAILABLE
)

const (
	HandleOwner       Handle = 0x40000001 // TPM_RH_OWNER
	HandleNull        Handle = 0x40000007 // TPM_RH_NULL
	HandleUnassigned  Handle = 0x40000008 // TPM_RH_UNASSIGNED
	HandlePW          Handle = 0x40000009 // TPM_RH_PW
	HandleLockout     Handle = 0x4000000a // TPM_RH_LOCKOUT
	HandleEndorsement Handle = 0x4000000b // TPM_RH_ENDORSEMENT
	HandlePlatform    Handle = 0x4000000c // TPM_RH_PLATFORM
	HandlePlatformNV  Handle = 0x4000000d // TPM_RH_PLATFORM_NV
)

const (
	HandleTypePCR              Handle = 0x00000000 // HR_PCR
	HandleTypeNVIndex          Handle = 0x01000000 // HR_NV_INDEX
	HandleTypeHMACSession      Handle = 0x02000000 // HR_HMAC_SESSION
	HandleTypePolicySession    Handle = 0x03000000 // HR_POLICY_SESSION
	HandleTypePermanent        Handle = 0x40000000 // HR_PERMANENT
	HandleTypeTransientObject  Handle = 0x80000000 // HR_TRANSIENT
	HandleTypePersistentObject Handle = 0x81000000 // HR_PERSISTENT

	HandleTypeLoadedSession Handle = 0x02000000
	HandleTypeActiveSession Handle = 0x03000000
)

const (
	AlgorithmRSA            AlgorithmId = 0x0001 // TPM_ALG_RSA
	AlgorithmSHA1           AlgorithmId = 0x0004 // TPM_ALG_SHA1
	AlgorithmHMAC           AlgorithmId = 0x0005 // TPM_ALG_HMAC
	AlgorithmAES            AlgorithmId = 0x0006 // TPM_ALG_AES
	AlgorithmMGF1           AlgorithmId = 0x0007 // TPM_ALG_MGF1
	AlgorithmKeyedHash      AlgorithmId = 0x0008 // TPM_ALG_KEYEDHASH
	AlgorithmXOR            AlgorithmId = 0x000a // TPM_ALG_XOR
	AlgorithmSHA256         AlgorithmId = 0x000b // TPM_ALG_SHA256
	AlgorithmSHA384         AlgorithmId = 0x000c // TPM_ALG_SHA384
	AlgorithmSHA512         AlgorithmId = 0x000d // TPM_ALG_SHA512
	AlgorithmNull           AlgorithmId = 0x0010 // TPM_ALG_NULL
	AlgorithmSM3_256        AlgorithmId = 0x0012 // TPM_ALG_SM3_256
	AlgorithmSM4            AlgorithmId = 0x0013 // TPM_ALG_SM4
	AlgorithmRSASSA         AlgorithmId = 0x0014 // TPM_ALG_RSASSA
	AlgorithmRSAES          AlgorithmId = 0x0015 // TPM_ALG_RSAES
	AlgorithmRSAPSS         AlgorithmId = 0x0016 // TPM_ALG_RSAPSS
	AlgorithmOAEP           AlgorithmId = 0x0017 // TPM_ALG_OAEP
	AlgorithmECDSA          AlgorithmId = 0x0018 // TPM_ALG_ECDSA
	AlgorithmECDH           AlgorithmId = 0x0019 // TPM_ALG_ECDH
	AlgorithmECDAA          AlgorithmId = 0x001a // TPM_ALG_ECDAA
	AlgorithmSM2            AlgorithmId = 0x001b // TPM_ALG_SM2
	AlgorithmECSCHNORR      AlgorithmId = 0x001c // TPM_ALG_ECSCHNORR
	AlgorithmECMQV          AlgorithmId = 0x001d // TPM_ALG_ECMQV
	AlgorithmKDF1_SP800_56A AlgorithmId = 0x0020 // TPM_ALG_KDF1_SP800_56A
	AlgorithmKDF2           AlgorithmId = 0x0021 // TPM_ALG_KDF2
	AlgorithmKDF1_SP800_108 AlgorithmId = 0x0022 // TPM_ALG_KDF1_SP800_108
	AlgorithmECC            AlgorithmId = 0x0023 // TPM_ALG_ECC
	AlgorithmSymCipher      AlgorithmId = 0x0025 // TPM_ALG_SYMCIPHER
	AlgorithmCamellia       AlgorithmId = 0x0026 // TPM_ALG_CAMELLIA
	AlgorithmCTR            AlgorithmId = 0x0040 // TPM_ALG_CTR
	AlgorithmOFB            AlgorithmId = 0x0041 // TPM_ALG_OFB
	AlgorithmCBC            AlgorithmId = 0x0042 // TPM_ALG_CBC
	AlgorithmCFB            AlgorithmId = 0x0043 // TPM_ALG_CFB
	AlgorithmECB            AlgorithmId = 0x0044 // TPM_ALG_ECB

	AlgorithmFirst AlgorithmId = AlgorithmRSA
)

const (
	AttrFixedTPM             ObjectAttributes = 1 << 1  // fixedTPM
	AttrStClear              ObjectAttributes = 1 << 2  // stClear
	AttrFixedParent          ObjectAttributes = 1 << 4  // fixedParent
	AttrSensitiveDataOrigin  ObjectAttributes = 1 << 5  // sensitiveDataOrigin
	AttrUserWithAuth         ObjectAttributes = 1 << 6  // userWithAuth
	AttrAdminWithPolicy      ObjectAttributes = 1 << 7  // adminWithPolicy
	AttrNoDA                 ObjectAttributes = 1 << 10 // noDA
	AttrEncryptedDuplication ObjectAttributes = 1 << 11 // encryptedDuplication
	AttrRestricted           ObjectAttributes = 1 << 16 // restricted
	AttrDecrypt              ObjectAttributes = 1 << 17 // decrypt
	AttrSign                 ObjectAttributes = 1 << 18 // sign
)

const (
	AttrNVPPWrite        NVAttributes = 1 << 0  // TPMA_NV_PPWRITE
	AttrNVOwnerWrite     NVAttributes = 1 << 1  // TPMA_NV_OWNERWRITE
	AttrNVAuthWrite      NVAttributes = 1 << 2  // TPMA_NV_AUTHWRITE
	AttrNVPolicyWrite    NVAttributes = 1 << 3  // TPMA_NV_POLICY_RITE
	AttrNVPolicyDelete   NVAttributes = 1 << 10 // TPMA_NV_POLICY_DELETE
	AttrNVWriteLocked    NVAttributes = 1 << 11 // TPMA_NV_WRITELOCKED
	AttrNVWriteAll       NVAttributes = 1 << 12 // TPMA_NV_WRITEALL
	AttrNVWriteDefine    NVAttributes = 1 << 13 // TPMA_NV_WRITEDEFINE
	AttrNVWriteStClear   NVAttributes = 1 << 14 // TPMA_NV_WRITE_STCLEAR
	AttrNVGlobalLock     NVAttributes = 1 << 15 // TPMA_NV_GLOBALLOCK
	AttrNVPPRead         NVAttributes = 1 << 16 // TPMA_NV_PPREAD
	AttrNVOwnerRead      NVAttributes = 1 << 17 // TPMA_NV_OWNERREAD
	AttrNVAuthRead       NVAttributes = 1 << 18 // TPMA_NV_AUTHREAD
	AttrNVPolicyRead     NVAttributes = 1 << 19 // TPMA_NV_POLICYREAD
	AttrNVNoDA           NVAttributes = 1 << 25 // TPMA_NV_NO_DA
	AttrNVOrderly        NVAttributes = 1 << 26 // TPMA_NV_ORDERLY
	AttrNVClearStClear   NVAttributes = 1 << 27 // TPMA_NV_CLEAR_STCLEAR
	AttrNVReadLocked     NVAttributes = 1 << 28 // TPMA_NV_READLOCKED
	AttrNVWritten        NVAttributes = 1 << 29 // TPMA_NV_WRITTEN
	AttrNVPlatformCreate NVAttributes = 1 << 30 // TPMA_NV_PLATFORMCREATE
	AttrNVReadStClear    NVAttributes = 1 << 31 // TPMA_NV_READ_STCLEAR
)

const (
	NVTypeOrdinary NVType = 0 // TPM_NT_ORDINARY
	NVTypeCounter  NVType = 1 // TPM_NT_COUNTER
	NVTypeBits     NVType = 2 // TPM_NT_BITS
	NVTypeExtend   NVType = 4 // TPM_NT_EXTEND
	NVTypePinFail  NVType = 8 // TPM_NT_PIN_FAIL
	NVTypePinPass  NVType = 9 // TPM_NT_PIN_PASS
)

const (
	LocalityZero  Locality = 0 // TPM_LOC_ZERO
	LocalityOne   Locality = 1 // TPM_LOC_ONE
	LocalityTwo   Locality = 2 // TPM_LOC_TWO
	LocalityThree Locality = 3 // TPM_LOC_THREE
	LocalityFour  Locality = 4 // TPM_LOC_FOUR
)

const (
	CapabilityAlgs          Capability = 0 // TPM_CAP_ALGS
	CapabilityHandles       Capability = 1 // TPM_CAP_HANDLES
	CapabilityCommands      Capability = 2 // TPM_CAP_COMMANDS
	CapabilityPPCommands    Capability = 3 // TPM_CAP_PP_COMMANDS
	CapabilityAuditCommands Capability = 4 // TPM_CAP_AUDIT_COMMANDS
	CapabilityPCRs          Capability = 5 // TPM_CAP_PCRS
	CapabilityTPMProperties Capability = 6 // TPM_CAP_TPM_PROPERTIES
	CapabilityPCRProperties Capability = 7 // TPM_CAP_PCR_PROPERTIES
	CapabilityECCCurves     Capability = 8 // TPM_CAP_ECC_CURVES
	CapabilityAuthPolicies  Capability = 9 // TPM_CAP_AUTH_POLICIES
)

const (
	CapabilityMaxProperties uint32 = math.MaxUint32
)

const (
	// These constants represent properties that only change when the firmware in the TPM changes.
	PropertyFamilyIndicator   Property = 0x100 // TPM_PT_FAMILY_INDICATOR
	PropertyLevel             Property = 0x101 // TPM_PT_LEVEL
	PropertyRevision          Property = 0x102 // TPM_PT_REVISION
	PropertyDayOfYear         Property = 0x103 // TPM_PT_DAY_OF_YEAR
	PropertyYear              Property = 0x104 // TPM_PT_YEAR
	PropertyManufacturer      Property = 0x105 // TPM_PT_MANUFACTURER
	PropertyVendorString1     Property = 0x106 // TPM_PT_VENDOR_STRING_1
	PropertyVendorString2     Property = 0x107 // TPM_PT_VENDOR_STRING_2
	PropertyVendorString3     Property = 0x108 // TPM_PT_VENDOR_STRING_3
	PropertyVendorString4     Property = 0x109 // TPM_PT_VENDOR_STRING_4
	PropertyVendorTPMType     Property = 0x10a // TPM_PT_VENDOR_TPM_TYPE
	PropertyFirmwareVersion1  Property = 0x10b // TPM_PT_FIRMWARE_VERSION_1
	PropertyFirmwareVersion2  Property = 0x10c // TPM_PT_FIRMWARE_VERSION_2
	PropertyInputBuffer       Property = 0x10d // TPM_PT_INPUT_BUFFER
	PropertyHRTransientMin    Property = 0x10e // TPM_PT_HR_TRANSIENT_MIN
	PropertyHRPersistentMin   Property = 0x10f // TPM_PT_HR_PERSISTENT_MIN
	PropertyHRLoadedMin       Property = 0x110 // TPM_PT_HR_LOADED_MIN
	PropertyActiveSessionsMax Property = 0x111 // TPM_PT_ACTIVE_SESSIONS_MAX
	PropertyPCRCount          Property = 0x112 // TPM_PT_PCR_COUNT
	PropertyPCRSelectMin      Property = 0x113 // TPM_PT_PCR_SELECT_MIN
	PropertyContextGapMax     Property = 0x114 // TPM_PT_CONTEXT_GAP_MAX
	PropertyNVCountersMax     Property = 0x116 // TPM_PT_NV_COUNTERS_MAX
	PropertyNVIndexMax        Property = 0x117 // TPM_PT_NV_INDEX_MAX
	PropertyMemory            Property = 0x118 // TPM_PT_MEMORY
	PropertyClockUpdate       Property = 0x119 // TPM_PT_CLOCK_UPDATE
	PropertyContextHash       Property = 0x11a // TPM_PT_CONTEXT_HASH
	PropertyContextSym        Property = 0x11b // TPM_PT_CONTEXT_SYM
	PropertyContextSymSize    Property = 0x11c // TPM_PT_CONTEXT_SYM_SIZE
	PropertyOrderlyCount      Property = 0x11d // TPM_PT_ORDERLY_COUNT
	PropertyMaxCommandSize    Property = 0x11e // TPM_PT_MAX_COMMAND_SIZE
	PropertyMaxResponseSize   Property = 0x11f // TPM_PT_MAX_RESPONSE_SIZE
	PropertyMaxDigest         Property = 0x120 // TPM_PT_MAX_DIGEST
	PropertyMaxObjectContext  Property = 0x121 // TPM_PT_MAX_OBJECT_CONTEXT
	PropertyMaxSessionContext Property = 0x122 // TPM_PT_MAX_SESSION_CONTEXT
	PropertyPSFamilyIndicator Property = 0x123 // TPM_PT_PS_FAMILY_INDICATOR
	PropertyPSLevel           Property = 0x124 // TPM_PT_PS_LEVEL
	PropertyPSRevision        Property = 0x125 // TPM_PT_PS_REVISION
	PropertyPSDayOfYear       Property = 0x126 // TPM_PT_PS_DAY_OF_YEAR
	PropertyPSYear            Property = 0x127 // TPM_PT_PS_YEAR
	PropertySplitMax          Property = 0x128 // TPM_PT_SPLIT_MAX
	PropertyTotalCommands     Property = 0x129 // TPM_PT_TOTAL_COMMANDS
	PropertyLibraryCommands   Property = 0x12a // TPM_PT_LIBRARY_COMMANDS
	PropertyVendorCommands    Property = 0x12b // TPM_PT_VENDOR_COMMANDS
	PropertyNVBufferMax       Property = 0x12c // TPM_PT_NV_BUFFER_MAX
	PropertyModes             Property = 0x12d // TPM_PT_MODES
	PropertyMaxCapBuffer      Property = 0x12e // TPM_PT_MAX_CAP_BUFFER

	PropertyFixed Property = PropertyFamilyIndicator
)

const (
	// These constants represent properties that change for reasons other than a firmware upgrade. Some of
	// them may not persist across power cycles.
	PropertyPermanent         Property = 0x200 // TPM_PT_PERMANENT
	PropertyStartupClear      Property = 0x201 // TPM_PT_STARTUP_CLEAR
	PropertyHRNVIndex         Property = 0x202 // TPM_PT_HR_NV_INDEX
	PropertyHRLoaded          Property = 0x203 // TPM_PT_HR_LOADED
	PropertyHRLoadedAvail     Property = 0x204 // TPM_PT_HR_LOADED_AVAIL
	PropertyHRActive          Property = 0x205 // TPM_PT_HR_ACTIVE
	PropertyHRActiveAvail     Property = 0x206 // TPM_PT_HR_ACTIVE_AVAIL
	PropertyHRTransientAvail  Property = 0x207 // TPM_PT_HR_TRANSIENT_AVAIL
	PropertyHRPersistent      Property = 0x208 // TPM_PT_HR_PERSISTENT
	PropertyHRPersistentAvail Property = 0x209 // TPM_PT_HR_PERSISTENT_AVAIL
	PropertyNVCounters        Property = 0x20a // TPM_PT_NV_COUNTERS
	PropertyNVCountersAvail   Property = 0x20b // TPM_PT_NV_COUNTERS_AVAIL
	PropertyAlgorithmSet      Property = 0x20c // TPM_PT_ALGORITHM_SET
	PropertyLoadedCurves      Property = 0x20d // TPM_PT_LOADED_CURVES
	PropertyLockoutCounter    Property = 0x20e // TPM_PT_LOCKOUT_COUNTER
	PropertyMaxAuthFail       Property = 0x20f // TPM_PT_MAX_AUTH_FAIL
	PropertyLockoutInterval   Property = 0x210 // TPM_PT_LOCKOUT_INTERVAL
	PropertyLockoutRecovery   Property = 0x211 // TPM_PT_LOCKOUT_RECOVERY
	PropertyNVWriteRecovery   Property = 0x212 // TPM_PT_NV_WRITE_RECOVERY
	PropertyAuditCounter0     Property = 0x213 // TPM_PT_AUDIT_COUNTER_0
	PropertyAuditCounter1     Property = 0x214 // TPM_PT_AUDIT_COUNTER_1

	PropertyVar Property = PropertyPermanent
)

const (
	PropertyPCRSave        PropertyPCR = 0x00 // TPM_PT_PCR_SAVE
	PropertyPCRExtendL0    PropertyPCR = 0x01 // TPM_PT_PCR_EXTEND_L0
	PropertyPCRResetL0     PropertyPCR = 0x02 // TPM_PT_PCR_RESET_L0
	PropertyPCRExtendL1    PropertyPCR = 0x03 // TPM_PT_PCR_EXTEND_L1
	PropertyPCRResetL1     PropertyPCR = 0x04 // TPM_PT_PCR_RESET_L1
	PropertyPCRExtendL2    PropertyPCR = 0x05 // TPM_PT_PCR_EXTEND_L2
	PropertyPCRResetL2     PropertyPCR = 0x06 // TPM_PT_PCR_RESET_L2
	PropertyPCRExtendL3    PropertyPCR = 0x07 // TPM_PT_PCR_EXTEND_L3
	PropertyPCRResetL3     PropertyPCR = 0x08 // TPM_PT_PCR_RESET_L3
	PropertyPCRExtendL4    PropertyPCR = 0x09 // TPM_PT_PCR_EXTEND_L4
	PropertyPCRResetL4     PropertyPCR = 0x0a // TPM_PT_PCR_RESET_L4
	PropertyPCRNoIncrement PropertyPCR = 0x11 // TPM_PT_PCR_NO_INCREMENT
	PropertyPCRDRTMReset   PropertyPCR = 0x12 // TPM_PT_PCR_DRTM_RESET
	PropertyPCRPolicy      PropertyPCR = 0x13 // TPM_PT_PCR_POLICY
	PropertyPCRAuth        PropertyPCR = 0x14 // TPM_PT_PCR_AUTH

	PropertyPCRFirst PropertyPCR = PropertyPCRSave
)

const (
	AttrAsymmetric AlgorithmAttributes = 1 << 0
	AttrSymmetric  AlgorithmAttributes = 1 << 1
	AttrHash       AlgorithmAttributes = 1 << 2
	AttrObject     AlgorithmAttributes = 1 << 3
	AttrSigning    AlgorithmAttributes = 1 << 8
	AttrEncrypting AlgorithmAttributes = 1 << 9
	AttrMethod     AlgorithmAttributes = 1 << 10
)

const (
	AttrNV        CommandAttributes = 1 << 22
	AttrExtensive CommandAttributes = 1 << 23
	AttrFlushed   CommandAttributes = 1 << 24
	AttrRHandle   CommandAttributes = 1 << 28
	AttrV         CommandAttributes = 1 << 29
)

const (
	ECCCurveNIST_P192 ECCCurve = 0x0001 // TPM_ECC_NIST_P192
	ECCCurveNIST_P224 ECCCurve = 0x0002 // TPM_ECC_NIST_P224
	ECCCurveNIST_P256 ECCCurve = 0x0003 // TPM_ECC_NIST_P256
	ECCCurveNIST_P384 ECCCurve = 0x0004 // TPM_ECC_NIST_P384
	ECCCurveNIST_P521 ECCCurve = 0x0005 // TPM_ECC_NIST_P521
	ECCCurveBN_P256   ECCCurve = 0x0010 // TPM_ECC_BN_P256
	ECCCurveBN_P638   ECCCurve = 0x0011 // TPM_ECC_BN_P638
	ECCCurveSM2_P256  ECCCurve = 0x0020 // TPM_ECC_SM2_P256

	ECCCurveFirst ECCCurve = ECCCurveNIST_P192
)

const (
	SessionTypeHMAC   SessionType = 0x00 // TPM_SE_HMAC
	SessionTypePolicy SessionType = 0x01 // TPM_SE_POLICY
	SessionTypeTrial  SessionType = 0x03 // TPM_SE_TRIAL
)

const (
	AttrOwnerAuthSet       PermanentAttributes = 1 << 0  // ownerAuthSet
	AttrEndorsementAuthSet PermanentAttributes = 1 << 1  // endorsementAuthSet
	AttrLockoutAuthSet     PermanentAttributes = 1 << 2  // lockoutAuthSet
	AttrDisableClear       PermanentAttributes = 1 << 8  // disableClear
	AttrInLockout          PermanentAttributes = 1 << 9  // inLockout
	AttrTPMGeneratedEPS    PermanentAttributes = 1 << 10 // tpmGeneratedEPS
)

const (
	AttrPhEnable   StartupClearAttributes = 1 << 0  // phEnable
	AttrShEnable   StartupClearAttributes = 1 << 1  // shEnable
	AttrEhEnable   StartupClearAttributes = 1 << 2  // ehEnable
	AttrPhEnableNV StartupClearAttributes = 1 << 3  // phEnableNV
	AttrOrderly    StartupClearAttributes = 1 << 31 // orderly
)
