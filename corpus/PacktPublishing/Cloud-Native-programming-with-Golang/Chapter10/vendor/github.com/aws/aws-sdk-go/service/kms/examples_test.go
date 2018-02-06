// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleKMS_CancelKeyDeletion() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.CancelKeyDeletionInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.CancelKeyDeletion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_CreateAlias() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.CreateAliasInput{
		AliasName:   aws.String("AliasNameType"), // Required
		TargetKeyId: aws.String("KeyIdType"),     // Required
	}
	resp, err := svc.CreateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_CreateGrant() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.CreateGrantInput{
		GranteePrincipal: aws.String("PrincipalIdType"), // Required
		KeyId:            aws.String("KeyIdType"),       // Required
		Constraints: &kms.GrantConstraints{
			EncryptionContextEquals: map[string]*string{
				"Key": aws.String("EncryptionContextValue"), // Required
				// More values...
			},
			EncryptionContextSubset: map[string]*string{
				"Key": aws.String("EncryptionContextValue"), // Required
				// More values...
			},
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		Name: aws.String("GrantNameType"),
		Operations: []*string{
			aws.String("GrantOperation"), // Required
			// More values...
		},
		RetiringPrincipal: aws.String("PrincipalIdType"),
	}
	resp, err := svc.CreateGrant(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_CreateKey() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.CreateKeyInput{
		BypassPolicyLockoutSafetyCheck: aws.Bool(true),
		Description:                    aws.String("DescriptionType"),
		KeyUsage:                       aws.String("KeyUsageType"),
		Origin:                         aws.String("OriginType"),
		Policy:                         aws.String("PolicyType"),
		Tags: []*kms.Tag{
			{ // Required
				TagKey:   aws.String("TagKeyType"),   // Required
				TagValue: aws.String("TagValueType"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.CreateKey(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_Decrypt() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.DecryptInput{
		CiphertextBlob: []byte("PAYLOAD"), // Required
		EncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
	}
	resp, err := svc.Decrypt(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_DeleteAlias() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.DeleteAliasInput{
		AliasName: aws.String("AliasNameType"), // Required
	}
	resp, err := svc.DeleteAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_DeleteImportedKeyMaterial() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.DeleteImportedKeyMaterialInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.DeleteImportedKeyMaterial(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_DescribeKey() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.DescribeKeyInput{
		KeyId: aws.String("KeyIdType"), // Required
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
	}
	resp, err := svc.DescribeKey(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_DisableKey() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.DisableKeyInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.DisableKey(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_DisableKeyRotation() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.DisableKeyRotationInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.DisableKeyRotation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_EnableKey() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.EnableKeyInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.EnableKey(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_EnableKeyRotation() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.EnableKeyRotationInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.EnableKeyRotation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_Encrypt() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.EncryptInput{
		KeyId:     aws.String("KeyIdType"), // Required
		Plaintext: []byte("PAYLOAD"),       // Required
		EncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
	}
	resp, err := svc.Encrypt(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GenerateDataKey() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.GenerateDataKeyInput{
		KeyId: aws.String("KeyIdType"), // Required
		EncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		KeySpec:       aws.String("DataKeySpec"),
		NumberOfBytes: aws.Int64(1),
	}
	resp, err := svc.GenerateDataKey(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GenerateDataKeyWithoutPlaintext() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.GenerateDataKeyWithoutPlaintextInput{
		KeyId: aws.String("KeyIdType"), // Required
		EncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		KeySpec:       aws.String("DataKeySpec"),
		NumberOfBytes: aws.Int64(1),
	}
	resp, err := svc.GenerateDataKeyWithoutPlaintext(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GenerateRandom() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.GenerateRandomInput{
		NumberOfBytes: aws.Int64(1),
	}
	resp, err := svc.GenerateRandom(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GetKeyPolicy() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.GetKeyPolicyInput{
		KeyId:      aws.String("KeyIdType"),      // Required
		PolicyName: aws.String("PolicyNameType"), // Required
	}
	resp, err := svc.GetKeyPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GetKeyRotationStatus() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.GetKeyRotationStatusInput{
		KeyId: aws.String("KeyIdType"), // Required
	}
	resp, err := svc.GetKeyRotationStatus(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_GetParametersForImport() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.GetParametersForImportInput{
		KeyId:             aws.String("KeyIdType"),       // Required
		WrappingAlgorithm: aws.String("AlgorithmSpec"),   // Required
		WrappingKeySpec:   aws.String("WrappingKeySpec"), // Required
	}
	resp, err := svc.GetParametersForImport(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ImportKeyMaterial() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.ImportKeyMaterialInput{
		EncryptedKeyMaterial: []byte("PAYLOAD"),       // Required
		ImportToken:          []byte("PAYLOAD"),       // Required
		KeyId:                aws.String("KeyIdType"), // Required
		ExpirationModel:      aws.String("ExpirationModelType"),
		ValidTo:              aws.Time(time.Now()),
	}
	resp, err := svc.ImportKeyMaterial(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListAliases() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.ListAliasesInput{
		Limit:  aws.Int64(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListAliases(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListGrants() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.ListGrantsInput{
		KeyId:  aws.String("KeyIdType"), // Required
		Limit:  aws.Int64(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListGrants(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListKeyPolicies() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.ListKeyPoliciesInput{
		KeyId:  aws.String("KeyIdType"), // Required
		Limit:  aws.Int64(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListKeyPolicies(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListKeys() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.ListKeysInput{
		Limit:  aws.Int64(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListKeys(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListResourceTags() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.ListResourceTagsInput{
		KeyId:  aws.String("KeyIdType"), // Required
		Limit:  aws.Int64(1),
		Marker: aws.String("MarkerType"),
	}
	resp, err := svc.ListResourceTags(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ListRetirableGrants() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.ListRetirableGrantsInput{
		RetiringPrincipal: aws.String("PrincipalIdType"), // Required
		Limit:             aws.Int64(1),
		Marker:            aws.String("MarkerType"),
	}
	resp, err := svc.ListRetirableGrants(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_PutKeyPolicy() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.PutKeyPolicyInput{
		KeyId:                          aws.String("KeyIdType"),      // Required
		Policy:                         aws.String("PolicyType"),     // Required
		PolicyName:                     aws.String("PolicyNameType"), // Required
		BypassPolicyLockoutSafetyCheck: aws.Bool(true),
	}
	resp, err := svc.PutKeyPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ReEncrypt() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.ReEncryptInput{
		CiphertextBlob:   []byte("PAYLOAD"),       // Required
		DestinationKeyId: aws.String("KeyIdType"), // Required
		DestinationEncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
		GrantTokens: []*string{
			aws.String("GrantTokenType"), // Required
			// More values...
		},
		SourceEncryptionContext: map[string]*string{
			"Key": aws.String("EncryptionContextValue"), // Required
			// More values...
		},
	}
	resp, err := svc.ReEncrypt(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_RetireGrant() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.RetireGrantInput{
		GrantId:    aws.String("GrantIdType"),
		GrantToken: aws.String("GrantTokenType"),
		KeyId:      aws.String("KeyIdType"),
	}
	resp, err := svc.RetireGrant(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_RevokeGrant() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.RevokeGrantInput{
		GrantId: aws.String("GrantIdType"), // Required
		KeyId:   aws.String("KeyIdType"),   // Required
	}
	resp, err := svc.RevokeGrant(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_ScheduleKeyDeletion() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.ScheduleKeyDeletionInput{
		KeyId:               aws.String("KeyIdType"), // Required
		PendingWindowInDays: aws.Int64(1),
	}
	resp, err := svc.ScheduleKeyDeletion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_TagResource() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.TagResourceInput{
		KeyId: aws.String("KeyIdType"), // Required
		Tags: []*kms.Tag{ // Required
			{ // Required
				TagKey:   aws.String("TagKeyType"),   // Required
				TagValue: aws.String("TagValueType"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.TagResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_UntagResource() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.UntagResourceInput{
		KeyId: aws.String("KeyIdType"), // Required
		TagKeys: []*string{ // Required
			aws.String("TagKeyType"), // Required
			// More values...
		},
	}
	resp, err := svc.UntagResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_UpdateAlias() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.UpdateAliasInput{
		AliasName:   aws.String("AliasNameType"), // Required
		TargetKeyId: aws.String("KeyIdType"),     // Required
	}
	resp, err := svc.UpdateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleKMS_UpdateKeyDescription() {
	sess := session.Must(session.NewSession())

	svc := kms.New(sess)

	params := &kms.UpdateKeyDescriptionInput{
		Description: aws.String("DescriptionType"), // Required
		KeyId:       aws.String("KeyIdType"),       // Required
	}
	resp, err := svc.UpdateKeyDescription(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
