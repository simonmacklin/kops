// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAlreadyExistsException for service response error code
	// "AlreadyExistsException".
	//
	// The request was rejected because it attempted to create a resource that already
	// exists.
	ErrCodeAlreadyExistsException = "AlreadyExistsException"

	// ErrCodeCloudHsmClusterInUseException for service response error code
	// "CloudHsmClusterInUseException".
	//
	// The request was rejected because the specified CloudHSM cluster is already
	// associated with an CloudHSM key store in the account, or it shares a backup
	// history with an CloudHSM key store in the account. Each CloudHSM key store
	// in the account must be associated with a different CloudHSM cluster.
	//
	// CloudHSM clusters that share a backup history have the same cluster certificate.
	// To view the cluster certificate of an CloudHSM cluster, use the DescribeClusters
	// (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
	// operation.
	ErrCodeCloudHsmClusterInUseException = "CloudHsmClusterInUseException"

	// ErrCodeCloudHsmClusterInvalidConfigurationException for service response error code
	// "CloudHsmClusterInvalidConfigurationException".
	//
	// The request was rejected because the associated CloudHSM cluster did not
	// meet the configuration requirements for an CloudHSM key store.
	//
	//    * The CloudHSM cluster must be configured with private subnets in at least
	//    two different Availability Zones in the Region.
	//
	//    * The security group for the cluster (https://docs.aws.amazon.com/cloudhsm/latest/userguide/configure-sg.html)
	//    (cloudhsm-cluster-<cluster-id>-sg) must include inbound rules and outbound
	//    rules that allow TCP traffic on ports 2223-2225. The Source in the inbound
	//    rules and the Destination in the outbound rules must match the security
	//    group ID. These rules are set by default when you create the CloudHSM
	//    cluster. Do not delete or change them. To get information about a particular
	//    security group, use the DescribeSecurityGroups (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSecurityGroups.html)
	//    operation.
	//
	//    * The CloudHSM cluster must contain at least as many HSMs as the operation
	//    requires. To add HSMs, use the CloudHSM CreateHsm (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_CreateHsm.html)
	//    operation. For the CreateCustomKeyStore, UpdateCustomKeyStore, and CreateKey
	//    operations, the CloudHSM cluster must have at least two active HSMs, each
	//    in a different Availability Zone. For the ConnectCustomKeyStore operation,
	//    the CloudHSM must contain at least one active HSM.
	//
	// For information about the requirements for an CloudHSM cluster that is associated
	// with an CloudHSM key store, see Assemble the Prerequisites (https://docs.aws.amazon.com/kms/latest/developerguide/create-keystore.html#before-keystore)
	// in the Key Management Service Developer Guide. For information about creating
	// a private subnet for an CloudHSM cluster, see Create a Private Subnet (https://docs.aws.amazon.com/cloudhsm/latest/userguide/create-subnets.html)
	// in the CloudHSM User Guide. For information about cluster security groups,
	// see Configure a Default Security Group (https://docs.aws.amazon.com/cloudhsm/latest/userguide/configure-sg.html)
	// in the CloudHSM User Guide .
	ErrCodeCloudHsmClusterInvalidConfigurationException = "CloudHsmClusterInvalidConfigurationException"

	// ErrCodeCloudHsmClusterNotActiveException for service response error code
	// "CloudHsmClusterNotActiveException".
	//
	// The request was rejected because the CloudHSM cluster associated with the
	// CloudHSM key store is not active. Initialize and activate the cluster and
	// try the command again. For detailed instructions, see Getting Started (https://docs.aws.amazon.com/cloudhsm/latest/userguide/getting-started.html)
	// in the CloudHSM User Guide.
	ErrCodeCloudHsmClusterNotActiveException = "CloudHsmClusterNotActiveException"

	// ErrCodeCloudHsmClusterNotFoundException for service response error code
	// "CloudHsmClusterNotFoundException".
	//
	// The request was rejected because KMS cannot find the CloudHSM cluster with
	// the specified cluster ID. Retry the request with a different cluster ID.
	ErrCodeCloudHsmClusterNotFoundException = "CloudHsmClusterNotFoundException"

	// ErrCodeCloudHsmClusterNotRelatedException for service response error code
	// "CloudHsmClusterNotRelatedException".
	//
	// The request was rejected because the specified CloudHSM cluster has a different
	// cluster certificate than the original cluster. You cannot use the operation
	// to specify an unrelated cluster for an CloudHSM key store.
	//
	// Specify an CloudHSM cluster that shares a backup history with the original
	// cluster. This includes clusters that were created from a backup of the current
	// cluster, and clusters that were created from the same backup that produced
	// the current cluster.
	//
	// CloudHSM clusters that share a backup history have the same cluster certificate.
	// To view the cluster certificate of an CloudHSM cluster, use the DescribeClusters
	// (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/API_DescribeClusters.html)
	// operation.
	ErrCodeCloudHsmClusterNotRelatedException = "CloudHsmClusterNotRelatedException"

	// ErrCodeCustomKeyStoreHasCMKsException for service response error code
	// "CustomKeyStoreHasCMKsException".
	//
	// The request was rejected because the custom key store contains KMS keys.
	// After verifying that you do not need to use the KMS keys, use the ScheduleKeyDeletion
	// operation to delete the KMS keys. After they are deleted, you can delete
	// the custom key store.
	ErrCodeCustomKeyStoreHasCMKsException = "CustomKeyStoreHasCMKsException"

	// ErrCodeCustomKeyStoreInvalidStateException for service response error code
	// "CustomKeyStoreInvalidStateException".
	//
	// The request was rejected because of the ConnectionState of the custom key
	// store. To get the ConnectionState of a custom key store, use the DescribeCustomKeyStores
	// operation.
	//
	// This exception is thrown under the following conditions:
	//
	//    * You requested the ConnectCustomKeyStore operation on a custom key store
	//    with a ConnectionState of DISCONNECTING or FAILED. This operation is valid
	//    for all other ConnectionState values. To reconnect a custom key store
	//    in a FAILED state, disconnect it (DisconnectCustomKeyStore), then connect
	//    it (ConnectCustomKeyStore).
	//
	//    * You requested the CreateKey operation in a custom key store that is
	//    not connected. This operations is valid only when the custom key store
	//    ConnectionState is CONNECTED.
	//
	//    * You requested the DisconnectCustomKeyStore operation on a custom key
	//    store with a ConnectionState of DISCONNECTING or DISCONNECTED. This operation
	//    is valid for all other ConnectionState values.
	//
	//    * You requested the UpdateCustomKeyStore or DeleteCustomKeyStore operation
	//    on a custom key store that is not disconnected. This operation is valid
	//    only when the custom key store ConnectionState is DISCONNECTED.
	//
	//    * You requested the GenerateRandom operation in an CloudHSM key store
	//    that is not connected. This operation is valid only when the CloudHSM
	//    key store ConnectionState is CONNECTED.
	ErrCodeCustomKeyStoreInvalidStateException = "CustomKeyStoreInvalidStateException"

	// ErrCodeCustomKeyStoreNameInUseException for service response error code
	// "CustomKeyStoreNameInUseException".
	//
	// The request was rejected because the specified custom key store name is already
	// assigned to another custom key store in the account. Try again with a custom
	// key store name that is unique in the account.
	ErrCodeCustomKeyStoreNameInUseException = "CustomKeyStoreNameInUseException"

	// ErrCodeCustomKeyStoreNotFoundException for service response error code
	// "CustomKeyStoreNotFoundException".
	//
	// The request was rejected because KMS cannot find a custom key store with
	// the specified key store name or ID.
	ErrCodeCustomKeyStoreNotFoundException = "CustomKeyStoreNotFoundException"

	// ErrCodeDependencyTimeoutException for service response error code
	// "DependencyTimeoutException".
	//
	// The system timed out while trying to fulfill the request. You can retry the
	// request.
	ErrCodeDependencyTimeoutException = "DependencyTimeoutException"

	// ErrCodeDisabledException for service response error code
	// "DisabledException".
	//
	// The request was rejected because the specified KMS key is not enabled.
	ErrCodeDisabledException = "DisabledException"

	// ErrCodeDryRunOperationException for service response error code
	// "DryRunOperationException".
	//
	// The request was rejected because the DryRun parameter was specified.
	ErrCodeDryRunOperationException = "DryRunOperationException"

	// ErrCodeExpiredImportTokenException for service response error code
	// "ExpiredImportTokenException".
	//
	// The request was rejected because the specified import token is expired. Use
	// GetParametersForImport to get a new import token and public key, use the
	// new public key to encrypt the key material, and then try the request again.
	ErrCodeExpiredImportTokenException = "ExpiredImportTokenException"

	// ErrCodeIncorrectKeyException for service response error code
	// "IncorrectKeyException".
	//
	// The request was rejected because the specified KMS key cannot decrypt the
	// data. The KeyId in a Decrypt request and the SourceKeyId in a ReEncrypt request
	// must identify the same KMS key that was used to encrypt the ciphertext.
	ErrCodeIncorrectKeyException = "IncorrectKeyException"

	// ErrCodeIncorrectKeyMaterialException for service response error code
	// "IncorrectKeyMaterialException".
	//
	// The request was rejected because the key material in the request is, expired,
	// invalid, or is not the same key material that was previously imported into
	// this KMS key.
	ErrCodeIncorrectKeyMaterialException = "IncorrectKeyMaterialException"

	// ErrCodeIncorrectTrustAnchorException for service response error code
	// "IncorrectTrustAnchorException".
	//
	// The request was rejected because the trust anchor certificate in the request
	// to create an CloudHSM key store is not the trust anchor certificate for the
	// specified CloudHSM cluster.
	//
	// When you initialize the CloudHSM cluster (https://docs.aws.amazon.com/cloudhsm/latest/userguide/initialize-cluster.html#sign-csr),
	// you create the trust anchor certificate and save it in the customerCA.crt
	// file.
	ErrCodeIncorrectTrustAnchorException = "IncorrectTrustAnchorException"

	// ErrCodeInternalException for service response error code
	// "KMSInternalException".
	//
	// The request was rejected because an internal exception occurred. The request
	// can be retried.
	ErrCodeInternalException = "KMSInternalException"

	// ErrCodeInvalidAliasNameException for service response error code
	// "InvalidAliasNameException".
	//
	// The request was rejected because the specified alias name is not valid.
	ErrCodeInvalidAliasNameException = "InvalidAliasNameException"

	// ErrCodeInvalidArnException for service response error code
	// "InvalidArnException".
	//
	// The request was rejected because a specified ARN, or an ARN in a key policy,
	// is not valid.
	ErrCodeInvalidArnException = "InvalidArnException"

	// ErrCodeInvalidCiphertextException for service response error code
	// "InvalidCiphertextException".
	//
	// From the Decrypt or ReEncrypt operation, the request was rejected because
	// the specified ciphertext, or additional authenticated data incorporated into
	// the ciphertext, such as the encryption context, is corrupted, missing, or
	// otherwise invalid.
	//
	// From the ImportKeyMaterial operation, the request was rejected because KMS
	// could not decrypt the encrypted (wrapped) key material.
	ErrCodeInvalidCiphertextException = "InvalidCiphertextException"

	// ErrCodeInvalidGrantIdException for service response error code
	// "InvalidGrantIdException".
	//
	// The request was rejected because the specified GrantId is not valid.
	ErrCodeInvalidGrantIdException = "InvalidGrantIdException"

	// ErrCodeInvalidGrantTokenException for service response error code
	// "InvalidGrantTokenException".
	//
	// The request was rejected because the specified grant token is not valid.
	ErrCodeInvalidGrantTokenException = "InvalidGrantTokenException"

	// ErrCodeInvalidImportTokenException for service response error code
	// "InvalidImportTokenException".
	//
	// The request was rejected because the provided import token is invalid or
	// is associated with a different KMS key.
	ErrCodeInvalidImportTokenException = "InvalidImportTokenException"

	// ErrCodeInvalidKeyUsageException for service response error code
	// "InvalidKeyUsageException".
	//
	// The request was rejected for one of the following reasons:
	//
	//    * The KeyUsage value of the KMS key is incompatible with the API operation.
	//
	//    * The encryption algorithm or signing algorithm specified for the operation
	//    is incompatible with the type of key material in the KMS key (KeySpec).
	//
	// For encrypting, decrypting, re-encrypting, and generating data keys, the
	// KeyUsage must be ENCRYPT_DECRYPT. For signing and verifying messages, the
	// KeyUsage must be SIGN_VERIFY. For generating and verifying message authentication
	// codes (MACs), the KeyUsage must be GENERATE_VERIFY_MAC. To find the KeyUsage
	// of a KMS key, use the DescribeKey operation.
	//
	// To find the encryption or signing algorithms supported for a particular KMS
	// key, use the DescribeKey operation.
	ErrCodeInvalidKeyUsageException = "InvalidKeyUsageException"

	// ErrCodeInvalidMarkerException for service response error code
	// "InvalidMarkerException".
	//
	// The request was rejected because the marker that specifies where pagination
	// should next begin is not valid.
	ErrCodeInvalidMarkerException = "InvalidMarkerException"

	// ErrCodeInvalidStateException for service response error code
	// "KMSInvalidStateException".
	//
	// The request was rejected because the state of the specified resource is not
	// valid for this request.
	//
	// This exceptions means one of the following:
	//
	//    * The key state of the KMS key is not compatible with the operation. To
	//    find the key state, use the DescribeKey operation. For more information
	//    about which key states are compatible with each KMS operation, see Key
	//    states of KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
	//    in the Key Management Service Developer Guide .
	//
	//    * For cryptographic operations on KMS keys in custom key stores, this
	//    exception represents a general failure with many possible causes. To identify
	//    the cause, see the error message that accompanies the exception.
	ErrCodeInvalidStateException = "KMSInvalidStateException"

	// ErrCodeKMSInvalidMacException for service response error code
	// "KMSInvalidMacException".
	//
	// The request was rejected because the HMAC verification failed. HMAC verification
	// fails when the HMAC computed by using the specified message, HMAC KMS key,
	// and MAC algorithm does not match the HMAC specified in the request.
	ErrCodeKMSInvalidMacException = "KMSInvalidMacException"

	// ErrCodeKMSInvalidSignatureException for service response error code
	// "KMSInvalidSignatureException".
	//
	// The request was rejected because the signature verification failed. Signature
	// verification fails when it cannot confirm that signature was produced by
	// signing the specified message with the specified KMS key and signing algorithm.
	ErrCodeKMSInvalidSignatureException = "KMSInvalidSignatureException"

	// ErrCodeKeyUnavailableException for service response error code
	// "KeyUnavailableException".
	//
	// The request was rejected because the specified KMS key was not available.
	// You can retry the request.
	ErrCodeKeyUnavailableException = "KeyUnavailableException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The request was rejected because a quota was exceeded. For more information,
	// see Quotas (https://docs.aws.amazon.com/kms/latest/developerguide/limits.html)
	// in the Key Management Service Developer Guide.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMalformedPolicyDocumentException for service response error code
	// "MalformedPolicyDocumentException".
	//
	// The request was rejected because the specified policy is not syntactically
	// or semantically correct.
	ErrCodeMalformedPolicyDocumentException = "MalformedPolicyDocumentException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The request was rejected because the specified entity or resource could not
	// be found.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeTagException for service response error code
	// "TagException".
	//
	// The request was rejected because one or more tags are not valid.
	ErrCodeTagException = "TagException"

	// ErrCodeUnsupportedOperationException for service response error code
	// "UnsupportedOperationException".
	//
	// The request was rejected because a specified parameter is not supported or
	// a specified resource is not valid for this operation.
	ErrCodeUnsupportedOperationException = "UnsupportedOperationException"

	// ErrCodeXksKeyAlreadyInUseException for service response error code
	// "XksKeyAlreadyInUseException".
	//
	// The request was rejected because the (XksKeyId) is already associated with
	// a KMS key in this external key store. Each KMS key in an external key store
	// must be associated with a different external key.
	ErrCodeXksKeyAlreadyInUseException = "XksKeyAlreadyInUseException"

	// ErrCodeXksKeyInvalidConfigurationException for service response error code
	// "XksKeyInvalidConfigurationException".
	//
	// The request was rejected because the external key specified by the XksKeyId
	// parameter did not meet the configuration requirements for an external key
	// store.
	//
	// The external key must be an AES-256 symmetric key that is enabled and performs
	// encryption and decryption.
	ErrCodeXksKeyInvalidConfigurationException = "XksKeyInvalidConfigurationException"

	// ErrCodeXksKeyNotFoundException for service response error code
	// "XksKeyNotFoundException".
	//
	// The request was rejected because the external key store proxy could not find
	// the external key. This exception is thrown when the value of the XksKeyId
	// parameter doesn't identify a key in the external key manager associated with
	// the external key proxy.
	//
	// Verify that the XksKeyId represents an existing key in the external key manager.
	// Use the key identifier that the external key store proxy uses to identify
	// the key. For details, see the documentation provided with your external key
	// store proxy or key manager.
	ErrCodeXksKeyNotFoundException = "XksKeyNotFoundException"

	// ErrCodeXksProxyIncorrectAuthenticationCredentialException for service response error code
	// "XksProxyIncorrectAuthenticationCredentialException".
	//
	// The request was rejected because the proxy credentials failed to authenticate
	// to the specified external key store proxy. The specified external key store
	// proxy rejected a status request from KMS due to invalid credentials. This
	// can indicate an error in the credentials or in the identification of the
	// external key store proxy.
	ErrCodeXksProxyIncorrectAuthenticationCredentialException = "XksProxyIncorrectAuthenticationCredentialException"

	// ErrCodeXksProxyInvalidConfigurationException for service response error code
	// "XksProxyInvalidConfigurationException".
	//
	// The request was rejected because the Amazon VPC endpoint service configuration
	// does not fulfill the requirements for an external key store proxy. For details,
	// see the exception message.
	ErrCodeXksProxyInvalidConfigurationException = "XksProxyInvalidConfigurationException"

	// ErrCodeXksProxyInvalidResponseException for service response error code
	// "XksProxyInvalidResponseException".
	//
	// KMS cannot interpret the response it received from the external key store
	// proxy. The problem might be a poorly constructed response, but it could also
	// be a transient network issue. If you see this error repeatedly, report it
	// to the proxy vendor.
	ErrCodeXksProxyInvalidResponseException = "XksProxyInvalidResponseException"

	// ErrCodeXksProxyUriEndpointInUseException for service response error code
	// "XksProxyUriEndpointInUseException".
	//
	// The request was rejected because the concatenation of the XksProxyUriEndpoint
	// is already associated with an external key store in the Amazon Web Services
	// account and Region. Each external key store in an account and Region must
	// use a unique external key store proxy address.
	ErrCodeXksProxyUriEndpointInUseException = "XksProxyUriEndpointInUseException"

	// ErrCodeXksProxyUriInUseException for service response error code
	// "XksProxyUriInUseException".
	//
	// The request was rejected because the concatenation of the XksProxyUriEndpoint
	// and XksProxyUriPath is already associated with an external key store in the
	// Amazon Web Services account and Region. Each external key store in an account
	// and Region must use a unique external key store proxy API address.
	ErrCodeXksProxyUriInUseException = "XksProxyUriInUseException"

	// ErrCodeXksProxyUriUnreachableException for service response error code
	// "XksProxyUriUnreachableException".
	//
	// KMS was unable to reach the specified XksProxyUriPath. The path must be reachable
	// before you create the external key store or update its settings.
	//
	// This exception is also thrown when the external key store proxy response
	// to a GetHealthStatus request indicates that all external key manager instances
	// are unavailable.
	ErrCodeXksProxyUriUnreachableException = "XksProxyUriUnreachableException"

	// ErrCodeXksProxyVpcEndpointServiceInUseException for service response error code
	// "XksProxyVpcEndpointServiceInUseException".
	//
	// The request was rejected because the specified Amazon VPC endpoint service
	// is already associated with an external key store in the Amazon Web Services
	// account and Region. Each external key store in an Amazon Web Services account
	// and Region must use a different Amazon VPC endpoint service.
	ErrCodeXksProxyVpcEndpointServiceInUseException = "XksProxyVpcEndpointServiceInUseException"

	// ErrCodeXksProxyVpcEndpointServiceInvalidConfigurationException for service response error code
	// "XksProxyVpcEndpointServiceInvalidConfigurationException".
	//
	// The request was rejected because the Amazon VPC endpoint service configuration
	// does not fulfill the requirements for an external key store proxy. For details,
	// see the exception message and review the requirements (https://docs.aws.amazon.com/kms/latest/developerguide/vpc-connectivity.html#xks-vpc-requirements)
	// for Amazon VPC endpoint service connectivity for an external key store.
	ErrCodeXksProxyVpcEndpointServiceInvalidConfigurationException = "XksProxyVpcEndpointServiceInvalidConfigurationException"

	// ErrCodeXksProxyVpcEndpointServiceNotFoundException for service response error code
	// "XksProxyVpcEndpointServiceNotFoundException".
	//
	// The request was rejected because KMS could not find the specified VPC endpoint
	// service. Use DescribeCustomKeyStores to verify the VPC endpoint service name
	// for the external key store. Also, confirm that the Allow principals list
	// for the VPC endpoint service includes the KMS service principal for the Region,
	// such as cks.kms.us-east-1.amazonaws.com.
	ErrCodeXksProxyVpcEndpointServiceNotFoundException = "XksProxyVpcEndpointServiceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AlreadyExistsException":                                  newErrorAlreadyExistsException,
	"CloudHsmClusterInUseException":                           newErrorCloudHsmClusterInUseException,
	"CloudHsmClusterInvalidConfigurationException":            newErrorCloudHsmClusterInvalidConfigurationException,
	"CloudHsmClusterNotActiveException":                       newErrorCloudHsmClusterNotActiveException,
	"CloudHsmClusterNotFoundException":                        newErrorCloudHsmClusterNotFoundException,
	"CloudHsmClusterNotRelatedException":                      newErrorCloudHsmClusterNotRelatedException,
	"CustomKeyStoreHasCMKsException":                          newErrorCustomKeyStoreHasCMKsException,
	"CustomKeyStoreInvalidStateException":                     newErrorCustomKeyStoreInvalidStateException,
	"CustomKeyStoreNameInUseException":                        newErrorCustomKeyStoreNameInUseException,
	"CustomKeyStoreNotFoundException":                         newErrorCustomKeyStoreNotFoundException,
	"DependencyTimeoutException":                              newErrorDependencyTimeoutException,
	"DisabledException":                                       newErrorDisabledException,
	"DryRunOperationException":                                newErrorDryRunOperationException,
	"ExpiredImportTokenException":                             newErrorExpiredImportTokenException,
	"IncorrectKeyException":                                   newErrorIncorrectKeyException,
	"IncorrectKeyMaterialException":                           newErrorIncorrectKeyMaterialException,
	"IncorrectTrustAnchorException":                           newErrorIncorrectTrustAnchorException,
	"KMSInternalException":                                    newErrorInternalException,
	"InvalidAliasNameException":                               newErrorInvalidAliasNameException,
	"InvalidArnException":                                     newErrorInvalidArnException,
	"InvalidCiphertextException":                              newErrorInvalidCiphertextException,
	"InvalidGrantIdException":                                 newErrorInvalidGrantIdException,
	"InvalidGrantTokenException":                              newErrorInvalidGrantTokenException,
	"InvalidImportTokenException":                             newErrorInvalidImportTokenException,
	"InvalidKeyUsageException":                                newErrorInvalidKeyUsageException,
	"InvalidMarkerException":                                  newErrorInvalidMarkerException,
	"KMSInvalidStateException":                                newErrorInvalidStateException,
	"KMSInvalidMacException":                                  newErrorKMSInvalidMacException,
	"KMSInvalidSignatureException":                            newErrorKMSInvalidSignatureException,
	"KeyUnavailableException":                                 newErrorKeyUnavailableException,
	"LimitExceededException":                                  newErrorLimitExceededException,
	"MalformedPolicyDocumentException":                        newErrorMalformedPolicyDocumentException,
	"NotFoundException":                                       newErrorNotFoundException,
	"TagException":                                            newErrorTagException,
	"UnsupportedOperationException":                           newErrorUnsupportedOperationException,
	"XksKeyAlreadyInUseException":                             newErrorXksKeyAlreadyInUseException,
	"XksKeyInvalidConfigurationException":                     newErrorXksKeyInvalidConfigurationException,
	"XksKeyNotFoundException":                                 newErrorXksKeyNotFoundException,
	"XksProxyIncorrectAuthenticationCredentialException":      newErrorXksProxyIncorrectAuthenticationCredentialException,
	"XksProxyInvalidConfigurationException":                   newErrorXksProxyInvalidConfigurationException,
	"XksProxyInvalidResponseException":                        newErrorXksProxyInvalidResponseException,
	"XksProxyUriEndpointInUseException":                       newErrorXksProxyUriEndpointInUseException,
	"XksProxyUriInUseException":                               newErrorXksProxyUriInUseException,
	"XksProxyUriUnreachableException":                         newErrorXksProxyUriUnreachableException,
	"XksProxyVpcEndpointServiceInUseException":                newErrorXksProxyVpcEndpointServiceInUseException,
	"XksProxyVpcEndpointServiceInvalidConfigurationException": newErrorXksProxyVpcEndpointServiceInvalidConfigurationException,
	"XksProxyVpcEndpointServiceNotFoundException":             newErrorXksProxyVpcEndpointServiceNotFoundException,
}
