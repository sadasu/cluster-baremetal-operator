package provisioning

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	fakekube "k8s.io/client-go/kubernetes/fake"
	faketesting "k8s.io/client-go/testing"
)

const testNamespace = "test-namespce"

func TestGenerateRandomPassword(t *testing.T) {
	pwd1, err := generateRandomPassword()
	if err != nil {
		t.Errorf("Unexpected error while generating random password: %s", err)
	}
	if pwd1 == "" {
		t.Errorf("Expected a valid string but got null")
	}
	pwd2, err := generateRandomPassword()
	if err != nil {
		t.Errorf("Unexpected error while re-generating random password: %s", err)
	} else {
		assert.False(t, pwd1 == pwd2, "regenerated random password should not match pervious one")
	}
}

func TestCreateMariadbPasswordSecretNew(t *testing.T) {

	cases := []struct {
		name string

		secretError   *errors.StatusError
		expectedError error
	}{
		{
			name:          "new-secret",
			expectedError: nil,
		},
		{
			name: "error-fetching-secret",

			secretError:   errors.NewServiceUnavailable("an error"),
			expectedError: errors.NewServiceUnavailable("an error"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			secretsResource := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "secrets"}

			kubeClient := fakekube.NewSimpleClientset(nil...)

			if tc.secretError != nil {
				kubeClient.Fake.PrependReactor("get", "secrets", func(action faketesting.Action) (handled bool, ret runtime.Object, err error) {
					return true, &v1.Secret{}, tc.secretError
				})
			}

			err := CreateMariadbPasswordSecret(kubeClient.CoreV1(), testNamespace)
			assert.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				secret, _ := kubeClient.Tracker().Get(secretsResource, testNamespace, "metal3-mariadb-password")
				assert.NotEmpty(t, secret.(*v1.Secret).StringData[baremetalSecretKey])
			}
		})
	}
}

func TestCreateIronicPasswordSecret(t *testing.T) {
	kubeClient := fakekube.NewSimpleClientset(nil...)
	client := kubeClient.CoreV1()

	err := CreateIronicPasswordSecret(client, testNamespace)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	// Check if Ironic secret exits
	secret, err := client.Secrets(testNamespace).Get(context.Background(), ironicSecretName, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		t.Errorf("Error creating Ironic secret.")
	}
	assert.True(t, strings.Compare(secret.StringData[ironicUsernameKey], ironicUsername) == 0, "ironic password created incorrectly")
	return
}

func TestCreateInspectorPasswordSecret(t *testing.T) {
	kubeClient := fakekube.NewSimpleClientset(nil...)
	client := kubeClient.CoreV1()

	err := CreateInspectorPasswordSecret(client, testNamespace)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	// Check if Ironic Inspector secret exits
	secret, err := client.Secrets(testNamespace).Get(context.Background(), inspectorSecretName, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		t.Errorf("Error creating Ironic Inspector secret.")
	}
	assert.True(t, strings.Compare(secret.StringData[ironicUsernameKey], inspectorUsername) == 0, "inspector password created incorrectly")
	return
}
