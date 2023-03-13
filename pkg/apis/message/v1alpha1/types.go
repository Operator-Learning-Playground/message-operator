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
}



type MessageSpec struct {
	Sender  Sender   `json:"sender"`
}

type Sender struct {
	remote   string	 `json:"remote"`
	port     int     `json:"port"`
	email    string  `json:"email"`
	password string  `json:"password"`
	targets  string  `json:"targets"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MessageList
type MessageList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Message `json:"items"`
}


