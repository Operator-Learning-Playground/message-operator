package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Message
type Message struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MessageSpec `json:"spec,omitempty"`

	Status MessageStatus `json:"status,omitempty"`
}

type MessageSpec struct {
	Sender Sender `json:"sender"`
	Feishu Feishu `json:"feishu"`
}

type MessageStatus struct {
	Generation int64 `json:"generation"`
}

type Sender struct {
	Open     bool   `json:"open"`
	Remote   string `json:"remote"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Targets  string `json:"targets"`
}

type Feishu struct {
	Open    bool   `json:"open"`
	Webhook string `json:"webhook"`
	Type    string `json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MessageList
type MessageList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Message `json:"items"`
}
